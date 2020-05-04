package hasher

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
)

// DataHash is the type of a hash value
type DataHash [32]byte

// InvalidHash is an invalid hash value - not initialised
var InvalidHash DataHash

// toBytes uses gob to construct a []byte slice, which ensures the []byte
// representation of the interface{} traverses references to other objects
func toBytes(d interface{}) (data []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("toBytes() panicked: %v", r)
		}
	}()

	var stream bytes.Buffer
	enc := gob.NewEncoder(&stream)
	e := enc.Encode(d)
	if e != nil {
		return nil, fmt.Errorf("Failed to covert %v to []byte: %v", d, e)
	}
	return stream.Bytes(), nil
}

// Hash returns a DataHash of the supplied interface{}.
// If the interface{} is a struct, then only public attributes will
// be used to construct the DataHash
func Hash(i interface{}) (DataHash, error) {
	if i == nil {
		return InvalidHash, errors.New("i must not be nil")
	}
	b, ok := i.([]byte)
	if !ok {
		var err error
		b, err = toBytes(i)
		if err != nil {
			return InvalidHash, err
		}
	}
	return sha256.Sum256(b), nil
}
