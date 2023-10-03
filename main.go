package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kehiy/PACTX/client"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	tutil "github.com/pactus-project/pactus/util/testsuite"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

var locktime uint32

func main() {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"
	t := tutil.NewTestSuiteForSeed(1)
	ctx := context.Background()

	c, err := client.NewClient(os.Args[1])
	if err != nil {
		panic(err)
	}

	pk, err := bls.PrivateKeyFromString(os.Args[2])
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 1000; i++ {
		info, err := c.BlockchainClient.GetBlockchainInfo(ctx, &pactus.GetBlockchainInfoRequest{})
		if err != nil {
			panic(err)
		}
		locktime = info.LastBlockHeight
		amt := t.RandInt64(1e9)
		fee, err := c.TransactionClient.CalculateFee(ctx,
			&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
		if err != nil {
			panic(err)
		}

		receiver := t.RandAccAddress()
		tx := tx.NewTransferTx(locktime, pk.PublicKeyNative().AccountAddress(),
			receiver, amt, fee.Fee, "test")

		tx.SetPublicKey(pk.PublicKey())
		signBytes := tx.SignBytes()
		sign := pk.Sign(signBytes)
		tx.SetSignature(sign)

		btx, err := tx.Bytes()
		if err != nil {
			panic(err)
		}

		res, err := c.TransactionClient.SendRawTransaction(ctx,
			&pactus.SendRawTransactionRequest{Data: btx})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TX ID: %x\nTX Num: %d\nAmount: %v\nReceiver: %v\n", res.Id, i, amt, receiver)
	}
}
