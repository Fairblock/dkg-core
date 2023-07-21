package tss

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	//"os"

	//"os"

	"math/rand"
	"strconv"
	"time"

	"fmt"
	"sort"

	//tss2 "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bls "github.com/drand/kyber-bls12381"
	"github.com/tendermint/tendermint/abci/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"

	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tss "github.com/fairblock/dkg-core/x/tss/types"
)
type MPK struct{
	Mpk []byte
	Id string
}
func SetMpk(mpk []byte, id string){
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
	//CReal []byte `json:"cReal"`
}
type P2pSad struct {
	VssComplaint []ShareInfoDispute `json:"vss_complaint"`
}
var mpkFinal = MPK{}
var round = 0
var blocks = 20
var received = 0
var indices = []int{}
var numOfP = 0
var messageBuff = map[int]types.Event{}
func findMissingNumbers(numbers []int, n int) []int {
	present := make(map[int]bool)

	// Mark the numbers present in the list
	for _, num := range numbers {
		present[num] = true
	}

	// Find the missing numbers from 1 to n
	missing := make([]int, 0)
	for i := 1; i <= n; i++ {
		if !present[i] {
			missing = append(missing, i)
		}
	}

	return missing
}
func (mgr *Mgr) CheckTimeout(e types.Event) error {

	//mgr.currentHeight = height
	// if mgr.startHeight > 0 {

		// if height > mgr.startHeight+blocks {
			if mgr.keyId == string(e.Attributes[1].Value) {
			//fmt.Println("height and end of era: ", height,mgr.startHeight+blocks)
			_, ok := mgr.getKeygenStream(mgr.keyId)
			if ok {
				// mgr.startHeight = mgr.startHeight + blocks
				if string(e.Attributes[0].Value) == "0" {

					if len(indices) < numOfP {
						fmt.Println("round 1 missing")
						fmt.Println(indices)
						missing := findMissingNumbers(indices, numOfP)

						for i := 0; i < len(missing); i++ {
							mgr.findMissing(uint64(missing[i]))
						}
						// if len(indices) < numOfP {
						// 	msg := dkgnet.MsgTimeout{Creator: mgr.principalAddr, Round: strconv.FormatUint(uint64(round)+1, 10)}
						// 	_, err := mgr.broadcaster.BroadcastTx(&msg, false)

						// 	if err != nil {
						// 		return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
						// 	}
						// }
					}

				}
				if string(e.Attributes[0].Value) == "1" {

					if len(indices) < numOfP*(numOfP+1) {
						fmt.Println("round 2 missing")
						fmt.Println(indices)
						missing := findMissingNumbers(indices, numOfP*(numOfP+1))
						for i := 0; i < len(missing); i++ {
							mgr.findMissing(uint64(missing[i]))
						}
						// if len(indices) < numOfP*(numOfP+1) {
						// 	msg := dkgnet.MsgTimeout{Creator: mgr.principalAddr, Round: strconv.FormatUint(uint64(round)+1, 10)}
						// 	_, err := mgr.broadcaster.BroadcastTx(&msg, false)

						// 	if err != nil {
						// 		return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
						// 	}
						// }
					}
					fmt.Println("after: ",len(indices))
				}
				r, _ := strconv.Atoi(string(e.Attributes[0].Value))
				round = r + 1
				//fmt.Println(round)

				mgr.ProcessTimeout()
			}
		}
		//}

//	}
	return nil
}

