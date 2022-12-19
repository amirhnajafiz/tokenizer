package internal

import "errors"

var (
	ErrJsonObject  = errors.New("input is not JSON object or array")
	ErrKeyNotFound = errors.New("no matched key")
)
