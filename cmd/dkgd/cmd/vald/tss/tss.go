package tss

import (
	"context"
	"encoding/hex"
	"fmt"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	broadcast "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/parse"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss/rpc"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"
	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tss "github.com/fairblock/dkg-core/x/tss/types"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
	"io"
	"strconv"
	"sync"
	"time"
)

type KeygenEvent struct {
	Type       string             `json:"type"`
	Attributes []KeygenAttributes `json:"attributes"`
}
type KeygenAttributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type EventMsg struct {
	MsgIndex int           `json:"msg_index"`
	Events   []KeygenEvent `json:"events"`
}
type Session struct {
	ID        string
	TimeoutAt int64
	timeout   chan struct{}
}

var faultersMap = make(map[string]bool)

func (s *Session) Timeout() {
	close(s.timeout)
}
func (s *Session) WaitForTimeout() {
	<-s.timeout
}

type TimeoutQueue struct {
	lock  sync.RWMutex
	queue []*Session
}

func (q *TimeoutQueue) Enqueue(ID string, timeoutAt int64) *Session {
	q.lock.Lock()
	defer q.lock.Unlock()
	session := Session{ID: ID, TimeoutAt: timeoutAt, timeout: make(chan struct{})}
	q.queue = append(q.queue, &session)
	return &session
}
func (q *TimeoutQueue) Dequeue() *Session {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
func (q *TimeoutQueue) Top() *Session {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.queue) == 0 {
		return nil
	}
	return q.queue[0]
}
func NewTimeoutQueue() *TimeoutQueue {
	return &TimeoutQueue{
		lock:  sync.RWMutex{},
		queue: []*Session{},
	}
}

type Stream interface {
	Send(in *tofnd.MessageIn) error
	Recv() (*tofnd.MessageOut, error)
	CloseSend() error
}
type LockableStream struct {
	sendLock sync.Mutex
	recvLock sync.Mutex
	stream   Stream
}

func NewLockableStream(stream Stream) *LockableStream {
	return &LockableStream{
		sendLock: sync.Mutex{},
		recvLock: sync.Mutex{},
		stream:   stream,
	}
}
func (l *LockableStream) Send(in *tofnd.MessageIn) error {
	l.sendLock.Lock()
	defer l.sendLock.Unlock()
	return l.stream.Send(in)
}
func (l *LockableStream) Recv() (*tofnd.MessageOut, error) {
	l.recvLock.Lock()
	defer l.recvLock.Unlock()
	return l.stream.Recv()
}
func (l *LockableStream) CloseSend() error {
	l.sendLock.Lock()
	defer l.sendLock.Unlock()
	return l.stream.CloseSend()
}

type Mgr struct {
	tmClient      *tmclient.HTTP
	client        rpc.Client
	cliCtx        sdkClient.Context
	keygen        *sync.RWMutex
	keyId         string
	keygenStreams map[string]*LockableStream
	timeoutQueue  *TimeoutQueue
	Timeout       time.Duration
	principalAddr string
	Logger        log.Logger
	broadcaster   *broadcast.CosmosClient
	cdc           *codec.LegacyAmino
	startHeight   int
	currentHeight int
	me            int
}

