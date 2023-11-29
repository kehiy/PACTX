package lib

import (
	"github.com/kehiy/PACTX/client"
	"github.com/pactus-project/pactus/crypto/bls"
)

// NetworkType is a type that determine which network you are using.
type NetworkType int

const (
	// TestNet network type is Pactus testnet.
	TestNet NetworkType = 0

	// MainNet network type is Pactus mainnet.
	MainNet NetworkType = 1

	// DevNet network type is Pactus dev-net or a private network on pactus protocol you are running.
	DevNet NetworkType = 2
)

// TxManager helps you to make, send and work with transaction in Pactus Blockchain.
type TxManager struct {
	// Provider is a RPC node url for sending and getting data from.
	Provider string

	// RPCClient is a Client for Pactus gRPC service.
	RPCClient *client.Client

	// PrivateKey is the private key of account on behalf of which transactions are made and sent.
	PrivateKey *bls.PrivateKey

	// NetworkType helps to determine which address prefixes and HRPs should be used, like: pc, tpc and more.
	NetworkType NetworkType
}

// NewTxManager returns a TxManager by provided parameters.
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
