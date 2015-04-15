package gorhythm

import (
	. "gopkg.in/check.v1"
)

type HeapSuite struct{}

var _ = Suite(&HeapSuite{})

func (s *HashTableSuite) TestHeap(c *C) {
	var heap Heap

	heap.Push(9)
	heap.Push(2)
	heap.Push(6)
	heap.Push(7)
	heap.Push(4)
	heap.Push(3)
	heap.Push(5)
	heap.Push(8)
	heap.Push(1)

	c.Check(heap.Pop(), Equals, 1)
	c.Check(heap.Pop(), Equals, 2)
	c.Check(heap.Pop(), Equals, 3)
	c.Check(heap.Pop(), Equals, 4)
	c.Check(heap.Pop(), Equals, 5)
	c.Check(heap.Pop(), Equals, 6)
	c.Check(heap.Pop(), Equals, 7)
	c.Check(heap.Pop(), Equals, 8)
	c.Check(heap.Pop(), Equals, 9)
}
