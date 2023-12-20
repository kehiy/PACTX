package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

// MakeBondTransaction makes a signed Bond transaction.
func (tm *TxManager) MakeBondTransaction(ctx context.Context, stake int64,
	receiverAddr string, lockTime uint32, memo, accName string,
) (Tx, error) {
	if tm.NetworkType == TestNet {
		crypto.AddressHRP = "tpc"
		crypto.PublicKeyHRP = "tpublic"
		crypto.PrivateKeyHRP = "tsecret"
		crypto.XPublicKeyHRP = "txpublic"
		crypto.XPrivateKeyHRP = "txsecret"
	} else {
		crypto.AddressHRP = "pc"
		crypto.PublicKeyHRP = "public"
		crypto.PrivateKeyHRP = "secret"
		crypto.XPublicKeyHRP = "xpublic"
		crypto.XPrivateKeyHRP = "xsecret"
	}

	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: stake, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return Tx{}, err
	}

	// converting receiver address to publicKey.
	receiverPublicKey, err := bls.PublicKeyFromString(receiverAddr)
	if err != nil {
		return Tx{}, err
	}

	senderAddr, ok := tm.Accounts[accName]
	if !ok {
		return Tx{}, ErrAccountNotFound
	}

	// making raw transaction.
	rawTx := tx.NewBondTx(lockTime, senderAddr.Address,
		receiverPublicKey.AccountAddress(), receiverPublicKey, stake, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}

	// setting publicKey, getting bytes for signing, signing the Tx and setting signature for it.
	rawTx.SetPublicKey(senderAddr.PublicKey)
	signBytes := rawTx.SignBytes()
	sign := senderAddr.PrivateKey.Sign(signBytes)
	rawTx.SetSignature(sign)

	// getting bytes of signed transaction.
	signedTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}
	return Tx{SignedTx: signedTxBytes, RawTx: rawTxBytes}, nil
}

// MakeUnsignedBondTransaction makes a unsigned (raw) Bond transaction.
func (tm *TxManager) MakeUnsignedBondTransaction(ctx context.Context, stake int64,
	receiverAddr string, lockTime uint32, memo, accName string,
) (Tx, error) {
	if tm.NetworkType == TestNet {
		crypto.AddressHRP = "tpc"
		crypto.PublicKeyHRP = "tpublic"
		crypto.PrivateKeyHRP = "tsecret"
		crypto.XPublicKeyHRP = "txpublic"
		crypto.XPrivateKeyHRP = "txsecret"
	} else {
		crypto.AddressHRP = "pc"
		crypto.PublicKeyHRP = "public"
		crypto.PrivateKeyHRP = "secret"
		crypto.XPublicKeyHRP = "xpublic"
		crypto.XPrivateKeyHRP = "xsecret"
	}

	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: stake, PayloadType: pactus.PayloadType_BOND_PAYLOAD})
	if err != nil {
		return Tx{}, err
	}

	// converting receiver address to publicKey.
	receiverPublicKey, err := bls.PublicKeyFromString(receiverAddr)
	if err != nil {
		return Tx{}, err
	}

	senderAddr, ok := tm.Accounts[accName]
	if !ok {
		return Tx{}, ErrAccountNotFound
	}

	// making raw transaction.
	rawTx := tx.NewBondTx(lockTime, senderAddr.Address,
		receiverPublicKey.AccountAddress(), receiverPublicKey, stake, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}

	return Tx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}
