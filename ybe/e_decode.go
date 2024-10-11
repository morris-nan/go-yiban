package ybe

import "errors"

var (
	ErrBlockSize = errors.New("ciphertext is not a multiple of the block size")
)
