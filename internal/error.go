package internal

import "errors"

// errors.
var (
	ErrJsonObject     = errors.New("input is not json object or array")
	ErrKeyNotFound    = errors.New("no matched key found")
	ErrArrayStructure = errors.New("wrong array structure")
)
