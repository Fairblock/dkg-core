package tss

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"os"
	//"math/big"
	"math/rand"
	"strconv"
	"time"

	//	"github.com/drand/kyber"
	//bls "github.com/drand/kyber-bls12381"
	//"encoding/json"
	// "crypto/sha256"
	"fmt"
	"sort"

	//"strconv"

	//"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bls "github.com/drand/kyber-bls12381"
	"github.com/tendermint/tendermint/abci/types"

	//tmtypes "github.com/tendermint/tendermint/types"

	//sdkFlags "github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	//	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster/types"
	//"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/parse"
	//	"github.com/fairblock/dkg-core/utils"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"

	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tss "github.com/fairblock/dkg-core/x/tss/types"
)

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

var round = 0
var blocks = 150
var received = 0
func (mgr *Mgr) CheckTimeout(height int) error {
	
	mgr.currentHeight = height
	if mgr.startHeight > 0 {
		
		if height > mgr.startHeight+blocks {
			//fmt.Println("height and end of era: ", height,mgr.startHeight+blocks)
			_, ok := mgr.getKeygenStream(mgr.keyId)
			if ok {
				mgr.startHeight = mgr.startHeight + blocks
				round = round + 1
				fmt.Println(round)
				mgr.ProcessTimeout()
			}

		}
		
	}
	return nil
}

// ProcessKeygenStart starts the communication with the keygen protocol
func (mgr *Mgr) ProcessKeygenStart(e []EventMsg, height int64) error {
	mgr.startHeight = int(height)
	keyID, threshold, participants, timeout, err := parseKeygenStartParams(e)
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
	errChan := make(chan error, 4)
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
	received = received+1
	fmt.Println("received message: me:%s num received:%s ",mgr.me, received)
	keyID, from, payload := parseMsgParams(e)
	fmt.Println("received:%s id:%s",received,keyID)
	msgIn := prepareTrafficIn(mgr.principalAddr, from, keyID, payload, mgr.Logger)
	
	fmt.Println("received:%s round:%s", received,msgIn.Data)
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

	msgIn := prepareTrafficIn(mgr.principalAddr, mgr.principalAddr, mgr.keyId, &tofnd.TrafficOut{ToPartyUid: strconv.Itoa(round-1), Payload: []byte("timeout"+strconv.Itoa(round)), IsBroadcast: true}, mgr.Logger)

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
			NewKeyUid: keyID,
			Threshold: threshold,
			MyPartyIndex: myIndex,
			PartyUids: participants,
		},
	}

fmt.Println(keygenInit)

	if err := stream.Send(&tofnd.MessageIn{Data: keygenInit}); err != nil {
		cancel()
		return nil, nil, err
	}

	return stream, cancel, nil
}

func (mgr *Mgr) handleIntermediateKeygenMsgs(keyID string, intermediate <-chan *tofnd.TrafficOut) error {
	
	
	for msg := range intermediate {
		
		num, _ := strconv.Atoi(msg.RoundNum)
		fmt.Println("message round and round: ", num, round)
		if num != round {
			//fmt.Println("waiting round:", msg.RoundNum)
				for {
					time.Sleep(150 * time.Millisecond)
					fmt.Println("waiting:", num, round)
				if num == round {
					
		// if mgr.currentHeight >= mgr.startHeight+blocks {
		// 	if mgr.currentHeight > mgr.startHeight+(2*blocks) {
		// 		panic("timeout")
		// 	}
		// 	fmt.Println("new round:", msg.RoundNum)
			
		// 	round = msg.RoundNum
			break}}

		
	}
	rand.Seed(time.Now().UnixNano())

	// Generate a random delay between 0 and 1000 milliseconds
	//delay := rand.Intn(11)

	//fmt.Printf("Waiting for %d milliseconds...\n", delay)
	minDelay := 1 // Minimum delay in seconds
	maxDelay := 50 // Maximum delay in seconds
	delay := rand.Intn(maxDelay-minDelay+1) + minDelay
	time.Sleep(time.Duration(delay)* 200 * time.Millisecond)
fmt.Println(delay)
		// mgr.Logger.Info(fmt.Sprintf("outgoing keygen msg: key [%.20s] from me [%.20s] to [%.20s] broadcast [%t]\n",
		// 	keyID, mgr.principalAddr, msg.ToPartyUid, msg.IsBroadcast))
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
				byteSlice := make([]byte, 32) // Assuming you want to convert to a 4-byte slice

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

		_, err := mgr.broadcaster.BroadcastTx(refundableMsg, false)

		if err != nil {
			return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing keygen msg")
		}

	}
	return nil
}

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
		//fmt.Println(share)
		shareB := share[4 : len(share)-2]
		index := share[len(share)-1]

		for i, j := 0, len(shareB)-1; i < j; i, j = i+1, j-1 {
			shareB[i], shareB[j] = shareB[j], shareB[i]
		}

		_ = suite
		//shareS := suite.G1().Scalar().SetBytes(shareB)

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
		os.Exit(0)
		// mgr.Logger.Info(fmt.Sprintf("handler goroutine: received pubkey bytes from server! : ", pkBytes))
		// mgr.Logger.Info(fmt.Sprintf("handler goroutine: received pubkey from server! : ", pk.String()))
		// mgr.Logger.Info(fmt.Sprintf("handler goroutine: received share from server! : ", shareB))
		// mgr.Logger.Info(fmt.Sprintf("handler goroutine: received share from server! : ", shareS.String()))
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
