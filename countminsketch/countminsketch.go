package countminsketch

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

type countMin struct {
	hashFuncs uint
	tableSize uint
	counts    [][]uint64
}

func New(hashFuncs, tableSize uint) *countMin {
	c := &countMin{
		hashFuncs: hashFuncs,
		tableSize: tableSize,
		counts:    make([][]uint64, hashFuncs),
	}

	for i := uint(0); i < hashFuncs; i++ {
		c.counts[i] = make([]uint64, tableSize)
	}

	return c
}

func (c *countMin) Add(data []byte, count uint64) {
	fmt.Println("counts", c.counts)
	for row, col := range c.locations(data) {
		c.counts[row][col] += count
	}
}

func (c *countMin) Estimate(data []byte) uint64 {
	var min uint64
	for row, col := range c.locations(data) {
		curr := c.counts[row][col]
		if row == 0 || curr < min {
			min = curr
		}
	}
	return min
}

func hashers(data []byte) (h1, h2 uint32) {
	hasher := fnv.New64()
	if _, err := hasher.Write(data); err != nil {
		panic(err)
	}

	sum := hasher.Sum(nil)

	return binary.BigEndian.Uint32(sum[0:4]), binary.BigEndian.Uint32(sum[4:8])
}

func (c *countMin) locations(data []byte) []uint {
	locs := make([]uint, c.hashFuncs)

	lower, upper := hashers(data)
	for i := uint(0); i < c.hashFuncs; i++ {
		locs[i] = (uint(lower) + uint(upper)*i) % c.tableSize
	}
	return locs
}
