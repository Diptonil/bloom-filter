package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)

var hasher hash.Hash32
var dataset []string
var datasetReplicated map[string]bool
var datasetUnique map[string]bool

const SIZE uint32 = 1000

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

func (bloomFilter BloomFilter) ShowFilter() {
	fmt.Println(bloomFilter.filter)
}

func NewBloomFilter(size uint32) BloomFilter {
	filter := make([]byte, size)
	return BloomFilter{filter, size}
}

func Hash(key string, size uint32) uint32 {
	hasher.Write([]byte(key))
	result := hasher.Sum32() % size
	hasher.Reset()
	return result
}

func GenerateDataset(size uint32) {
	for iterator := uint32(0); iterator < size/2; iterator++ {
		id := uuid.New().String()
		dataset = append(dataset, id)
		datasetUnique[id] = true
	}
	for iterator := uint32(0); iterator < size/2; iterator++ {
		id := uuid.New().String()
		dataset = append(dataset, id)
		datasetReplicated[id] = true
	}
}

func AddToBloomFilter(bloomFilter BloomFilter) {
	for key := range datasetUnique {
		bloomFilter.Add(key)
	}
}

func CheckRate(bloomFilter BloomFilter, falsePositives int) {
	for _, key := range dataset {
		if bloomFilter.Exists(key) && datasetReplicated[key] {
			falsePositives++
		}
	}
	fmt.Println(float32(falsePositives) / float32(len(dataset)))
}

func init() {
	hasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func main() {
	bloomFilter := NewBloomFilter(SIZE)
	falsePositives := 0

	GenerateDataset(bloomFilter.size)
	AddToBloomFilter(bloomFilter)
	CheckRate(bloomFilter, falsePositives)
}
