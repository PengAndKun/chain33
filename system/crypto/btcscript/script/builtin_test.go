package script_test

import (
	"testing"

	"github.com/33cn/chain33/client/mocks"
	cryptocli "github.com/33cn/chain33/common/crypto/client"
	nty "github.com/33cn/chain33/system/dapp/none/types"
	"github.com/stretchr/testify/mock"

	"github.com/33cn/chain33/system/crypto/btcscript"

	"github.com/33cn/chain33/system/crypto/btcscript/script"
	"github.com/33cn/chain33/types"
	"github.com/33cn/chain33/util"
	"github.com/stretchr/testify/require"
)

func Test_WalletRecoveryScript(t *testing.T) {

	_, controlKey := util.Genaddress()
	_, recoverKey := util.Genaddress()
	delayTime := int64(10) //10 block height

	pkScript, err := script.NewWalletRecoveryScript(
		controlKey.PubKey().Bytes(), recoverKey.PubKey().Bytes(), delayTime)

	require.Nil(t, err)
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	tx := util.CreateNoneTx(cfg, nil)
	signMsg := types.Encode(tx)

	// withdraw wallet balance with control address
	sig, pubKey, err := script.GetWalletRecoverySignature(false, signMsg, controlKey.Bytes(),
		pkScript, delayTime)

	require.Nil(t, err)

	tx.Signature = &types.Signature{
		Ty:        btcscript.ID,
		Pubkey:    pubKey,
		Signature: sig,
	}
	require.True(t, tx.CheckSign(0))

	api := new(mocks.QueueProtocolAPI)
	cryptocli.SetQueueAPI(api)
	cryptocli.SetCurrentBlock(delayTime-1, types.Now().Unix())
	runFn := func(args mock.Arguments) {
		execer := args.Get(0).(string)
		require.Equal(t, nty.NoneX, execer)
		funcName := args.Get(1).(string)
		require.Equal(t, nty.QueryGetDelayBegin, funcName)
		param := args.Get(2).(*types.ReqBytes)
		require.Equal(t, tx.Hash(), param.Data)
	}
	api.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(&types.Int64{}, nil).Run(runFn)

	// withdraw wallet balance with recover address
	sig, pubKey, err = script.GetWalletRecoverySignature(true, signMsg, recoverKey.Bytes(),
		pkScript, delayTime)

	require.Nil(t, err)

	tx.Signature = &types.Signature{
		Ty:        btcscript.ID,
		Pubkey:    pubKey,
		Signature: sig,
	}
	require.True(t, tx.CheckSign(10))

	// delay time not satisfied
	sig, pubKey, err = script.GetWalletRecoverySignature(true, signMsg, recoverKey.Bytes(),
		pkScript, delayTime-1)

	require.Nil(t, err)
	tx.Signature = &types.Signature{
		Ty:        btcscript.ID,
		Pubkey:    pubKey,
		Signature: sig,
	}
	require.False(t, tx.CheckSign(10))

	// without delay time
	sig, pubKey, err = script.GetWalletRecoverySignature(false, signMsg, recoverKey.Bytes(),
		pkScript, delayTime-1)

	require.Nil(t, err)
	tx.Signature = &types.Signature{
		Ty:        btcscript.ID,
		Pubkey:    pubKey,
		Signature: sig,
	}
	require.False(t, tx.CheckSign(0))
}

func Test_MultisigScript(t *testing.T) {
	_, priv1 := util.Genaddress()
	_, priv2 := util.Genaddress()

	_, err := script.NewMultiSigScript(nil, 0)
	require.Equal(t, script.ErrInvalidMultiSigRequiredNum, err)

	_, err = script.NewMultiSigScript([][]byte{priv1.PubKey().Bytes()}, 2)
	require.Equal(t, script.ErrInvalidMultiSigRequiredNum, err)

	_, err = script.NewMultiSigScript([][]byte{priv1.PubKey().Bytes()}, 1)
	require.Nil(t, err)

	_, err = script.NewMultiSigScript([][]byte{priv1.PubKey().Bytes(), priv2.PubKey().Bytes()}, 1)
	require.Nil(t, err)
}
