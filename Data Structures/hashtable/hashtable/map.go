package hashtable

import (
	"fmt"
	"hash/maphash"
)

type HashTable[K comparable, V any] struct {
	buckets []*Bucket[K, V]
	size int
	seed maphash.Seed
}

type HashTableOption[K comparable, V any] func(*HashTable[K, V])

func WithSize[K comparable, V any](size int) HashTableOption[K, V] {
	return func(ht *HashTable[K, V]) {
		ht.size = size
		ht.buckets = make([]*Bucket[K, V], size)
	}
}

func NewHashTable[K comparable, V any](opts ...HashTableOption[K, V]) *HashTable[K, V] {
	ht := &HashTable[K, V]{
		buckets: make([]*Bucket[K, V], 16), // default size
		size: 16,
			seed: maphash.MakeSeed(),
	}

	for _, opt := range opts {
		opt(ht)
	}
	
	for i := 0; i < ht.size; i++ {
		ht.buckets[i] = &Bucket[K, V]{}
		ht.buckets[i].head = nil
	}
	return ht
}

func (ht *HashTable[K, V]) Hash(key K) int {
	var h maphash.Hash
	h.SetSeed(ht.seed)
	fmt.Fprintf(&h, "%v", key)
	hashValue := h.Sum64()

	fmt.Printf("Value: %v\t hashValue: %v\t index: %v\n", key, hashValue, int(hashValue & uint64(ht.size-1)))
	// bitwise AND operation with size-1 to get the index
	return int(hashValue & uint64(ht.size-1))
}


func (ht *HashTable[K, V]) Insert(key K, value V) bool {
	index := ht.Hash(key)
	bucket := ht.buckets[index]

	start := bucket.head

	newNode := &BucketNode[K, V]{
		key: key,
		value: value,
		next: nil,
	}

	if start == nil {
		bucket.head = newNode
		return true
	}

	for start.next != nil {
		start = start.next
		if start.key == key {
			start.value = value
			return true
		}
	}

	start.next = newNode
	return true
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index := ht.Hash(key)
	bucket := ht.buckets[index]

	var val V
	start := bucket.head
	for start != nil {
		if start.key == key {
			val = start.value
			return val, true
		}
		start = start.next
	}
	return val, false
}