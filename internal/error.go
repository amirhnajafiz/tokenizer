package internal

import "errors"

var (
	ErrConfFileNotFound = errors.New("conf file not found")
	ErrKeyNotSet        = errors.New("no token is set for this key")
	ErrScanner          = errors.New("file scanner failure")
	ErrFileCreation     = errors.New("failed to create a new file to export")
)
