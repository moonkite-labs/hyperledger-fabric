package utils

import (
	"encoding/base64"
	"encoding/hex"
	"hash"

	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

var hashFactory = map[int]func() hash.Hash{
	256: sha3.New256,
	384: sha3.New384,
	512: sha3.New512,
}

// Return a hash result of given input using SHA3 algorithm of given size
func SHASum(input []byte, size int) ([]byte, error) {
	f := hashFactory[size]
	if f == nil {
		return nil, errors.New("Failed to retrieve a hash function! Invalid size provided.")
	}
	h := f()
	h.Write(input)
	return h.Sum(nil), nil
}

// Return a hex encoded hash result of given input using SHA3 algoritm of given size
func SHASumHex(input []byte, size int) (string, error) {
	shaBytes, err := SHASum(input, size)
	return hex.EncodeToString(shaBytes), err
}

// Return a base64 encoded hash result of given input using SHA3 algoritm of given size
func SHASumBase64(input []byte, size int, isUrl bool) (string, error) {
	shaBytes, err := SHASum(input, size)
	if isUrl {
		return base64.URLEncoding.EncodeToString(shaBytes), err
	}
	return base64.StdEncoding.EncodeToString(shaBytes), err
}
