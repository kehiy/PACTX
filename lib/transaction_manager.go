package lib

import (
	"github.com/kehiy/PACTX/client"
	"github.com/pactus-project/pactus/crypto/bls"
)

type NetworkType int

const (
	TestNet NetworkType = 0
	MainNet NetworkType = 1
	DevNet  NetworkType = 2
)

type TxManager struct {
	Provider    string
	RPCClient   *client.Client
	PrivateKey  *bls.PrivateKey
	NetworkType NetworkType
}

func NewTxManager(networkType NetworkType, rpcurl, privatekey string) (TxManager, error) {
	pk, err := bls.PrivateKeyFromString(privatekey)
	if err != nil {
		return TxManager{}, err
	}

	c, err := client.NewClient(rpcurl)
	if err != nil {
		return TxManager{}, err
	}

	return TxManager{
		Provider:    rpcurl,
		PrivateKey:  pk,
		RPCClient:   c,
		NetworkType: networkType,
	}, nil
}
