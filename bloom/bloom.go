package bloom

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

type BloomFilter struct {
	maxBits   uint
	hashFuncs uint
	set       *bitset.BitSet
}

func hashers(data []byte) [2]uint64 {
	hasher := murmur3.New64()
	if _, err := hasher.Write(data); err != nil {
		panic(err)
	}
	v1 := hasher.Sum64()

	a1 := []byte{1}
	if _, err := hasher.Write(a1); err != nil {
		panic(err)
	}
	v2 := hasher.Sum64()
	return [2]uint64{v1, v2}
}

func (f *BloomFilter) location(h [2]uint64, i uint) uint {
	return uint(h[i] % uint64(f.maxBits))
}

func (f *BloomFilter) Add(data []byte) {
	for i := uint(0); i < f.hashFuncs; i++ {
		loc := f.location(hashers(data), i)
		f.set.Set(loc)
	}
}

func (f *BloomFilter) Check(data []byte) bool {
	for i := uint(0); i < f.hashFuncs; i++ {
		loc := f.location(hashers(data), i)
		if !f.set.Test(loc) {
			return false
		}
	}

	return true
}

func New(maxBits, hashFuncs uint) *BloomFilter {
	return &BloomFilter{maxBits, hashFuncs, bitset.New(maxBits)}
}
