package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

type UnBondTx struct {
	RawTx    []byte
	SignedTx []byte
}

func (tm *TxManager) MakeUnBondTransaction(validatorAddr string, lockTime uint32, memo string,
) (UnBondTx, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	// converting receiver address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return UnBondTx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewUnbondTx(lockTime, validatorPublicKey.AccountAddress(), memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return UnBondTx{}, err
	}

	// setting publicKey, getting bytes for signing, signing the Tx and setting signature for it.
	rawTx.SetPublicKey(tm.PrivateKey.PublicKey())
	signBytes := rawTx.SignBytes()
	sign := tm.PrivateKey.Sign(signBytes)
	rawTx.SetSignature(sign)

	// getting bytes of signed transaction.
	signedTxBytes, err := rawTx.Bytes()
	if err != nil {
		return UnBondTx{}, err
	}
	return UnBondTx{SignedTx: signedTxBytes, RawTx: rawTxBytes}, nil
}

func (tm *TxManager) MakeUnsignedUnBondTransaction(validatorAddr string, lockTime uint32, memo string,
) (UnBondTx, error) {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"

	// converting receiver address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return UnBondTx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewUnbondTx(lockTime, validatorPublicKey.AccountAddress(), memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return UnBondTx{}, err
	}

	return UnBondTx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}

func (tt *UnBondTx) Send(ctx context.Context, tm TxManager) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: tt.SignedTx})
	if err != nil {
		return nil, err
	}
	return res.Id, nil
}

func (tt *UnBondTx) Raw() []byte {
	return tt.RawTx
}

func (tt *UnBondTx) Signed() []byte {
	return tt.SignedTx
}
