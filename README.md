# PACTX

Pactus Bulk Transaction Sender.
This CLI tool helps you to send bulk transaction in Pactus network.

You can also use PACTX as a library to make transactions with golang code.


# Usage

```bash
go run main.go [node_IP:gRPC-port] [private_key] [transction count (optional|default=1000)]
```

# Package

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
```

## Contribution

Contributions to the PACTX are welcomed.

## License

The PACTX is under MIT [license](./LICENSE).