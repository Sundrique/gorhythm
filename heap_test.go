package gorhythm

import (
	. "gopkg.in/check.v1"
)

type HeapSuite struct{}

var _ = Suite(&HeapSuite{})

func (s *HashTableSuite) TestPush(c *C) {
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

func (s *HashTableSuite) TestPushMultiple(c *C) {
	var heap Heap

	heap.Push(9, 11, 42, 23, 0, 88, 15, 404, 6)

	c.Check(heap.Pop(), Equals, 0)
	c.Check(heap.Pop(), Equals, 6)
	c.Check(heap.Pop(), Equals, 9)
	c.Check(heap.Pop(), Equals, 11)
	c.Check(heap.Pop(), Equals, 15)
	c.Check(heap.Pop(), Equals, 23)
	c.Check(heap.Pop(), Equals, 42)
	c.Check(heap.Pop(), Equals, 88)
	c.Check(heap.Pop(), Equals, 404)
}

func (s *HashTableSuite) TestPushSlice(c *C) {
	var heap Heap

	heap.Push([]interface{}{100, 21, 13, 304, 283, 111}...)

	c.Check(heap.Pop(), Equals, 13)
	c.Check(heap.Pop(), Equals, 21)
	c.Check(heap.Pop(), Equals, 100)
	c.Check(heap.Pop(), Equals, 111)
	c.Check(heap.Pop(), Equals, 283)
	c.Check(heap.Pop(), Equals, 304)
}
