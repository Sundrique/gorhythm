package gorhythm

import (
	. "gopkg.in/check.v1"
)

type StackSuite struct{}

var _ = Suite(&StackSuite{})

func (s *StackSuite) TestStack(c *C) {
	var stack Stack

	c.Check(stack.Pop(), IsNil)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	c.Check(stack.Pop().(int), Equals, 3)
	c.Check(stack.Pop().(int), Equals, 2)
	c.Check(stack.Pop().(int), Equals, 1)
	c.Check(stack.Pop(), IsNil)
}

func (s *StackSuite) TestPushMultiple(c *C) {
	var stack Stack

	c.Check(stack.Pop(), IsNil)

	stack.Push(1, 2, 3)

	c.Check(stack.Pop().(int), Equals, 3)
	c.Check(stack.Pop().(int), Equals, 2)
	c.Check(stack.Pop().(int), Equals, 1)
	c.Check(stack.Pop(), IsNil)
}

func (s *StackSuite) TestPushSlice(c *C) {
	var stack Stack

	c.Check(stack.Pop(), IsNil)

	stack.Push([]interface{}{1, 2, 3}...)

	c.Check(stack.Pop().(int), Equals, 3)
	c.Check(stack.Pop().(int), Equals, 2)
	c.Check(stack.Pop().(int), Equals, 1)
	c.Check(stack.Pop(), IsNil)
}
