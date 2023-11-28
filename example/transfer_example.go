package example

import (
	"context"
	"fmt"

	pt "github.com/kehiy/PACTX/lib"
)

func Example() {
	tm, err := pt.NewTxManager(0, "url", "privateKey") // 0 is testnet network type.
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	transferTx, err := tm.MakeTransferTransaction(ctx, 1000, "addr", 8000, "testTX")
	if err != nil {
		panic(err)
	}

	result, err := transferTx.Send(ctx, tm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(result))
}
