package client

import (
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a gRPC client of Pactus Blockchain for transactions, blockchain and network.
type Client struct {
	NetworkClient     pactus.NetworkClient
	BlockchainClient  pactus.BlockchainClient
	TransactionClient pactus.TransactionClient
	Conn              *grpc.ClientConn
}

// NewClient get a RPC url as input and returns a Client.
func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		NetworkClient:     pactus.NewNetworkClient(conn),
		BlockchainClient:  pactus.NewBlockchainClient(conn),
		TransactionClient: pactus.NewTransactionClient(conn),
		Conn:              conn,
	}, nil
}

// Close will close the connection to RPC node.
func (c *Client) Close() error {
	return c.Conn.Close()
}
