package example

import (
	"context"
	"fmt"

	pt "github.com/kehiy/PACTX/lib"
)

func Example() {
	tm, err := pt.NewTxManager("url", "privateKey")
	if err != nil {
		panic(err)
	}

	rawTX, err := tm.MakeTransferTransaction(context.Background(), 1000, "addr", "testTX")
	if err != nil {
		panic(err)
	}

	result, err := tm.SendTransferTransaction(context.Background(), rawTX)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(result))
}
