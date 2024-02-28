package main

import "hash/fnv"

// The type of hashing strategy to use
type HashStrategy interface {
	Hash([]byte, int) int
}

type FNVHashStrategy struct{}

// Hash hash using a FNV
func (f FNVHashStrategy) Hash(data []byte, seed int) int {
	h := fnv.New32a()
	h.Write(data)
	// Convert the hash to an integer, might need adjustments based on your requirements.
	return int(h.Sum32()) % (seed + 1)
}

// Bloom is a struct that holds the bloom filter information
type Bloom struct {
	bitArray    []byte
	numOfHashes int

	// The hashing strategy to use
	hashStrategy HashStrategy
}

// Return a new bloom filter
func NewBloom(bitArrayLen int, numOfHashes int, hashStrategy HashStrategy) Bloom {
	return Bloom{
		bitArray:     make([]byte, bitArrayLen),
		numOfHashes:  numOfHashes,
		hashStrategy: hashStrategy,
	}
}

// Add adds a new object to the bloom filter
func (b *Bloom) Add(o []byte) {
	for i := 0; i < b.numOfHashes; i++ {
		r := b.hashStrategy.Hash(o, i)

		b.bitArray[r%len(b.bitArray)] = 1
	}
}

// Check checks if an item exists on the filter
func (b Bloom) Check(o []byte) bool {
	// Run the our hashes and check against the bitArray
	for i := 0; i < b.numOfHashes; i++ {
		r := b.hashStrategy.Hash(o, i)
		if b.bitArray[r%len(b.bitArray)] == 0 {
			return false
		}
	}

	return true
}
