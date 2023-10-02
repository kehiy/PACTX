package client

import (
	"github.com/k0kubun/pp"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	NetworkClient     pactus.NetworkClient
	BlockchainClient  pactus.BlockchainClient
	TransactionClient pactus.TransactionClient
	Conn              *grpc.ClientConn
}

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	pp.Println("connection established...")

	return &Client{
		NetworkClient:     pactus.NewNetworkClient(conn),
		BlockchainClient:  pactus.NewBlockchainClient(conn),
		TransactionClient: pactus.NewTransactionClient(conn),
		Conn:              conn,
	}, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
