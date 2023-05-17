package tss

import (
	"context"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	//"encoding/json"
	// "crypto/sha256"
	"fmt"
	"sort"

	//"strconv"

	//tmEvents "github.com/axelarnetwork/tm-events/events"
	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmtypes "github.com/tendermint/tendermint/types"

	//sdkFlags "github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	//	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcaster/types"
	//"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/parse"
	//	"github.com/axelarnetwork/axelar-core/utils"
	axelarnet "github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	//tssexported "github.com/axelarnetwork/axelar-core/x/tss/exported"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	tss "github.com/axelarnetwork/axelar-core/x/tss/types"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
)



// ProcessKeygenStart starts the communication with the keygen protocol
func (mgr *Mgr) ProcessKeygenStart(e tmtypes.TMEventData) error {
    keyID, threshold, participants, timeout, height, err := parseKeygenStartParams(e)
	if err != nil {
		return err
	}
fmt.Println("this")
index := -1
var list []string
for i := 0; i < len(participants); i++ {
	list = append(list, sdk.AccAddress([]byte(participants[i])).String())
	if mgr.principalAddr == participants[i] {
		fmt.Println("here")
		index = i
	}
}
fmt.Println("this2")
	// argAddr := sdk.AccAddress([]byte(validator1))
	// argAddr2 := sdk.AccAddress([]byte(validator2))
	//me := sdk.AccAddress([]byte(mgr.principalAddr))
	
	
	//participants := []string{argAddr.String(), argAddr2.String()}
	return mgr.thresholdKeygenStart(height, keyID, timeout, threshold, index, list)
	// case tssexported.Multisig.SimpleString():
	// 	return mgr.multiSigKeygenStart(keyID, participantShareCounts[myIndex])

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
	session := mgr.timeoutQueue.Enqueue(keyID, height +timeout)
rand.Seed(time.Now().UnixNano())

	// participants = []string{"rr", "tt"}
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
func (mgr *Mgr) ProcessKeygenMsg(e tmtypes.TMEventData) error {
	//fmt.Println("process")
	keyID, from, payload := parseMsgParams(e)
	//fmt.Println(keyID)
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

func parseKeygenStartParams(e tmtypes.TMEventData) (keyID string, threshold uint32, participants []string, timeout int64, height int64, err error) {
	height = e.(tmtypes.EventDataTx).Height
	dd := e.(tmtypes.EventDataTx).Tx
	//fmt.Println(string(dd))
	ddd := (dd[29:194])
	msg:= new(axelarnet.MsgStartKeygen)
	err = msg.Unmarshal(ddd)
// for i := 0; i < len(dd); i++ {
// 	for j := i + 1; j < len(dd); j++ {
// 		ddd := (dd[i:j])
// 	msg:= new(axelarnet.MsgStartKeygen)

// 	err = msg.Unmarshal(ddd)
// 		fmt.Println(j)
// 		fmt.Println(i)
// 	fmt.Println(msg)
// 	}}
//fmt.Println(msg.Participants)
var participant_list []string
err = json.Unmarshal([]byte(msg.Participants), &participant_list)
if err != nil {
	panic(err)
}
fmt.Println("here")
t, err := strconv.ParseUint(msg.Threshold,10,32)
timeout, err = strconv.ParseInt(msg.Timeout,10,64)
fmt.Println(msg.KeyID)
fmt.Println("here2")


return msg.KeyID,uint32(t),participant_list , timeout, height, nil

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
			PartyUids: participants,
			//PartyShareCounts: participantShareCounts,
			MyPartyIndex: myIndex,
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

		mgr.Logger.Info(fmt.Sprintf("outgoing keygen msg: key [%.20s] from me [%.20s] to [%.20s] broadcast [%t]\n",
			keyID, mgr.principalAddr, msg.ToPartyUid, msg.IsBroadcast))
		// sender is set by broadcaster
		argAddr := sdk.AccAddress([]byte(mgr.principalAddr))
		
		tssMsg := &tss.ProcessKeygenTrafficRequest{Sender: argAddr, SessionID: keyID, Payload: msg}

		refundableMsg := axelarnet.NewRefundMsgRequest(mgr.principalAddr,argAddr, tssMsg)
		//m,_:=tssMsg.Marshal()
		//fmt.Println(refundableMsg.Sender)
		//fmt.Println(tssMsg.Marshal())
		resp, err := mgr.broadcaster.BroadcastTx(refundableMsg, false)
		fmt.Println(resp)
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
		if res.Data.GetGroupRecoverInfo() == nil {
			return fmt.Errorf("group recovery data missing from the result")
		}
		if res.Data.GetPrivateRecoverInfo() == nil {
			return fmt.Errorf("private recovery data missing from the result")
		}

		btcecPK, err := btcec.ParsePubKey(res.Data.GetPubKey(), btcec.S256())
		if err != nil {
			return sdkerrors.Wrap(err, "failure to deserialize pubkey")
		}

		mgr.Logger.Info(fmt.Sprintf("handler goroutine: received pubkey from server! [%v]", btcecPK.ToECDSA()))
	default:
		return fmt.Errorf("invalid data type")
	}

	pollKey := voting.NewPollKey(tss.ModuleName, keyID)
	vote := &tss.VotePubKeyRequest{Sender: mgr.cliCtx.FromAddress, PollKey: pollKey, Result: result}
	_ = axelarnet.NewRefundMsgRequest(mgr.principalAddr,mgr.cliCtx.FromAddress, vote)

	//_, err := mgr.broadcaster.Broadcast(mgr.cliCtx.WithBroadcastMode(sdkFlags.BroadcastBlock), refundableMsg)
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