// ProcessKeygenStart starts the communication with the keygen protocol
func (mgr *Mgr) ProcessKeygenStart(e []EventMsg, height int64) error {
	mgr.startHeight = int(height)
	keyID, threshold, participants, timeout, err := parseKeygenStartParams(e)
	blocks = int(timeout)
	mgr.keyId = keyID
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
	return mgr.thresholdKeygenStart(height, keyID, timeout, threshold, index, list)

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
func (mgr *Mgr) thresholdKeygenStart(height int64, keyID string, timeout int64, threshold uint32, myIndex int, participants []string) error {
	done := false
	session := mgr.timeoutQueue.Enqueue(keyID, height+timeout)
	rand.Seed(time.Now().UnixNano())

	stream, cancel, err := mgr.startKeygen(randSeq(35), threshold, uint32(myIndex), participants)
	if err != nil {
		return err
	}
	mgr.setKeygenStream(keyID, stream)

	// use error channel to coordinate errors during communication with sign protocol
	errChan := make(chan error, 100)
	intermediateMsgs, result, streamErrChan := handleStream(stream, cancel, mgr.Logger)
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

// ProcessKeygenMsg forwards blockchain messages to the keygen protocol
func (mgr *Mgr) ProcessKeygenMsg(e []types.Event, h int64) error {

	received = received + 1
	for i := 4; i < len(e); i++ {
	keyID, from, payload, index := parseMsgParamsOne(e[i])
	if keyID != ""{
		//fmt.Println(i, index, " ----> loop")
	indices = append(indices, int(index))

	msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)

	stream, ok := mgr.getKeygenStream(keyID)
	if !ok {
		mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
		return nil
	}

	if err := stream.Send(msgIn); err != nil {
		return sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server")
	}}
	}
	
	return nil
}
func (mgr *Mgr) findMissing(index uint64) {
	event, exist := messageBuff[int(index)]
	if  exist {
		
		received = int(index)
		keyID, from, payload, i := parseMsgParamsOne(event)
		
		// fmt.Println("---------------------------------------------------------------")
		// fmt.Println(e)
		// fmt.Println("---------------------------------------------------------------")
		//keyID, from, payload, i := parseMsgParams(e)

		//fmt.Println("fetched idex: ", i, mgr.me)
		msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)

		stream, ok := mgr.getKeygenStream(keyID)
		if !ok {
			mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))

		}
	
		if err := stream.Send(msgIn); err != nil {
			mgr.Logger.Info("failure to send incoming msg to gRPC server")
		}
		indices = append(indices, int(i))
		fmt.Println("msg and len: ", payload, len(indices))
		return
	}
	received = int(index)
	query := "keygen.index = " + strconv.FormatUint(index, 10)
	//fmt.Println(query, mgr.me)
	page := 1
	limit := 1
	orderBy := "desc"
	result, err := mgr.tmClient.TxSearch(context.Background(), query, false, &page, &limit, orderBy)
	if err != nil {

		mgr.Logger.Error(err.Error())
	}

	for _, tx := range result.Txs {

		e := tx.TxResult.Events
		for j := 4; j < len(e); j++ {
			keyID, from, payload, i := parseMsgParamsOne(e[j])
			if i == index {
		// fmt.Println("---------------------------------------------------------------")
		// fmt.Println(e)
		// fmt.Println("---------------------------------------------------------------")
		//keyID, from, payload, i := parseMsgParams(e)

	//	fmt.Println("fetched idex: ", i, mgr.me)
		msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
		

		stream, ok := mgr.getKeygenStream(keyID)
		if !ok {
			mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))

		}
		
		if err := stream.Send(msgIn); err != nil {
			mgr.Logger.Info("failure to send incoming msg to gRPC server")
		}
		indices = append(indices, int(i))
		fmt.Println("msg and len: ", payload, len(indices))

	}
	if index != i {
		if i != 1000000000000 {
		messageBuff[int(i)] = e[j]
	}}
	
}}}