func Connect(host string, port string, timeout time.Duration, logger log.Logger) (*grpc.ClientConn, error) {
	tofndServerAddress := host + ":" + port
	logger.Info(fmt.Sprintf("initiate connection to tofnd gRPC server: %s", tofndServerAddress))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return grpc.DialContext(ctx, tofndServerAddress, grpc.WithInsecure(), grpc.WithBlock())
}
func NewMgr(c *tmclient.HTTP, client rpc.Client, cliCtx sdkClient.Context, timeout time.Duration, principalAddr string, broadcaster *broadcast.CosmosClient, logger log.Logger, cdc *codec.LegacyAmino) *Mgr {
	return &Mgr{
		tmClient:      c,
		client:        client,
		cliCtx:        cliCtx,
		keygen:        &sync.RWMutex{},
		keygenStreams: make(map[string]*LockableStream),
		timeoutQueue:  NewTimeoutQueue(),
		Timeout:       timeout,
		principalAddr: principalAddr,
		Logger:        logger.With("listener", "tss"),
		broadcaster:   broadcaster,
		cdc:           cdc,
	}
}
func (mgr *Mgr) Recover(recoverJSON []byte) error {
	var requests []tofnd.RecoverRequest
	err := mgr.cdc.UnmarshalJSON(recoverJSON, &requests)
	if err != nil {
		panic(sdkerrors.Wrapf(err, "failed to unmarshal recovery requests"))
	}
	for _, request := range requests {
		uid := request.KeygenInit.PartyUids[int(request.KeygenInit.MyPartyIndex)]
		if mgr.principalAddr != uid {
			return fmt.Errorf("party UID mismatch (expected %s, got %s)", mgr.principalAddr, uid)
		}
		grpcCtx, cancel := context.WithTimeout(context.Background(), mgr.Timeout)
		defer cancel()
		response, err := mgr.client.Recover(grpcCtx, &request)
		if err != nil {
			panic(sdkerrors.Wrap(err,
				fmt.Sprintf("failed tofnd gRPC call Recover for key ID %s", request.KeygenInit.NewKeyUid)))
		}
		if response.GetResponse() == tofnd.RecoverResponse_RESPONSE_FAIL {
			return fmt.Errorf("failed to recover tofnd shares for validator %s and key ID %s", uid, request.KeygenInit.NewKeyUid)
		}
		mgr.Logger.Info(
			fmt.Sprintf("successfully recovered tofnd shares for validator %s and key ID %s", uid, request.KeygenInit.NewKeyUid))
	}
	return nil
}
func (mgr *Mgr) abortKeygen(keyID string) (err error) {
	stream, ok := mgr.getKeygenStream(keyID)
	if !ok {
		return nil
	}
	return abort(stream)
}
func (mgr *Mgr) extractRefundMsgResponse(res *sdk.TxResponse) (dkgnet.RefundMsgResponse, error) {
	var txMsg sdk.TxMsgData
	var refundReply dkgnet.RefundMsgResponse
	bz, err := hex.DecodeString(res.Data)
	if err != nil {
		return dkgnet.RefundMsgResponse{}, err
	}
	err = mgr.cdc.Unmarshal(bz, &txMsg)
	if err != nil {
		return dkgnet.RefundMsgResponse{}, err
	}
	for _, msg := range txMsg.Data {
		err = refundReply.Unmarshal(msg.Data)
		if err != nil {
			mgr.Logger.Debug(fmt.Sprintf("not a refund response: %v", err))
			continue
		}
		return refundReply, nil
	}
	return dkgnet.RefundMsgResponse{}, fmt.Errorf("no refund response found in tx response")
}
func abort(stream Stream) error {
	msg := &tofnd.MessageIn{
		Data: &tofnd.MessageIn_Abort{
			Abort: true,
		},
	}
	if err := stream.Send(msg); err != nil {
		panic(sdkerrors.Wrap(err, "failure to send abort msg to gRPC server"))
	}
	return nil
}
func handleStream(stream Stream, cancel context.CancelFunc, logger log.Logger, me int) (broadcast <-chan *tofnd.TrafficOut, result <-chan interface{}, err <-chan error) {
	broadcastChan := make(chan *tofnd.TrafficOut)
	resChan := make(chan interface{})
	errChan := make(chan error, 2)
	go func() {
		defer cancel()
		defer close(errChan)
		defer close(broadcastChan)
		defer close(resChan)
		defer func() {
			if err := stream.CloseSend(); err != nil {
				panic("handler goroutine: failure to CloseSend stream")
			}
		}()
		for {
			msgOneof, err := stream.Recv()
			if err == io.EOF {
				logger.Debug("handler goroutine: gRPC stream closed by server")
				return
			}
			if err != nil {
				fmt.Println("handler goroutine: failure to receive msg from gRPC server stream", me, err)
				return
			}
			switch msg := msgOneof.GetData().(type) {
			case *tofnd.MessageOut_Traffic:
				broadcastChan <- msg.Traffic
			case *tofnd.MessageOut_KeygenResult_:
				resChan <- msg.KeygenResult
				return
			default:
				errChan <- fmt.Errorf("handler goroutine: server stream should send only msg type")
				return
			}
		}
	}()
	return broadcastChan, resChan, errChan
}
func parseHeartBeatParams(cdc *codec.LegacyAmino, attributes map[string]string) []tss.KeyInfo {
	parsers := []*parse.AttributeParser{
		{Key: tss.AttributeKeyKeyInfos, Map: func(s string) (interface{}, error) {
			var keyInfos []tss.KeyInfo
			cdc.MustUnmarshalJSON([]byte(s), &keyInfos)
			return keyInfos, nil
		}},
	}
	results, err := parse.Parse(attributes, parsers)
	if err != nil {
		panic(err)
	}
	return results[0].([]tss.KeyInfo)
}
func parseMsgParamsOne(e types.Event) (sessionID string, from string, payload *tofnd.TrafficOut, index uint64) {
	
	if len(e.Attributes) == 0 {
		return "", "", nil, 1000000000000
	}
	innerMsg := e.Attributes[0].Value
	if len(e.Attributes) < 3 {
		return "", "", nil, 1000000000000
	}
	indexS := e.Attributes[2].Value
	index, err := strconv.ParseUint(string(indexS), 10, 64)
	if err != nil {
		fmt.Println("processig error")
		return "", "", nil, 1000000000000
	}
	msg := new(dkgnet.MsgRefundMsgRequest)
	err = msg.Unmarshal(innerMsg)
	if err != nil {
		fmt.Println("processig error")
		return "", "", nil, 1000000000000
	}
	msgVal := new(tss.ProcessKeygenTrafficRequest)
	err = msgVal.Unmarshal(msg.InnerMessage.Value)
	if err != nil {
		fmt.Println("processig error")
		return "", "", nil, 1000000000000
	}
	
	return msgVal.SessionID, msgVal.Sender.String(), msgVal.Payload, index
}
func parseMsgParams(e []types.Event) (sessionID string, from string, payload *tofnd.TrafficOut, index uint64) {
	innerMsg := e[4].Attributes[0].Value
	indexS := e[4].Attributes[2].Value
	index, _ = strconv.ParseUint(string(indexS), 10, 64)
	msg := new(dkgnet.MsgRefundMsgRequest)
	msg.Unmarshal(innerMsg)
	msgVal := new(tss.ProcessKeygenTrafficRequest)
	msgVal.Unmarshal(msg.InnerMessage.Value)
	return msgVal.SessionID, msgVal.Sender.String(), msgVal.Payload, index
}
func parseMsgParamsDispute(e KeygenEvent) (sessionID string, from string, payload *tofnd.TrafficOut, id uint64) {
	
	if len(e.Attributes) < 4 {
		return
	}
	innerMsg := (e.Attributes[0].Value)
	keyId := e.Attributes[1].Value
	from = e.Attributes[2].Value
	idString := string(e.Attributes[3].Value)
	i, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Oops, an error occurred:", err)
		return
	}
	value, err := strconv.Atoi(innerMsg)
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	b := []byte{byte(value)}
	id = uint64(i)
	
	tofndT := tofnd.TrafficOut{ToPartyUid: "", Payload: b, IsBroadcast: true}
	return keyId, sdk.AccAddress([]byte(from)).String(), &tofndT, id
}
func parseMsgParamsDisputeOne(e []types.Event) (sessionID string, from string, payload *tofnd.TrafficOut, id uint64) {
	
	if len(e[4].Attributes) < 4 {
		return
	}
	innerMsg := e[4].Attributes[0].Value
	sessionID = string(e[4].Attributes[1].Value)
	from = string(e[4].Attributes[2].Value)
	idString := string(e[4].Attributes[3].Value)
	i, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Oops, an error occurred:", err)
		return
	}
	value, err := strconv.Atoi(string(innerMsg))
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	b := []byte{byte(value)}
	id = uint64(i)
	tofndT := tofnd.TrafficOut{ToPartyUid: "", Payload: b, IsBroadcast: true}
	return sessionID, sdk.AccAddress([]byte(from)).String(), &tofndT, id
}
func prepareTrafficIn(principalAddr string, from string, sessionID string, payload *tofnd.TrafficOut, logger log.Logger) *tofnd.MessageIn {
	msgIn := &tofnd.MessageIn{
		Data: &tofnd.MessageIn_Traffic{
			Traffic: &tofnd.TrafficIn{
				Payload:      payload.Payload,
				IsBroadcast:  payload.IsBroadcast,
				FromPartyUid: from,
				RoundNum:     payload.RoundNum,
			},
		},
	}
	logger.Debug(fmt.Sprintf("incoming msg to tofnd: session [%.20s] from [%.20s] to [%.20s] broadcast [%t] me [%.20s]",
		sessionID, from, payload.ToPartyUid, payload.IsBroadcast, principalAddr))
	return msgIn
}
