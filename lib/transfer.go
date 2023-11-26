package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

func (tm *TxManager) MakeTransferTransaction(ctx context.Context, amt int64,
	receiverAddr, memo string,
) ([]byte, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	info, err := tm.RPCClient.BlockchainClient.GetBlockchainInfo(ctx, &pactus.GetBlockchainInfoRequest{})
	if err != nil {
		return nil, err
	}
	lockTime := info.LastBlockHeight

	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return nil, err
	}

	recAddr, err := bls.PublicKeyFromString(receiverAddr)
	if err != nil {
		return nil, err
	}

	rawTx := tx.NewTransferTx(lockTime, tm.PrivateKey.PublicKeyNative().AccountAddress(),
		recAddr.AccountAddress(), amt, fee.Fee, memo)

	rawTx.SetPublicKey(tm.PrivateKey.PublicKey())
	signBytes := rawTx.SignBytes()
	sign := tm.PrivateKey.Sign(signBytes)
	rawTx.SetSignature(sign)

	txBytes, err := rawTx.Bytes()
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

func (tm *TxManager) SendTransferTransaction(ctx context.Context, rawTx []byte) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx, &pactus.SendRawTransactionRequest{Data: rawTx})
	if err != nil {
		return nil, err
	}
	return res.Id, nil
}
