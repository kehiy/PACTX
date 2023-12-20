package lib

import (
	"context"

	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

// MakeWithdrawTransaction makes a signed Withdraw transaction.
func (tm *TxManager) MakeWithdrawTransaction(ctx context.Context,
	validatorAddr, accountAddr string, amt int64, lockTime uint32, memo, accName string,
) (Tx, error) {
	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		return Tx{}, err
	}

	// converting validator address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return Tx{}, err
	}

	senderAddr, ok := tm.Accounts[accName]
	if !ok {
		return Tx{}, ErrAccountNotFound
	}

	// converting receiver account address to publicKey.
	accountPublicKey, err := bls.PublicKeyFromString(accountAddr)
	if err != nil {
		return Tx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewWithdrawTx(lockTime, validatorPublicKey.AccountAddress(), accountPublicKey.AccountAddress(),
		amt, fee.Fee, memo)

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

// MakeUnsignedWithdrawTransaction makes a unsigned (raw) Withdraw transaction.
func (tm *TxManager) MakeUnsignedWithdrawTransaction(ctx context.Context,
	validatorAddr, accountAddr string, amt int64, lockTime uint32, memo string,
) (Tx, error) {
	// getting transaction fee from network.
	// TODO: should we get this as input?
	fee, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_WITHDRAW_PAYLOAD})
	if err != nil {
		return Tx{}, err
	}

	// converting validator address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return Tx{}, err
	}

	// converting receiver account address to publicKey.
	accountPublicKey, err := bls.PublicKeyFromString(accountAddr)
	if err != nil {
		return Tx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewWithdrawTx(lockTime, validatorPublicKey.AccountAddress(), accountPublicKey.AccountAddress(),
		amt, fee.Fee, memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}

	return Tx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}
