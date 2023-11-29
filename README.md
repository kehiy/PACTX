# PACTx

<img alt="pactx" src="https://github.com/kehiy/PACTX/assets/89645414/7b82344a-634f-49c8-b94a-c3b8b2a98ee9" width="150" />

Pactus Transaction is a tool written in golang which help you to make, send and manage transactions in [Pactus Blockchain](https://pactus.org) using multiple account with a good control.

> PACTx cli tool also helps you to send bulk transactions with different types and more without leaving terminal.

# Usage

Install:

```bash
go get -u github.com/kehiy/PACTX@latest
```

Example:

```go
package example

import (
	"context"
	"fmt"

	pt "github.com/kehiy/PACTX/lib"
)

func Example() {
	ctx := context.Background()

	/*
	0 is testnet network type.

	account name will point to this specific private key.
	you can have multiple private keys with different names

	consider to get private key from and env, config file and ...
	*/
	tm, err := pt.NewTxManager(0, "url", "private-key", "first-account-name")
	if err != nil {
		panic(err)
	}

	transferTx, err := tm.MakeTransferTransaction(ctx, 1000, "addr", 8000,
	 "testTX", "first-account-name")
	if err != nil {
		panic(err)
	}

	result, err := transferTx.Send(ctx, tm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(result)) // result is your transaction ID.

	err = tm.AddAccount("second-private-key", "second-account-name")
	if err != nil {
		panic(err)
	}
}
```

## Contribution

Contributions to the PACTX are welcomed.

## License

The PACTX is under MIT [license](./LICENSE).
