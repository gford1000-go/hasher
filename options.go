package hasher

// HasherType determine the valid Hashing algos that can be used
type HashType int

const (
	UnknownHashType HashType = iota
	Sha256
	Sha384
	Sha512
	InvalidHashType
)

func (h HashType) isValid() bool {
	return h > UnknownHashType && h < InvalidHashType
}

// Options changes the behaviour of Hash()
type Options struct {
	HashType HashType
}

// WithHashType allows different hashing algos to be used
func WithHashType(ht HashType) func(*Options) {
	return func(o *Options) {
		if ht.isValid() {
			o.HashType = ht
		}
	}
}
