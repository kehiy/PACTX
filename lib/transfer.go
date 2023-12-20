package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

// MakeTransferTransaction makes a signed Transfer transaction.
func (tm *TxManager) MakeTransferTransaction(ctx context.Context, amt int64,
	receiverAddr string, lockTime uint32, memo, accName string,
) (Tx, error) {
	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
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
	rawTx := tx.NewTransferTx(lockTime, senderAddr.Address,
		receiverPublicKey.AccountAddress(), amt, fee.Fee, memo)

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

// MakeUnsignedTransferTransaction makes a unsigned (raw) Transfer transaction.
func (tm *TxManager) MakeUnsignedTransferTransaction(ctx context.Context, amt int64,
	receiverAddr string, lockTime uint32, memo, accName string,
) (Tx, error) {
	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
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
	rawTx := tx.NewTransferTx(lockTime, senderAddr.Address,
		receiverPublicKey.AccountAddress(), amt, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}

	return Tx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}
