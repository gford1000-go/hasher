package hasher

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/gob"
	"errors"
	"fmt"
)

// InvalidHash is an invalid hash value - not initialised
var InvalidHash = []byte{}

// toBytes uses gob to construct a []byte slice, which ensures the []byte
// representation of the interface{} traverses references to other objects
func toBytes(d any) (data []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("toBytes() panicked: %v", r)
		}
	}()

	var stream bytes.Buffer
	enc := gob.NewEncoder(&stream)
	e := enc.Encode(d)
	if e != nil {
		return nil, fmt.Errorf("failed to covert %v to []byte: %v", d, e)
	}
	return stream.Bytes(), nil
}

// Hash returns a DataHash of the supplied interface{}.
// If the interface{} is a struct, then only public attributes will
// be used to construct the DataHash
func Hash(i any, opts ...func(*Options)) ([]byte, error) {

	o := Options{
		HashType: Sha256,
	}
	for _, opt := range opts {
		opt(&o)
	}

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
	switch o.HashType {
	case Sha256:
		v := sha256.Sum256(b)
		return v[:], nil
	case Sha384:
		v := sha512.Sum384(b)
		return v[:], nil
	case Sha512:
		v := sha512.Sum512(b)
		return v[:], nil
	default:
		return nil, errors.New("unexpected error on hash type")
	}
}
