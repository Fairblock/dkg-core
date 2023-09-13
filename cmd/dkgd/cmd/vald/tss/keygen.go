package tss

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bls "github.com/drand/kyber-bls12381"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"
	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tss "github.com/fairblock/dkg-core/x/tss/types"
	"github.com/tendermint/tendermint/abci/types"
)

type MPK struct {
	Mpk []byte
	Id  string
}

func SetMpk(mpk []byte, id string) {
	mpkFinal.Mpk = mpk
	mpkFinal.Id = id
}

type Share struct {
	Scalar []byte `json:"scalar"`
	Index  int    `json:"index"`
}
type KeyShareRecoveryInfo struct {
	X_i_ciphertext []byte `json:"x_i_ciphertext"`
}
type ShareInfoDispute struct {
	Share     Share       `json:"share"`
	Kij       []byte      `json:"kij"`
	Proof     [3][32]byte `json:"proof"`
	Commit    [][]byte    `json:"commit"`
	Faulter   []byte      `json:"faulter"`
	Accuser   []byte      `json:"accuser"`
	AccuserId uint64      `json:"accuserId"`
	FaulterId uint64      `json:"faulterId"`
}
type P2pSad struct {
	VssComplaint []ShareInfoDispute `json:"vss_complaint"`
}

var mpkFinal = MPK{}
var round = 0
var blocks = 20
var received = 0
var indices = AtomicSlice{indicesList: []int{}}
var numOfP = 0
var messageBuff = map[int]types.Event{}

type AtomicSlice struct {
	indicesList []int
	mu   sync.Mutex
}
func (a *AtomicSlice) Append(item int) {
	a.mu.Lock()         // Lock the mutex
	defer a.mu.Unlock() // Ensure the mutex is unlocked when we're done

	a.indicesList = append(a.indicesList, item)
}

