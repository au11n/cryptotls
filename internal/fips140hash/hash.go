package fips140hash

import "hash"

func Unwrap(h hash.Hash) hash.Hash { return h }

func UnwrapNew[Hash hash.Hash](newHash func() Hash) func() hash.Hash {
	return func() hash.Hash { return newHash() }
}
