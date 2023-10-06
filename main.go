package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/kehiy/PACTX/client"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/types/tx"
	tutil "github.com/pactus-project/pactus/util/testsuite"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

var (
	lockTime uint32
	txNum    uint32
	wg       = &sync.WaitGroup{}
)

func main() {
	crypto.AddressHRP = "tpc"
	crypto.PublicKeyHRP = "tpublic"
	crypto.PrivateKeyHRP = "tsecret"
	crypto.XPublicKeyHRP = "txpublic"
	crypto.XPrivateKeyHRP = "txsecret"
	t := tutil.NewTestSuiteForSeed(1)

	pk, err := bls.PrivateKeyFromString(os.Args[2])
	if err != nil {
		panic(err.Error())
	}

	transactionCount := int64(1000)
	concurrentGoroutine := 200

	if len(os.Args) > 3 {
		transactionCountArg := os.Args[3]
		if strings.TrimSpace(transactionCountArg) != "" {
			transactionCount, err = strconv.ParseInt(transactionCountArg, 10, 64)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	totalTransaction := transactionCount / int64(concurrentGoroutine)

	for i := int64(0); i < totalTransaction; i++ {
		wg.Add(concurrentGoroutine)
		for j := 0; j < concurrentGoroutine; j++ {
			go sendTransaction(txNum, t, pk)
			txNum++
		}
		wg.Wait()
	}
}

func sendTransaction(i uint32, t *tutil.TestSuite, pk *bls.PrivateKey) {
	defer wg.Done()
	ctx := context.Background()

	c, err := client.NewClient(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer c.Close()

	info, err := c.BlockchainClient.GetBlockchainInfo(ctx, &pactus.GetBlockchainInfoRequest{})
	if err != nil {
		panic(err)
	}

	lockTime = info.LastBlockHeight
	amt := t.RandInt64(1e9)
	fee, err := c.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pactus.PayloadType_TRANSFER_PAYLOAD})
	if err != nil {
		panic(err)
	}

	receiver := t.RandAccAddress()
	transferTx := tx.NewTransferTx(lockTime, pk.PublicKeyNative().AccountAddress(),
		receiver, amt, fee.Fee, "test")

	transferTx.SetPublicKey(pk.PublicKey())
	signBytes := transferTx.SignBytes()
	sign := pk.Sign(signBytes)
	transferTx.SetSignature(sign)

	btx, err := transferTx.Bytes()
	if err != nil {
		panic(err)
	}

	res, err := c.TransactionClient.SendRawTransaction(ctx,
		&pactus.SendRawTransactionRequest{Data: btx})
	if err != nil {
		panic(err)
	}
	log.Printf("TX ID: %x\nTX Num: %d\nAmount: %v\nReceiver: %v\n", res.Id, i, amt, receiver)
}
