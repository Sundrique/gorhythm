package gorhythm

import (
	. "gopkg.in/check.v1"
)

type HeapSuite struct{}

var _ = Suite(&HeapSuite{})

func (s *HashTableSuite) TestHeap(c *C) {
	var heap Heap

	heap.Push(2)
	heap.Push(1)
	heap.Push(4)
	heap.Push(3)
	heap.Push(5)

	c.Check(heap.Pop(), Equals, 1)
	c.Check(heap.Pop(), Equals, 2)
	c.Check(heap.Pop(), Equals, 3)
	c.Check(heap.Pop(), Equals, 4)
	c.Check(heap.Pop(), Equals, 5)
}
