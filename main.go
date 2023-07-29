package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

var hasher hash.Hash32

type BloomFilter struct {
	filter []byte
	size   uint32
}

func (bloomFilter BloomFilter) Add(key string) {
	index := Hash(key, bloomFilter.size)
	bloomFilter.filter[index] = 1
}

func (bloomFilter BloomFilter) Exists(key string) bool {
	index := Hash(key, bloomFilter.size)
	if bloomFilter.filter[index] == 0 {
		return false
	}
	return true
}

func (bloomFilter BloomFilter) Filter() {
	fmt.Println(bloomFilter.filter)
}

func NewBloomFilter(size uint32) BloomFilter {
	filter := make([]byte, size)
	return BloomFilter{filter, size}
}

func init() {
	hasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func Hash(key string, size uint32) uint32 {
	hasher.Write([]byte(key))
	result := hasher.Sum32() % size
	hasher.Reset()
	return result
}

func main() {
	bloomFilter := NewBloomFilter(16)
	keys := []string{"one", "two", "three"}

	for _, key := range keys {
		bloomFilter.Add(key)
	}
	for _, key := range keys {
		fmt.Println(key, bloomFilter.Exists(key))
	}

	bloomFilter.Filter()
}
