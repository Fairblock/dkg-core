package broadcaster

import (
	"context"
	"encoding/hex"
	"fmt"

	//"fmt"

	//"fmt"
	"log"
	"strings"
	"time"

	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"

	//"cosmossdk.io/math"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/fairblock/dkg-core/app"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	defaultGasAdjustment = 1.5
	defaultGasLimit      = 300000
)

type CosmosClient struct {
	authClient authtypes.QueryClient
	txClient   tx.ServiceClient
	grpcConn   grpc.ClientConn

	privateKey secp256k1.PrivKey
	publicKey  cryptotypes.PubKey
	account    authtypes.BaseAccount
	accAddress cosmostypes.AccAddress
	chainID    string
}
var messageBuff = []cosmostypes.Msg{}
func PrivateKeyToAccAddress(privateKeyHex string) (cosmostypes.AccAddress, error) {
	keyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, err
	}

	privateKey := secp256k1.PrivKey{Key: keyBytes}

	return cosmostypes.AccAddress(privateKey.PubKey().Address()), nil
}

func NewCosmosClient(
	endpoint string,
	privateKeyHex string,
	chainID string,
) (*CosmosClient, error) {
	grpcConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	authClient := authtypes.NewQueryClient(grpcConn)

	keyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, err
	}

	privateKey := secp256k1.PrivKey{Key: keyBytes}
	pubKey := privateKey.PubKey()
	address := pubKey.Address()

	accAddr := cosmostypes.AccAddress(address)
	addr := accAddr.String()

	var baseAccount authtypes.BaseAccount

	resp, err := authClient.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{Address: addr},
	)

	if err != nil {
		log.Println(cosmostypes.AccAddress(address).String())
		return nil, err
	}

	err = baseAccount.Unmarshal(resp.Account.Value)
	if err != nil {
		return nil, err
	}

	return &CosmosClient{

		authClient: authClient,
		txClient:   tx.NewServiceClient(grpcConn),
		grpcConn:   *grpcConn,
		privateKey: privateKey,
		account:    baseAccount,
		accAddress: accAddr,
		publicKey:  pubKey,
		chainID:    chainID,
	}, nil
}

func (c *CosmosClient) GetAddress() string {
	return c.account.Address
}
func (c *CosmosClient) GetPK() string {
	return c.publicKey.String()
}
func (c *CosmosClient) GetAccAddress() cosmostypes.AccAddress {
	return c.accAddress
}

func (c *CosmosClient) handleBroadcastResult(resp *cosmostypes.TxResponse, err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("make sure that your account has enough balance")
		}
		return err
	}

	if resp.Code > 0 {
		return errors.Errorf("error code: '%d' msg: '%s'", resp.Code, resp.RawLog)
	}
	return nil
}
func (c *CosmosClient) BroadcastTxDispute(msgs []dkgnet.MsgFileDispute, adjustGas bool, id int) (*cosmostypes.TxResponse, error) {
	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&: ",msg.Creator)
	msgsList := []cosmostypes.Msg{}
	numOfP := len(msgs)
	for i := 0; i < numOfP; i++ {
		msgsList = append(msgsList, &msgs[i])
	}
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@: ", "msgsList")
	//fmt.Println("=========================================================", (messageBuff))
	n := 1
	div := (numOfP)/n
	add := (numOfP) % n
	if add != 0 {
		add = 1
	}
	//fmt.Println(div)
	delay := id * 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
		for i := 0; i < (div+add); i++ {
			
			if i == (div){
				
				txBytes, err := c.signTxMsgs(msgsList[i*n:], false)
				//fmt.Println("txBytes last: ", txBytes, id)
				if err != nil {
			
					return nil, err
				}
				c.account.Sequence++

				resp, err := c.txClient.BroadcastTx(
					context.Background(),
					&tx.BroadcastTxRequest{
						TxBytes: txBytes,
						Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
					},
				)
				if err != nil {
					return nil, err
				}
				return resp.TxResponse, c.handleBroadcastResult(resp.TxResponse, err)
			}
		//messages := msgsList[i*n:((i*n)+n)]
			txBytes, err := c.signTxMsgs(msgsList[i*n:((i*n)+n)], false)
			//fmt.Println("&&&&&&&&&&&&&&&&&&& dispute: ",len(txBytes))
				if err != nil {
			
					return nil, err
				}
				c.account.Sequence++

	_ , err = c.txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			TxBytes: txBytes,
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
		},
	)
	
	
	if err != nil {
		panic("error------------------------------------------------------------")
	}
			// batch := []sdk.Msg{}
			delay := id * 100
	time.Sleep(time.Duration(delay) * time.Millisecond)
			//fmt.Println(i*n, (i*n)+n)	// batch = append(batch, messageBuff[i*10:i*10+10]) messageBuff[i*10:i*10+10]
		}		
	
return nil,nil


	
}

func (c *CosmosClient) BroadcastTx(msg cosmostypes.Msg, adjustGas bool) (*cosmostypes.TxResponse, error) {
	// fmt.Println()

	txBytes, err := c.signTxMsg(msg, adjustGas)
	if err != nil {

		return nil, err
	}

	c.account.Sequence++

	resp, err := c.txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			TxBytes: txBytes,
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.TxResponse, c.handleBroadcastResult(resp.TxResponse, err)
}