func findMissingNumbers(numbers []int, n int) []int {
	present := make(map[int]bool)
	for _, num := range numbers {
		present[num] = true
	}
	missing := make([]int, 0)
	for i := 1; i <= n; i++ {
		if !present[i] {
			missing = append(missing, i)
		}
	}
	return missing
}
func (mgr *Mgr) CheckTimeout(e types.Event) error {
	if mgr.keyId == string(e.Attributes[1].Value) {
		_, ok := mgr.getKeygenStream(mgr.keyId)
		if ok {
			if string(e.Attributes[0].Value) == "0" {
				if len(indices.indicesList) < numOfP {
					fmt.Println("round 1 missing")
					
					missing := findMissingNumbers(indices.indicesList, numOfP)
					for i := 0; i < len(missing); i++ {
						mgr.findMissing(uint64(missing[i]))
					}
					missing = findMissingNumbers(indices.indicesList, numOfP)
					fmt.Println(missing)
				}
			}
			if string(e.Attributes[0].Value) == "1" {
				
				if len(indices.indicesList) < numOfP*(numOfP+1) {
					fmt.Println("round 2 missing")

					missing := findMissingNumbers(indices.indicesList, numOfP*(numOfP+1))
					
					for i := 0; i < len(missing); i++ {
						mgr.findMissing(uint64(missing[i]))
					}
				}

			}
			if string(e.Attributes[0].Value) == "2" {
				for i := numOfP*(numOfP+1) + 1; i < (numOfP*numOfP*2)+1; i++ {
					skip := false
					for k := 0; k < len(indices.indicesList); k++ {
						if indices.indicesList[k] == i {
							skip = true
							break
						}
					}
					if !skip {
						found := mgr.findMissingDispute(uint64(i))
						fmt.Println(found, i)
						if !found {
							break
						}
					}
				}
			}
			r, _ := strconv.Atoi(string(e.Attributes[0].Value))
			round = r + 1
			mgr.ProcessTimeout()
		}
	}
	return nil
}
func (mgr *Mgr) findMissingDispute(index uint64) bool {
	//fmt.Println("looking for disputes...", index)
	event, exist := messageBuff[int(index)]
	if exist {
		//fmt.Println(exist)
		received = int(index)
		keyID, from, payload, _ := parseMsgParamsDisputeOne([]types.Event{event})
		if payload == nil {
			return false
		}
		if keyID != mgr.keyId {
			return false
		}
	//	fmt.Println("dispute fetched idex: ", i, mgr.me)
		msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
		time.Sleep(time.Duration(50) * time.Millisecond)
		stream, ok := mgr.getKeygenStream(keyID)
		if !ok {
			mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
		}
		if err := stream.Send(msgIn); err != nil {
			mgr.Logger.Info("failure to send incoming msg to gRPC server")
		}
		return true
	}
	received = int(index)
	query := "keygen.index = " + strconv.FormatUint(index, 10)
	page := 1
	limit := 1
	orderBy := "desc"
	result, err := mgr.tmClient.TxSearch(context.Background(), query, false, &page, &limit, orderBy)
	if err != nil {
		return false
	}
	found := false
	for _, tx := range result.Txs {
		e := tx.TxResult.Events
		//fmt.Println("fetched dispute ")
		keyID, from, payload, i := parseMsgParamsDisputeOne(e)
		//fmt.Println("fetched dispute : ", keyID, from, payload, i, index)
		if keyID != mgr.keyId {
			
			return false
		}
		if i == index {
			found = true
			time.Sleep(time.Duration(10) * time.Millisecond)
		//	fmt.Println("dispute fetched idex: ", i, mgr.me)
			msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
			stream, ok := mgr.getKeygenStream(keyID)
			if !ok {
				mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
			}
			if err := stream.Send(msgIn); err != nil {
				mgr.Logger.Info("failure to send incoming msg to gRPC server")
			}
		}
		if index != i {
			if i != 1000000000000 {
				messageBuff[int(i)] = e[4]
			}
		}
	}
	return found
}
func (mgr *Mgr) ProcessKeygenStart(e types.Event) error {
	fmt.Println("start")
	round = 0
	mpkFinal = MPK{}
	received = 0
	indices = AtomicSlice{indicesList: []int{}}
	numOfP = 0
	messageBuff = map[int]types.Event{}
	// mgr.startHeight = int(height)
	keyID, threshold, participants, timeout, err := parseKeygenStartParams(e)
	blocks = int(timeout)
	mgr.keyId = keyID
	//fmt.Println(keyID, threshold, participants, timeout, err)
	if err != nil {
		return err
	}
	index := -1
	var list []string
	for i := 0; i < len(participants); i++ {
		list = append(list, sdk.AccAddress([]byte(participants[i])).String())
		if mgr.principalAddr == participants[i] {
			index = i
		}
	}
	numOfP = len(list)
	mgr.me = index
	return mgr.thresholdKeygenStart(keyID, timeout, threshold, index, list)
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func (mgr *Mgr) thresholdKeygenStart(keyID string, timeout int64, threshold uint32, myIndex int, participants []string) error {
	done := false
	session := mgr.timeoutQueue.Enqueue(keyID, timeout)
	rand.Seed(time.Now().UnixNano())
	stream, cancel, err := mgr.startKeygen(randSeq(35), threshold, uint32(myIndex), participants)
	if err != nil {
		return err
	}
	mgr.setKeygenStream(keyID, stream)
	errChan := make(chan error, 120)
	intermediateMsgs, result, streamErrChan := handleStream(stream, cancel, mgr.Logger, mgr.me)
	go func() {
		err, ok := <-streamErrChan
		if ok {
			errChan <- err
		}
	}()
	go func() {
		err := mgr.handleIntermediateKeygenMsgs(keyID, intermediateMsgs)
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		session.WaitForTimeout()
		if done {
			return
		}
		errChan <- mgr.abortKeygen(keyID)
		mgr.Logger.Info(fmt.Sprintf("aborted keygen protocol %s due to timeout", keyID))
	}()
	go func() {
		err := mgr.handleKeygenResult(keyID, result)
		done = true
		errChan <- err
	}()
	return <-errChan
}
func (mgr *Mgr) ProcessKeygenMsg(e []types.Event, h int64) error {
	received = received + 1
	
	for i := 4; i < len(e); i++ {
		keyID, from, payload, index := parseMsgParamsOne(e[i])
		if payload != nil {
			
			if payload.RoundNum != strconv.Itoa(round) {
				fmt.Println("wrong round ", payload.RoundNum, round)
				return nil
			}
			if round == 1 {
				if index <= uint64(numOfP) {
					fmt.Println("wrong message 1")
					return nil
				}
			}
			if round == 0 {
				if index > uint64(numOfP) {
					fmt.Println("wrong message 0")
					return nil
				}
			}
			for k := 0; k < len(indices.indicesList); k++ {
				if indices.indicesList[k] == int(index) {
					fmt.Println("repeated message")
					return nil
				}
			}
			if keyID != mgr.keyId {
				fmt.Println("wrong id")
				return nil
			}
			
			indices.Append(int(index))
			//indices = append(indices, int(index))
			fmt.Println(mgr.me, index, " ----> loop")
			msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
			stream, ok := mgr.getKeygenStream(keyID)
			if !ok {
				mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
				return nil
			}
			if err := stream.Send(msgIn); err != nil {
				panic(sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server"))
			}
		}
	}
	return nil
}
func (mgr *Mgr) findMissing(index uint64) {

	event, exist := messageBuff[int(index)]
	if exist {
		fmt.Println(exist)
		received = int(index)
		keyID, from, payload, i := parseMsgParamsOne(event)
		if payload == nil {
			panic("wrong----")
		}
		if keyID != mgr.keyId {
			return 
		}
		fmt.Println("fetched idex: ", i, mgr.me)
		
		msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
		time.Sleep(time.Duration(50) * time.Millisecond)
		stream, ok := mgr.getKeygenStream(keyID)
		if !ok {
			mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
		}
		if err := stream.Send(msgIn); err != nil {
			mgr.Logger.Info("failure to send incoming msg to gRPC server")
		}
		indices.Append(int(i))
		// indices = append(indices, int(i))
		return
	}
	received = int(index)
	query := "keygen.index = " + strconv.FormatUint(index, 10)
	page := 1
	limit := 1
	orderBy := "desc"
	result, err := mgr.tmClient.TxSearch(context.Background(), query, false, &page, &limit, orderBy)
	if err != nil {
		panic(err.Error())
	}
	for _, tx := range result.Txs {
		e := tx.TxResult.Events
		for j := 4; j < len(e); j++ {
			keyID, from, payload, i := parseMsgParamsOne(e[j])
			if keyID != mgr.keyId {
				return
			}
			if i == index {
				time.Sleep(time.Duration(10) * time.Millisecond)
				fmt.Println("fetched idex: ", i, mgr.me)
				msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
				stream, ok := mgr.getKeygenStream(keyID)
				if !ok {
					mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
				}
				if err := stream.Send(msgIn); err != nil {
					mgr.Logger.Info("failure to send incoming msg to gRPC server")
				}
				indices.Append(int(i))
				// indices = append(indices, int(i))
			}
			if index != i {
				if i != 1000000000000 {
					messageBuff[int(i)] = e[j]
				}
			}
		}
	}
}
func (mgr *Mgr) ProcessKeygenMsgDispute(e []KeygenEvent) error {
//	fmt.Println("ProcessKeygenMsgDispute")
	for {
		if round > 1 {
			break
		}
	}
	for i := 0; i < len(e); i++ {
		keyID, from, payload, index := parseMsgParamsDispute(e[i])
		//fmt.Println("dispute has been received")
		if index > uint64(numOfP*(numOfP+1)) {
			if keyID == mgr.keyId {
				fmt.Println("dispute: ", index)
				indices.Append(int(index))
				// indices = append(indices, int(index))
				msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
				stream, ok := mgr.getKeygenStream(keyID)
				if !ok {
					mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
					return nil
				}
				if err := stream.Send(msgIn); err != nil {
					panic(sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server"))
				}
			}
		}
	}
	return nil
}
func (mgr *Mgr) ProcessTimeout() error {
	msgIn := prepareTrafficIn(mgr.principalAddr, mgr.principalAddr, mgr.keyId, &tofnd.TrafficOut{ToPartyUid: strconv.Itoa(round - 1), Payload: []byte("timeout" + strconv.Itoa(round)), IsBroadcast: true}, mgr.Logger)
	stream, ok := mgr.getKeygenStream(mgr.keyId)
	if !ok {
		mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", mgr.keyId))
		return nil
	}
	if err := stream.Send(msgIn); err != nil {
		panic(sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server"))
	}
	return nil
}
func parseKeygenStartParams(e types.Event) (string, uint32, []string, int64, error) {
	keyID := string(e.Attributes[0].Value)
	t := string(e.Attributes[1].Value)
	participants := e.Attributes[2].Value
	timeout := string(e.Attributes[3].Value)
	var participant_list []string
	err := json.Unmarshal([]byte(participants), &participant_list)
	if err != nil {
		panic(err)
	}
	threshold_uint64, _ := strconv.ParseUint(t, 10, 32)
	timeout_int, _ := strconv.ParseInt(timeout, 10, 64)
	return keyID, uint32(threshold_uint64), participant_list, timeout_int, nil
}
func (mgr *Mgr) startKeygen(keyID string, threshold uint32, myIndex uint32, participants []string) (Stream, context.CancelFunc, error) {
	if _, ok := mgr.getKeygenStream(keyID); ok {
		return nil, nil, fmt.Errorf("keygen protocol for ID %s already in progress", keyID)
	}
	grpcCtx, cancel := context.WithTimeout(context.Background(), mgr.Timeout)
	stream, err := mgr.client.Keygen(grpcCtx)
	if err != nil {
		cancel()
		return nil, nil, sdkerrors.Wrap(err, "failed tofnd gRPC call Keygen")
	}
	keygenInit := &tofnd.MessageIn_KeygenInit{
		KeygenInit: &tofnd.KeygenInit{
			NewKeyUid:    keyID,
			Threshold:    threshold,
			MyPartyIndex: myIndex,
			PartyUids:    participants,
		},
	}
	if err := stream.Send(&tofnd.MessageIn{Data: keygenInit}); err != nil {
		cancel()
		return nil, nil, err
	}
	return stream, cancel, nil
}
func (mgr *Mgr) handleIntermediateKeygenMsgs(keyID string, intermediate <-chan *tofnd.TrafficOut) error {
	for msg := range intermediate {
		num, _ := strconv.Atoi(msg.RoundNum)
		if num != round {
			for {
				time.Sleep(150 * time.Millisecond)
				if num == round {
					break
				}
			}
		}
		rand.Seed(time.Now().UnixNano())
		argAddr := sdk.AccAddress([]byte(mgr.principalAddr))
		if msg.RoundNum == "2" {
			var p2pSad P2pSad
			c := 0
			for {
				err := json.Unmarshal((msg.Payload[c : len(msg.Payload)-1]), &p2pSad)
				if err != nil {

					c = c + 1
				}
				if err == nil {

					break
				}
			}
			r3MsgList := []dkgnet.MsgFileDispute{}
			for i := 0; i < len(p2pSad.VssComplaint); i++ {
				complaint := p2pSad.VssComplaint[i]
				byteSlice := make([]byte, 32)
				binary.BigEndian.PutUint32(byteSlice, uint32(complaint.Share.Index))

				msgR3 := dkgnet.MsgFileDispute{Creator: mgr.principalAddr, Dispute: &dkgnet.Dispute{AddressOfAccuser: complaint.Accuser, AddressOfAccusee: complaint.Faulter, Share: &dkgnet.Share{Value: complaint.Share.Scalar, Index: byteSlice, Id: uint64(complaint.Share.Index)}, Commit: &dkgnet.Commit{Commitments: complaint.Commit}, Kij: complaint.Kij, CZkProof: complaint.Proof[0][:], RZkProof: complaint.Proof[1][:], Id: 1, AccuserId: uint64(complaint.AccuserId), FaulterId: uint64(complaint.FaulterId), CReal: complaint.Proof[2][:]}, IdOfAccuser: uint64(complaint.Share.Index), KeyId: keyID}
				r3MsgList = append(r3MsgList, msgR3)
			}
			_, err := mgr.broadcaster.BroadcastTxDispute(r3MsgList, false, mgr.me)
			if err != nil {
				panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg"))
			}
			return nil
		}
		tssMsg := &tss.ProcessKeygenTrafficRequest{Sender: argAddr, SessionID: keyID, Payload: msg}
		refundableMsg := dkgnet.NewMsgRefundMsgRequest(mgr.principalAddr, argAddr, tssMsg)
		if msg.RoundNum == "1" {
			if msg.IsBroadcast {

				delay := mgr.me * 100
				time.Sleep(time.Duration(delay) * time.Millisecond)
				_, err := mgr.broadcaster.BroadcastTx(refundableMsg, false)
				if err != nil {
					panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg"))
				}
			}
			if !msg.IsBroadcast {
				_, err := mgr.broadcaster.BroadcastTxs(refundableMsg, false, numOfP, mgr.me)
				if err != nil {
					panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg"))
				}
			}
		}
		if msg.RoundNum == "0" {
			_, err := mgr.broadcaster.BroadcastTx(refundableMsg, false)
			if err != nil {
				panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg"))
			}
		}
	}
	return nil
}
func (mgr *Mgr) handleKeygenResult(keyID string, resultChan <-chan interface{}) error {
	defer func() {
		mgr.keygen.Lock()
		defer mgr.keygen.Unlock()
		delete(mgr.keygenStreams, keyID)
	}()
	r, ok := <-resultChan
	if !ok {
		return fmt.Errorf("failed to receive keygen result, channel was closed by the server")
	}
	result, ok := r.(*tofnd.MessageOut_KeygenResult)
	if !ok {
		return fmt.Errorf("failed to receive keygen result, received unexpected type %T", r)
	}
	mgr.Logger.Debug(fmt.Sprintf("handler goroutine: received keygen result for %s [%+v]", keyID, result))
	switch res := result.GetKeygenResultData().(type) {
	case *tofnd.MessageOut_KeygenResult_Criminals:
		sort.Stable(res.Criminals)
	case *tofnd.MessageOut_KeygenResult_Data:
		if res.Data.GetPubKey() == nil {
			return fmt.Errorf("public key missing from the result")
		}
		pkBytes := res.Data.GetPubKey()
		suite := bls.NewBLS12381Suite()
		pk := suite.G1().Point()
		pk.UnmarshalBinary(pkBytes)
		share := res.Data.GetPrivateRecoverInfo()
		shareB := share[4 : len(share)-1]
		index := share[len(share)-1]
		for i, j := 0, len(shareB)-1; i < j; i, j = i+1, j-1 {
			shareB[i], shareB[j] = shareB[j], shareB[i]
		}
		_ = suite
		shareS := suite.G1().Scalar().SetBytes(shareB)
		commitment := suite.G1().Point().Mul(shareS, pk)
		filenamePk := fmt.Sprintf("pk-%d.txt", index)
		filenameShare := fmt.Sprintf("share-%d.txt", index)
		err := ioutil.WriteFile(filenamePk, pkBytes, 0644)
		if err != nil {
			fmt.Printf("Failed to write to file: %s\n", err)
		}
		err = ioutil.WriteFile(filenameShare, shareB, 0644)
		if err != nil {
			fmt.Printf("Failed to write to file: %s\n", err)
		}
		me := strconv.Itoa(mgr.me)
		msg := dkgnet.MsgKeygenResult{Creator: mgr.principalAddr, MyIndex: me, Commitment: commitment.String()}
		resp, err := mgr.broadcaster.BroadcastTx(&msg, false)
		if err != nil {
			panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg"))
		}
		msgr := dkgnet.MsgRegisterValidator{Creator: mgr.principalAddr, Participation: true}

		_, err = mgr.broadcaster.BroadcastTx(&msgr, false)
		if err != nil {
			panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing register msg"))
		}
		mgr.Logger.Info(fmt.Sprintf("mpk bytes: ", pkBytes, "me: ", mgr.me, "resp:", resp))
	default:
		return fmt.Errorf("invalid data type")
	}
	return nil
}
func (mgr *Mgr) getKeygenStream(keyID string) (Stream, bool) {
	mgr.keygen.RLock()
	defer mgr.keygen.RUnlock()
	stream, ok := mgr.keygenStreams[keyID]
	return stream, ok
}
func (mgr *Mgr) setKeygenStream(keyID string, stream Stream) {
	mgr.keygen.Lock()
	defer mgr.keygen.Unlock()
	mgr.keygenStreams[keyID] = NewLockableStream(stream)
}


