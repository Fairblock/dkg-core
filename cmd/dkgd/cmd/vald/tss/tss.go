package tss

import (
	//"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	broadcast "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"
	"github.com/tendermint/tendermint/abci/types"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"

	//sdkFlags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"

	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/parse"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss/rpc"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"

	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tss "github.com/fairblock/dkg-core/x/tss/types"
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

// Session defines a tss session which is either signing or keygen
type Session struct {
	ID        string
	TimeoutAt int64
	timeout   chan struct{}
}
var faultersMap = make(map[string]bool)
// Timeout signals a session has timed out
func (s *Session) Timeout() {
	close(s.timeout)
}

// WaitForTimeout waits until the session has timed out
func (s *Session) WaitForTimeout() {
	<-s.timeout
}

// TimeoutQueue is a queue of sessions order by timeoutAt
type TimeoutQueue struct {
	lock  sync.RWMutex
	queue []*Session
}

// Enqueue adds a new session with ID and timeoutAt into the queue
func (q *TimeoutQueue) Enqueue(ID string, timeoutAt int64) *Session {
	q.lock.Lock()
	defer q.lock.Unlock()

	session := Session{ID: ID, TimeoutAt: timeoutAt, timeout: make(chan struct{})}
	q.queue = append(q.queue, &session)

	return &session
}

// Dequeue pops the first session in queue
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

// Top returns the first session in queue
func (q *TimeoutQueue) Top() *Session {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if len(q.queue) == 0 {
		return nil
	}

	return q.queue[0]
}

// NewTimeoutQueue is the constructor for TimeoutQueue
func NewTimeoutQueue() *TimeoutQueue {
	return &TimeoutQueue{
		lock:  sync.RWMutex{},
		queue: []*Session{},
	}
}

// Stream is the abstracted communication stream with tofnd
type Stream interface {
	Send(in *tofnd.MessageIn) error
	Recv() (*tofnd.MessageOut, error)
	CloseSend() error
}

// LockableStream is a thread-safe Stream
type LockableStream struct {
	sendLock sync.Mutex
	recvLock sync.Mutex
	stream   Stream
}

// NewLockableStream return a new thread-safe stream instance
func NewLockableStream(stream Stream) *LockableStream {
	return &LockableStream{
		sendLock: sync.Mutex{},
		recvLock: sync.Mutex{},
		stream:   stream,
	}
}

// Send implements the Stream interface
func (l *LockableStream) Send(in *tofnd.MessageIn) error {
	l.sendLock.Lock()
	defer l.sendLock.Unlock()

	return l.stream.Send(in)
}

// Recv implements the Stream interface
func (l *LockableStream) Recv() (*tofnd.MessageOut, error) {
	l.recvLock.Lock()
	defer l.recvLock.Unlock()

	return l.stream.Recv()
}

// CloseSend implements the Stream interface
func (l *LockableStream) CloseSend() error {
	l.sendLock.Lock()
	defer l.sendLock.Unlock()

	return l.stream.CloseSend()
}

// Mgr represents an object that manages all communication with the external tss process
type Mgr struct {
	tmClient *tmclient.HTTP
	client rpc.Client
	//multiSigClient rpc.MultiSigClient
	cliCtx sdkClient.Context
	keygen *sync.RWMutex
	keyId  string
	//sign           *sync.RWMutex
	keygenStreams map[string]*LockableStream
	//signStreams    map[string]*LockableStream
	timeoutQueue  *TimeoutQueue
	Timeout       time.Duration
	principalAddr string
	Logger        log.Logger
	broadcaster   *broadcast.CosmosClient
	cdc           *codec.LegacyAmino
	startHeight   int
	currentHeight int
	me int
}

// Connect connects to tofnd gRPC Server
func Connect(host string, port string, timeout time.Duration, logger log.Logger) (*grpc.ClientConn, error) {

	tofndServerAddress := host + ":" + port
	logger.Info(fmt.Sprintf("initiate connection to tofnd gRPC server: %s", tofndServerAddress))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return grpc.DialContext(ctx, tofndServerAddress, grpc.WithInsecure(), grpc.WithBlock())
}

// NewMgr returns a new tss manager instance
func NewMgr(c *tmclient.HTTP, client rpc.Client, cliCtx sdkClient.Context, timeout time.Duration, principalAddr string, broadcaster *broadcast.CosmosClient, logger log.Logger, cdc *codec.LegacyAmino) *Mgr {
	return &Mgr{
		tmClient: c,
		client: client,
		//	multiSigClient: multiSigClient,
		cliCtx: cliCtx,
		keygen: &sync.RWMutex{},
		//	sign:           &sync.RWMutex{},
		keygenStreams: make(map[string]*LockableStream),
		//	signStreams:    make(map[string]*LockableStream),
		timeoutQueue:  NewTimeoutQueue(),
		Timeout:       timeout,
		principalAddr: principalAddr,
		Logger:        logger.With("listener", "tss"),
		broadcaster:   broadcaster,
		cdc:           cdc,
	}
}

// Recover instructs tofnd to recover the node's shares given the recovery info provided
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

// func (mgr *Mgr) abortSign(sigID string) (err error) {
// 	stream, ok := mgr.getSignStream(sigID)
// 	if !ok {
// 		return nil
// 	}

// 	return abort(stream)
// }

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

func handleStream(stream Stream, cancel context.CancelFunc, logger log.Logger) (broadcast <-chan *tofnd.TrafficOut, result <-chan interface{}, err <-chan error) {
	broadcastChan := make(chan *tofnd.TrafficOut)
	// TODO: MessageOut_KeygenResult and MessageOut_SignResult should be merged into one type of message
	resChan := make(chan interface{})
	errChan := make(chan error, 2)

	// server handler https://grpc.io/docs/languages/go/basics/#bidirectional-streaming-rpc-1
	go func() {
		defer cancel()
		defer close(errChan)
		defer close(broadcastChan)
		defer close(resChan)
		defer func() {
			// close the stream on error or protocol completion
			if err := stream.CloseSend(); err != nil {
				panic("handler goroutine: failure to CloseSend stream")
			}
		}()

		for {
			msgOneof, err := stream.Recv() // blocking
			if err == io.EOF {             // output stream closed by server
				logger.Debug("handler goroutine: gRPC stream closed by server")
				return
			}
			if err != nil {
				panic("handler goroutine: failure to receive msg from gRPC server stream")
				return
				//break
			}

			switch msg := msgOneof.GetData().(type) {
			case *tofnd.MessageOut_Traffic:
				broadcastChan <- msg.Traffic
			case *tofnd.MessageOut_KeygenResult_:
				resChan <- msg.KeygenResult
				return
			// case *tofnd.MessageOut_SignResult_:
			// 	resChan <- msg.SignResult
			// 	return
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
fmt.Println("processig")
//fmt.Println(e[4])
if len(e.Attributes) == 0 {
	return "","",nil,1000000000000
}
	innerMsg := e.Attributes[0].Value
	
	//fmt.Println(innerMsg)
	if len(e.Attributes) < 3 {
		return "","",nil,1000000000000
	}
	indexS := e.Attributes[2].Value
	index, err := strconv.ParseUint(string(indexS), 10, 64)
	//fmt.Println([]byte(innerMsg))
	// tx := e.(tmtypes.EventDataTx).Tx
	if err != nil {
		fmt.Println("processig error")
		return "","", nil , 1000000000000
	}
	// tx_slice := (tx[39:246])

	msg := new(dkgnet.MsgRefundMsgRequest)

	err = msg.Unmarshal(innerMsg)
	if err != nil {
		fmt.Println("processig error")
		return "","", nil , 1000000000000
	}
	msgVal := new(tss.ProcessKeygenTrafficRequest)
	// for i := 0; i < len([]byte(innerMsg)); i++ {
	// 	fmt.Println(i)
	// 	msgVal.Unmarshal([]byte(innerMsg)[i:])
	// 	fmt.Println(msgVal)
	// }

	err = msgVal.Unmarshal(msg.InnerMessage.Value)
	if err != nil {
		fmt.Println("processig error")
		return "","", nil , 1000000000000
	}
	//msgVal.Payload.IsBroadcast = true
	//fmt.Println((innerMsg))
	//  msgVal2 := new(tss.ProcessKeygenTrafficRequest)

	// msgVal2.Unmarshal([]byte(msgVal.SessionID))
	// parsers := []*parse.AttributeParser{
	// 	{Key: tss.AttributeKeySessionID, Map: parse.IdentityMap},
	// 	{Key: sdk.AttributeKeySender, Map: parse.IdentityMap},
	// 	{Key: tss.AttributeKeyPayload, Map: func(s string) (interface{}, error) {
	// 		cdc.MustUnmarshalJSON([]byte(s), &payload)
	// 		return payload, nil
	// 	}},
	// }

	// results, err := parse.Parse(attributes, parsers)
	// if err != nil {
	// 	panic(err)
	// }
		fmt.Println(msgVal.Payload.IsBroadcast, round, " ******************************************")
	return msgVal.SessionID, msgVal.Sender.String(), msgVal.Payload, index
	// return "","",nil
}
func parseMsgParams(e []types.Event) (sessionID string, from string, payload *tofnd.TrafficOut, index uint64) {
	//fmt.Println("here")
//fmt.Println(e[4])

	innerMsg := e[4].Attributes[0].Value
	
	//fmt.Println(innerMsg)
	indexS := e[4].Attributes[2].Value
	index, _ = strconv.ParseUint(string(indexS), 10, 64)
	//fmt.Println([]byte(innerMsg))
	// tx := e.(tmtypes.EventDataTx).Tx

	// tx_slice := (tx[39:246])

	msg := new(dkgnet.MsgRefundMsgRequest)

	msg.Unmarshal(innerMsg)
	
	msgVal := new(tss.ProcessKeygenTrafficRequest)
	// for i := 0; i < len([]byte(innerMsg)); i++ {
	// 	fmt.Println(i)
	// 	msgVal.Unmarshal([]byte(innerMsg)[i:])
	// 	fmt.Println(msgVal)
	// }

	msgVal.Unmarshal(msg.InnerMessage.Value)
	//msgVal.Payload.IsBroadcast = true
	//fmt.Println((innerMsg))
	//  msgVal2 := new(tss.ProcessKeygenTrafficRequest)

	// msgVal2.Unmarshal([]byte(msgVal.SessionID))
	// parsers := []*parse.AttributeParser{
	// 	{Key: tss.AttributeKeySessionID, Map: parse.IdentityMap},
	// 	{Key: sdk.AttributeKeySender, Map: parse.IdentityMap},
	// 	{Key: tss.AttributeKeyPayload, Map: func(s string) (interface{}, error) {
	// 		cdc.MustUnmarshalJSON([]byte(s), &payload)
	// 		return payload, nil
	// 	}},
	// }

	// results, err := parse.Parse(attributes, parsers)
	// if err != nil {
	// 	panic(err)
	// }

	return msgVal.SessionID, msgVal.Sender.String(), msgVal.Payload, index
	// return "","",nil
}
func parseMsgParamsDispute(e KeygenEvent) (sessionID string, from string, payload *tofnd.TrafficOut, id uint64) {

	//panic(e)
	fmt.Println(e)
	if len(e.Attributes) < 4 {
		//fmt.Println("attributes less than 4", e)
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
	value, err := strconv.Atoi(innerMsg) // Convert string to int
	if err != nil {
		fmt.Println("Conversion failed:", err)
		return
	}
	b := []byte{byte(value)} 
	id = uint64(i)
	
	fmt.Println("innermsg : ",b, innerMsg)
	if faultersMap[innerMsg]{
		return
	}
	faultersMap[innerMsg] = true
	tofndT := tofnd.TrafficOut{ToPartyUid: "", Payload: b , IsBroadcast: true}
	return keyId, sdk.AccAddress([]byte(from)).String(), &tofndT, id
	// return "","",nil
}
func parseMsgParamsDisputeOne(e []types.Event) (sessionID string, from string, payload *tofnd.TrafficOut, id uint64) {

	//panic(e)
	if len(e[0].Attributes) < 4 {
		//fmt.Println("attributes less than 4", e)
		return 
	}
	innerMsg := e[0].Attributes[0].Value
	sessionID = string(e[0].Attributes[1].Value)
	from = string(e[0].Attributes[2].Value)
	idString := string(e[0].Attributes[3].Value)

	i, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Oops, an error occurred:", err)
		return
	}
	id = uint64(i)
	tofndT := tofnd.TrafficOut{ToPartyUid: "", Payload: []byte(innerMsg), IsBroadcast: true}
	return sessionID, sdk.AccAddress([]byte(from)).String(), &tofndT, id
	// return "","",nil
}
func prepareTrafficIn(principalAddr string, from string, sessionID string, payload *tofnd.TrafficOut, logger log.Logger) *tofnd.MessageIn {
	msgIn := &tofnd.MessageIn{
		Data: &tofnd.MessageIn_Traffic{
			Traffic: &tofnd.TrafficIn{
				Payload:      payload.Payload,
				IsBroadcast:  payload.IsBroadcast,
				FromPartyUid: from,
				RoundNum: payload.RoundNum,
			},
		},
	}

	logger.Debug(fmt.Sprintf("incoming msg to tofnd: session [%.20s] from [%.20s] to [%.20s] broadcast [%t] me [%.20s]",
		sessionID, from, payload.ToPartyUid, payload.IsBroadcast, principalAddr))

	return msgIn
}
