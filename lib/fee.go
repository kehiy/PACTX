package lib

import (
	"context"

	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

func (tm *TxManager) CalcFee(ctx context.Context, amt int64, payloadType string) (int64, error) {
	var pType pactus.PayloadType
	switch payloadType {
	case "transfer":
		pType = pactus.PayloadType_TRANSFER_PAYLOAD
	case "bond":
		pType = pactus.PayloadType_BOND_PAYLOAD
	case "unbond":
		return 0, nil
	case "withdraw":
		pType = pactus.PayloadType_WITHDRAW_PAYLOAD
	default:
		return 0, ErrInvalidPayLoadType
	}
	result, err := tm.RPCClient.TransactionClient.CalculateFee(ctx,
		&pactus.CalculateFeeRequest{Amount: amt, PayloadType: pType})
	if err != nil {
		return 0, err
	}

	return result.GetFee(), nil
}
