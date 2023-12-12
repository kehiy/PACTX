package lib

import (
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
)

// Account is a keeps a Pactus account private key, public key and address.
type Account struct {
	PrivateKey *bls.PrivateKey
	PublicKey  *bls.PublicKey
	Address    crypto.Address
}

func newAccount(privateKey string) (Account, error) {
	if privateKey == "" {
		return Account{}, nil // make an empty account if user provide an empty privateKey.
	}

	pk, err := bls.PrivateKeyFromString(privateKey)
	if err != nil {
		return Account{}, err
	}

	return Account{
		PrivateKey: pk,
		PublicKey:  pk.PublicKeyNative(),
		Address:    pk.PublicKeyNative().AccountAddress(),
	}, nil
}
