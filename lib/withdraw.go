package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

type WithdrawTx struct {
	RawTx    []byte
	SignedTx []byte
}

func (tm *TxManager) MakeWithdrawTransaction(ctx context.Context,
	validatorAddr, accountAddr string, amt int64, lockTime uint32, memo string,
) (WithdrawTx, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return WithdrawTx{}, err
	}

	// converting validator address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return WithdrawTx{}, err
	}

	// converting receiver account address to publicKey.
	accountPublicKey, err := bls.PublicKeyFromString(accountAddr)
	if err != nil {
		return WithdrawTx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewWithdrawTx(lockTime, validatorPublicKey.AccountAddress(), accountPublicKey.AccountAddress(),
		amt, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return WithdrawTx{}, err
	}

	// setting publicKey, getting bytes for signing, signing the Tx and setting signature for it.
	rawTx.SetPublicKey(tm.PrivateKey.PublicKey())
	signBytes := rawTx.SignBytes()
	sign := tm.PrivateKey.Sign(signBytes)
	rawTx.SetSignature(sign)

	// getting bytes of signed transaction.
	signedTxBytes, err := rawTx.Bytes()
	if err != nil {
		return WithdrawTx{}, err
	}
	return WithdrawTx{SignedTx: signedTxBytes, RawTx: rawTxBytes}, nil
}

func (tm *TxManager) MakeUnsignedWithdrawTransaction(ctx context.Context,
	validatorAddr, accountAddr string, amt int64, lockTime uint32, memo string,
) (WithdrawTx, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return WithdrawTx{}, err
	}

	// converting validator address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return WithdrawTx{}, err
	}

	// converting receiver account address to publicKey.
	accountPublicKey, err := bls.PublicKeyFromString(accountAddr)
	if err != nil {
		return WithdrawTx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewWithdrawTx(lockTime, validatorPublicKey.AccountAddress(), accountPublicKey.AccountAddress(),
		amt, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return WithdrawTx{}, err
	}

	return WithdrawTx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}

func (tt *WithdrawTx) Send(ctx context.Context, tm TxManager) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: tt.SignedTx})
	if err != nil {
		return nil, err
	}
	return res.Id, nil
}

func (tt *WithdrawTx) Raw() []byte {
	return tt.RawTx
}

func (tt *WithdrawTx) Signed() []byte {
	return tt.SignedTx
}
