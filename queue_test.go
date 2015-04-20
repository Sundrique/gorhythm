package gorhythm

import (
	. "gopkg.in/check.v1"
)

type QueueSuite struct{}

var _ = Suite(&QueueSuite{})

func (s *QueueSuite) TestQueue(c *C) {
	var queue Queue

	c.Check(queue.Dequeue(), IsNil)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	c.Check(queue.Dequeue().(int), Equals, 1)
	c.Check(queue.Dequeue().(int), Equals, 2)
	c.Check(queue.Dequeue().(int), Equals, 3)
	c.Check(queue.Dequeue(), IsNil)
}

func (s *QueueSuite) TestEnqueueMultiple(c *C) {
	var queue Queue

	c.Check(queue.Dequeue(), IsNil)

	queue.Enqueue(1, 2, 3)

	c.Check(queue.Dequeue().(int), Equals, 1)
	c.Check(queue.Dequeue().(int), Equals, 2)
	c.Check(queue.Dequeue().(int), Equals, 3)
	c.Check(queue.Dequeue(), IsNil)
}

func (s *QueueSuite) TestEnqueueSlice(c *C) {
	var queue Queue

	c.Check(queue.Dequeue(), IsNil)

	queue.Enqueue([]interface{}{1, 2, 3}...)

	c.Check(queue.Dequeue().(int), Equals, 1)
	c.Check(queue.Dequeue().(int), Equals, 2)
	c.Check(queue.Dequeue().(int), Equals, 3)
	c.Check(queue.Dequeue(), IsNil)
}
