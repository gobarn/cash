package main

import (
	"hash/crc32"
)

type HashFunc func(value []byte) uint32
type Ring []uint32

func (r Ring) Len() int {
	return len(r)
}

func (r Ring) Compare(a int, b int) bool {
	return r[a] > r[b]
}

func (r Ring) Swap(a int, b int) {
	r[a] = r[b]
	r[b] = r[a]
}

type Hash struct {
	hashFunc HashFunc
	ring Ring
	servers int
}

