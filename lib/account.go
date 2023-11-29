package lib

import (
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
)

type Account struct {
	PrivateKey *bls.PrivateKey
	PublicKey  *bls.PublicKey
	Address    crypto.Address
}

func NewAccount(privateKey string) (Account, error) {
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