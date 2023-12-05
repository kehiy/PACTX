package lib

import (
	"context"

	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

// Tx struct contains a specific transaction signed bytes and raw transaction bytes.
type Tx struct {
	RawTx    []byte
	SignedTx []byte

	// ID will be filled when Send method called.
	ID []byte
}

// Send sends (publish, broadcasts) a transaction to the network using RPC node urls.
func (tx *Tx) Send(ctx context.Context, tm *TxManager) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: tx.SignedTx})
	if err != nil {
		return nil, err
	}
	tx.ID = res.Id

	return res.Id, nil
}

// Raw returns bytes of a raw (unsigned) transaction.
func (tx *Tx) Raw() []byte {
	return tx.RawTx
}

// Signed returns bytes of transaction which is signed by a private key and have a signature.
func (tx *Tx) Signed() []byte {
	return tx.SignedTx
}

// GetInfo will get a transaction info from network.
func (tx *Tx) GetInfo(ctx context.Context, tm TxManager) (*pactus.GetTransactionResponse, error) {
	res, err := tm.RPCClient.TransactionClient.GetTransaction(ctx,
		&pactus.GetTransactionRequest{Id: tx.ID, Verbosity: pactus.TransactionVerbosity_TRANSACTION_INFO})
	if err != nil {
		return &pactus.GetTransactionResponse{}, err
	}

	return res, nil
}

// GetData will get a transaction data from network.
func (tx *Tx) GetData(ctx context.Context, tm TxManager) (*pactus.GetTransactionResponse, error) {
	res, err := tm.RPCClient.TransactionClient.GetTransaction(ctx,
		&pactus.GetTransactionRequest{Id: tx.ID, Verbosity: pactus.TransactionVerbosity_TRANSACTION_DATA})
	if err != nil {
		return &pactus.GetTransactionResponse{}, err
	}

	return res, nil
}
