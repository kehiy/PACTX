package lib

import (
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
)

// MakeUnBondTransaction makes a signed UnBond transaction.
func (tm *TxManager) MakeUnBondTransaction(validatorAddr string, lockTime uint32,
	memo, accName string,
) (Tx, error) {
	// converting receiver address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return Tx{}, err
	}

	senderAddr, ok := tm.Accounts[accName]
	if !ok {
		return Tx{}, ErrAccountNotFound
	}

	// making raw transaction.
	rawTx := tx.NewUnbondTx(lockTime, validatorPublicKey.AccountAddress(), memo)

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

// MakeUnsignedUnBondTransaction makes a unsigned (raw) UnBond transaction.
func (tm *TxManager) MakeUnsignedUnBondTransaction(validatorAddr string, lockTime uint32, memo string,
) (Tx, error) {
	// converting receiver address to publicKey.
	validatorPublicKey, err := bls.PublicKeyFromString(validatorAddr)
	if err != nil {
		return Tx{}, err
	}

	// making raw transaction.
	rawTx := tx.NewUnbondTx(lockTime, validatorPublicKey.AccountAddress(), memo)

	// keep raw transaction bytes for RawTx field in TransferTx.
	rawTxBytes, err := rawTx.Bytes()
	if err != nil {
		return Tx{}, err
	}

	return Tx{SignedTx: make([]byte, 0), RawTx: rawTxBytes}, nil
}