func (c *CosmosClient) BroadcastTxs(msg *dkgnet.MsgRefundMsgRequest, adjustGas bool, numOfP int, id int) (*cosmostypes.TxResponse, error) {
	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&: ",msg.Creator)

	messageBuff = append(messageBuff, msg)
	
	if len(messageBuff) != (numOfP - 1) {
		return nil,nil
	}
	//fmt.Println("=========================================================", (messageBuff))
	n := 30
	div := (numOfP-1)/n
	add := (numOfP-1) % n
	if add != 0 {
		add = 1
	}
	//fmt.Println(div)
	delay := id * 300
	time.Sleep(time.Duration(delay) * time.Millisecond)
		for i := 0; i < (div+add); i++ {
			
			if i == (div){
				
				txBytes, err := c.signTxMsgs(messageBuff[i*n:], adjustGas)
				//fmt.Println("txBytes last: ", txBytes, id)
				if err != nil {
			
					return nil, err
				}
				c.account.Sequence++

				resp, err := c.txClient.BroadcastTx(
					context.Background(),
					&tx.BroadcastTxRequest{
						TxBytes: txBytes,
						Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
					},
				)
				if err != nil {
					return nil, err
				}
				return resp.TxResponse, c.handleBroadcastResult(resp.TxResponse, err)
			}
		//messages := messageBuff[i*n:((i*n)+n)]
			txBytes, err := c.signTxMsgs(messageBuff[i*n:((i*n)+n)], adjustGas)
			//fmt.Println("&&&&&&&&&&&&&&&&&&&: ",len(txBytes))
				if err != nil {
			
					return nil, err
				}
				c.account.Sequence++

	_ , err = c.txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			TxBytes: txBytes,
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
		},
	)
	
	
	if err != nil {
		panic("error------------------------------------------------------------")
	}
			// batch := []sdk.Msg{}
			delay := id * 300
	time.Sleep(time.Duration(delay) * time.Millisecond)
			//fmt.Println(i*n, (i*n)+n)	// batch = append(batch, messageBuff[i*10:i*10+10]) messageBuff[i*10:i*10+10]
		}		
	
return nil,nil


	
}

func (c *CosmosClient) WaitForTx(hash string, rate time.Duration) (*tx.GetTxResponse, error) {
	for {
		resp, err := c.txClient.GetTx(context.Background(), &tx.GetTxRequest{Hash: hash})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				time.Sleep(rate)
				continue
			}
			return nil, err
		}
		return resp, err
	}
}

func (c *CosmosClient) signTxMsg(msg cosmostypes.Msg, adjustGas bool) ([]byte, error) {
	encodingCfg := app.MakeEncodingConfig()
	txBuilder := encodingCfg.TxConfig.NewTxBuilder()
	signMode := encodingCfg.TxConfig.SignModeHandler().DefaultMode()

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}

	var newGasLimit uint64 = defaultGasLimit
	if adjustGas {
		txf := clienttx.Factory{}.
			WithGas(defaultGasLimit).
			WithSignMode(signMode).
			WithTxConfig(encodingCfg.TxConfig).
			WithChainID(c.chainID).
			WithAccountNumber(c.account.AccountNumber).
			WithSequence(c.account.Sequence).
			WithGasAdjustment(defaultGasAdjustment)

		_, newGasLimit, err = clienttx.CalculateGas(&c.grpcConn, txf, msg)
		if err != nil {

			return nil, err
		}
	}

	txBuilder.SetGasLimit(newGasLimit)

	signerData := authsigning.SignerData{
		ChainID:       c.chainID,
		AccountNumber: c.account.AccountNumber,
		Sequence:      c.account.Sequence,
		// PubKey:        c.publicKey,
		//Address:       c.account.Address,
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   c.publicKey,
		Data:     &sigData,
		Sequence: c.account.Sequence,
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	sigV2, err := clienttx.SignWithPrivKey(
		signMode, signerData, txBuilder, &c.privateKey,
		encodingCfg.TxConfig, c.account.Sequence,
	)

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := encodingCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}
func (c *CosmosClient) signTxMsgs(msgs []cosmostypes.Msg, adjustGas bool) ([]byte, error) {
	encodingCfg := app.MakeEncodingConfig()
	txBuilder := encodingCfg.TxConfig.NewTxBuilder()
	signMode := encodingCfg.TxConfig.SignModeHandler().DefaultMode()

	err := txBuilder.SetMsgs(msgs...)  // here msgs is a slice of messages
	if err != nil {
		return nil, err
	}

	var newGasLimit uint64 = defaultGasLimit
	if adjustGas {
		txf := clienttx.Factory{}.
			WithGas(defaultGasLimit).
			WithSignMode(signMode).
			WithTxConfig(encodingCfg.TxConfig).
			WithChainID(c.chainID).
			WithAccountNumber(c.account.AccountNumber).
			WithSequence(c.account.Sequence).
			WithGasAdjustment(defaultGasAdjustment)

		// Since CalculateGas expects a single message, you might want to iterate 
		// through your messages and calculate gas for each, or create a new function
		// to calculate gas for all messages together.
		// Here's an example of how you could do it with a loop:
		for _, msg := range msgs {
			_, newGasLimit, err = clienttx.CalculateGas(&c.grpcConn, txf, msg)
			if err != nil {
				return nil, err
			}
		}
	}
	_ = newGasLimit 
	//fmt.Println("calculated gas: ", newGasLimit)
	txBuilder.SetGasLimit(newGasLimit)

	signerData := authsigning.SignerData{
		ChainID:       c.chainID,
		AccountNumber: c.account.AccountNumber,
		Sequence:      c.account.Sequence,
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   c.publicKey,
		Data:     &sigData,
		Sequence: c.account.Sequence,
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	sigV2, err := clienttx.SignWithPrivKey(
		signMode, signerData, txBuilder, &c.privateKey,
		encodingCfg.TxConfig, c.account.Sequence,
	)

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, err
	}

	txBytes, err := encodingCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}
