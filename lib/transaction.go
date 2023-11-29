package lib

import (
	"context"

	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

// Tx struct contains a specific transaction signed bytes and raw transaction bytes.
type Tx struct {
	RawTx    []byte
	SignedTx []byte
}

// Send sends (publish, broadcasts) a transaction to the network using RPC node urls.
func (tt *Tx) Send(ctx context.Context, tm TxManager) ([]byte, error) {
	res, err := tm.RPCClient.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: tt.SignedTx})
	if err != nil {
		return nil, err
	}
	return res.Id, nil
}

// Raw returns bytes of a raw (unsigned) transaction.
func (tt *Tx) Raw() []byte {
	return tt.RawTx
}

// Signed returns bytes of transaction which is signed by a private key and have a signature.
func (tt *Tx) Signed() []byte {
	return tt.SignedTx
}
