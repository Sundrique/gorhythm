package gorhythm

import (
	. "gopkg.in/check.v1"
)

type ListSuite struct{}

var _ = Suite(&ListSuite{})

func (s *ListSuite) TestPrepend(c *C) {
	list := DoubleLinkedList{}

	first := list.Prepend(1)

	c.Check(list.First().Value(), Equals, 1)
	c.Check(list.First().Next(), IsNil)

	list.Prepend(2)

	c.Check(list.First().Value(), Equals, 2)
	c.Check(list.First().Next(), DeepEquals, first)
}

func (s *ListSuite) TestAppend(c *C) {
	list := DoubleLinkedList{}

	first := list.Prepend(1)

	c.Check(list.First().Value(), Equals, 1)
	c.Check(list.First().Next(), IsNil)

	list.Prepend(2)
	list.Append(3)

	c.Assert(first.Next(), NotNil)
	c.Check(first.Next().Value(), Equals, 3)
	c.Assert(first.Prev(), NotNil)
	c.Check(first.Prev().Value(), Equals, 2)
}

func (s *ListSuite) TestInsertAfter(c *C) {
	list := DoubleLinkedList{}

	first := list.Append(1)
	second := list.InsertAfter(2, list.First())

	c.Check(list.First().Next(), DeepEquals, second)
	c.Check(second.Prev(), DeepEquals, first)
	c.Check(second.Value(), Equals, 2)

	third := list.InsertAfter(3, list.First())

	c.Check(second.Prev(), DeepEquals, third)
	c.Check(third.Prev(), DeepEquals, first)
	c.Check(third.Next(), DeepEquals, second)
}

func (s *ListSuite) TestCircularLinkedList(c *C) {
	list := CircularLinkedList{}

	list.Append(1)

	c.Check(list.First().Next(), DeepEquals, list.First())
	c.Check(list.First().Prev(), DeepEquals, list.First())

	list.Prepend(2)

	c.Check(list.First().Value(), Equals, 2)
	c.Check(list.First().Prev().Value(), Equals, 1)
	c.Check(list.First().Next().Value(), Equals, 1)
}
