package lib

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found in account map of transaction manager")
	ErrInvalidPayLoadType = errors.New("invalid payload type")
	ErrInvalidNetworkType = errors.New("invalid network type")
)
