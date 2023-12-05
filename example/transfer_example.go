package example

import (
	"context"
	"fmt"

	pt "github.com/kehiy/PACTX/lib"
)

func Example() {
	ctx := context.Background()

	/*
			account name will point to this specific private key.
		    (you can have multiple private keys with different names (multiple pacuts accounts))

			consider to get private key from and env, config file and ...
	*/
	tm, err := pt.NewTxManager(pt.TestNet, "url", "private-key", "first-account-name")
	if err != nil {
		panic(err)
	}

	transferTx, err := tm.MakeTransferTransaction(ctx, 1000, "addr", 8000, "testTX", "first-account-name")
	if err != nil {
		panic(err)
	}

	result, err := transferTx.Send(ctx, tm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x", result) // result is your transaction ID.

	err = tm.AddAccount("second-private-key", "second-account-name")
	if err != nil {
		panic(err)
	}
}
