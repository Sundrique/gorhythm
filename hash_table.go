package gorhythm

import (
	"hash/fnv"
	"math"
)

type TableRecord interface {
	Key() string
}

type HashTable struct {
	size int16
	data [][]TableRecord
}

func (h *HashTable) Insert(item TableRecord) {
	key := h.calculateHash(item.Key())

	neededLen := key + 1
	if h.length() < neededLen {
		data := make([][]TableRecord, neededLen)
		copy(data, h.data)
		h.data = data
	}

	h.data[key] = append(h.data[key], item)
}

func (h *HashTable) Get(s string) interface{} {
	key := h.calculateHash(s)

	if h.length() >= key+1 {
		bucket := h.data[key]
		if len(bucket) > 0 {
			for _, item := range bucket {
				if item.Key() == s {
					return item
				}
			}
		}
	}
	return nil
}

func (h *HashTable) calculateHash(s string) int16 {
	hash := fnv.New64()
	hash.Write([]byte(s))
	return int16(hash.Sum64() % uint64(h.maxSize()))
}

func (h *HashTable) maxSize() int16 {
	if h.size > math.MaxInt16 || h.size <= 0 {
		return math.MaxInt16
	}

	return h.size
}

func (h *HashTable) length() int16 {
	return int16(len(h.data))
}