func (mgr *Mgr) ProcessKeygenMsgDispute(e []KeygenEvent) error {

	keyID, from, payload := parseMsgParamsDispute(e)

	msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)

	stream, ok := mgr.getKeygenStream(keyID)
	if !ok {
		mgr.Logger.Info(fmt.Sprintf("no keygen session with id %s. This process does not participate", keyID))
		return nil
	}

	if err := stream.Send(msgIn); err != nil {
		return sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server")
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
		return sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server")
	}
	return nil
}
func parseKeygenStartParams(e []EventMsg) (string, uint32, []string, int64, error) {
	keyID := e[0].Events[0].Attributes[0].Value
	t := e[0].Events[0].Attributes[1].Value
	participants := e[0].Events[0].Attributes[2].Value
	timeout := e[0].Events[0].Attributes[3].Value

	var participant_list []string
	err := json.Unmarshal([]byte(participants), &participant_list)
	if err != nil {
		panic(err)
	}

	threshold_uint64, err := strconv.ParseUint(t, 10, 32)
	timeout_int, err := strconv.ParseInt(timeout, 10, 64)

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
		//fmt.Println("message round and round: ", num, round)
		if num != round {

			for {
				time.Sleep(150 * time.Millisecond)

				if num == round {

					break
				}
			}

		}
		rand.Seed(time.Now().UnixNano())

		// minDelay := 1  // Minimum delay in seconds
		// maxDelay := 10 // Maximum delay in seconds
		// delay := rand.Intn(maxDelay-minDelay+1) + minDelay
		//time.Sleep(150 * time.Millisecond)

		argAddr := sdk.AccAddress([]byte(mgr.principalAddr))
		// sender is set by broadcaster
		if msg.ToPartyUid == "r3" {
			var p2pSad P2pSad

			err := json.Unmarshal((msg.Payload[9 : len(msg.Payload)-1]), &p2pSad)

			if err != nil {
				fmt.Println("Error:", err)

			}

			for i := 0; i < len(p2pSad.VssComplaint); i++ {
				complaint := p2pSad.VssComplaint[i]
				byteSlice := make([]byte, 32)

				binary.BigEndian.PutUint32(byteSlice, uint32(complaint.Share.Index))
				fmt.Println("accuser: ", complaint.AccuserId)
				fmt.Println("accusee: ", complaint.FaulterId)
				msgR3 := dkgnet.MsgFileDispute{Creator: mgr.principalAddr, Dispute: &dkgnet.Dispute{AddressOfAccuser: complaint.Accuser, AddressOfAccusee: complaint.Faulter, Share: &dkgnet.Share{Value: complaint.Share.Scalar, Index: byteSlice, Id: uint64(complaint.Share.Index)}, Commit: &dkgnet.Commit{Commitments: complaint.Commit}, Kij: complaint.Kij, CZkProof: complaint.Proof[0][:], RZkProof: complaint.Proof[1][:], Id: 1, AccuserId: uint64(complaint.AccuserId), FaulterId: uint64(complaint.FaulterId), CReal: complaint.Proof[2][:]}, IdOfAccuser: uint64(complaint.Share.Index), KeyId: keyID}

				_, err := mgr.broadcaster.BroadcastTx(&msgR3, false)

				if err != nil {
					return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
				}
			}

			return nil
		}

		tssMsg := &tss.ProcessKeygenTrafficRequest{Sender: argAddr, SessionID: keyID, Payload: msg}

		refundableMsg := dkgnet.NewMsgRefundMsgRequest(mgr.principalAddr, argAddr, tssMsg)
		if msg.RoundNum == "1"{
		// messageBuff = append(messageBuff, refundableMsg)
		// for {
		// 	fmt.Println(len(messageBuff))
		// 	if len(messageBuff) == numOfP  {
		// 		break
		// 	}
		// }
		//fmt.Println("batch sending ========================================================================")
		// div := numOfP/10
		// for i := 0; i < (div+1); i++ {
		// 	if i == (div){
				// batch := []sdk.Msg{}
				// batch = append(batch, messageBuff[i*10:])
			// 	_, err := mgr.broadcaster.BroadcastTxs(messageBuff[i*10:], false)
		
			// if err != nil {
			// 	//fmt.Println("this", tssMsg)
			// 	return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
			// }
			// 	break
			// }
			// // batch := []sdk.Msg{}
			// // batch = append(batch, messageBuff[i*10:i*10+10]) messageBuff[i*10:i*10+10]
			
			_, err := mgr.broadcaster.BroadcastTxs(refundableMsg, false, numOfP)
		
			if err != nil {
				//fmt.Println("this", tssMsg)
				return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
			}
		}
		
	
	
	if msg.RoundNum != "1"{
		//fmt.Println("single sending ========================================================================")
		_, err := mgr.broadcaster.BroadcastTx(refundableMsg, false)

		if err != nil {
			//fmt.Println("this", tssMsg)
			return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")

	}}
	
}

return nil}

func (mgr *Mgr) handleKeygenResult(keyID string, resultChan <-chan interface{}) error {
	// Delete the reference to the keygen stream with keyID because entering this function means the tss protocol has completed
	defer func() {
		mgr.keygen.Lock()
		defer mgr.keygen.Unlock()
		delete(mgr.keygenStreams, keyID)
	}()

	r, ok := <-resultChan
	if !ok {
		return fmt.Errorf("failed to receive keygen result, channel was closed by the server")
	}

	// get result. Result will be implicity checked by Validate() during Braodcast(), so no checks are needed here
	result, ok := r.(*tofnd.MessageOut_KeygenResult)
	if !ok {
		return fmt.Errorf("failed to receive keygen result, received unexpected type %T", r)
	}

	mgr.Logger.Debug(fmt.Sprintf("handler goroutine: received keygen result for %s [%+v]", keyID, result))

	switch res := result.GetKeygenResultData().(type) {
	case *tofnd.MessageOut_KeygenResult_Criminals:
		// prepare criminals for Validate()
		// criminals have to be sorted in ascending order
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
		mgr.Logger.Info(fmt.Sprintf("share bytes: ", shareB))
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
		msg := dkgnet.MsgKeygenResult{Creator: mgr.principalAddr, Mpk: pk.String(), Commitment: commitment.String()}
		resp, err := mgr.broadcaster.BroadcastTx(&msg, false)

		if err != nil {
			return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
		}
		mgr.Logger.Info(fmt.Sprintf("mpk bytes: ", pkBytes))
		mgr.Logger.Info(fmt.Sprintf("handler goroutine: submitted mpk and commitment: ", resp))
		// for{
		// 	if mpkFinal.Id != "" {
		// 		os.Exit(0)
		// 	}
		// }

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
