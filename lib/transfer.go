package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

type TransferTx struct {
	RawTx    []byte
	SignedTx []byte
}

func (tm *TxManager) MakeTransferTransaction(ctx context.Context, amt int64,
	receiverAddr string, lockTime uint32, memo string,
) (TransferTx, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return TransferTx{}, err
	}

	recAddr, err := bls.PublicKeyFromString(receiverAddr)
	if err != nil {
		return TransferTx{}, err
	}

	rawTx := tx.NewTransferTx(lockTime, tm.PrivateKey.PublicKeyNative().AccountAddress(),
		recAddr.AccountAddress(), amt, fee.Fee, memo)

	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return TransferTx{}, err
	}

	rawTx.SetPublicKey(tm.PrivateKey.PublicKey())
	signBytes := rawTx.SignBytes()
	sign := tm.PrivateKey.Sign(signBytes)
	rawTx.SetSignature(sign)

	signedTxBytes, err := rawTx.Bytes()
	if err != nil {
		return TransferTx{}, err
	}
	return TransferTx{SignedTx: signedTxBytes, RawTx: rawTxBytes}, nil
}

func (tt *TransferTx) Send(ctx context.Context, tm TxManager) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: tt.SignedTx})
	if err != nil {
		return nil, err
	}
	return res.Id, nil
}

func (tt *TransferTx) Raw() []byte {
	return tt.RawTx
}

func (tt *TransferTx) Signed() []byte {
	return tt.SignedTx
}
