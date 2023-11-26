package lib

import (
	"github.com/kehiy/PACTX/client"
	"github.com/pactus-project/pactus/crypto/bls"
)

type TxManager struct {
	Provider   string
	RPCClient  *client.Client
	PrivateKey *bls.PrivateKey
}

func NewTxManager(rpcurl, privatekey string) (TxManager, error) {
	pk, err := bls.PrivateKeyFromString(privatekey)
	if err != nil {
		return TxManager{}, err
	}

	c, err := client.NewClient(rpcurl)
	if err != nil {
		return TxManager{}, err
	}

	return TxManager{
		Provider:   rpcurl,
		PrivateKey: pk,
		RPCClient:  c,
	}, nil
}
