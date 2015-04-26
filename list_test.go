package gorhythm

import (
	. "gopkg.in/check.v1"
)

type ListSuite struct{}

var _ = Suite(&ListSuite{})

func (s *ListSuite) TestElementConstructor(c *C) {

	item := NewItem(1)

	c.Check(item.Value(), Equals, 1)
	c.Check(item.Next(), IsNil)
}

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
	c.Check(list.Last().Prev(), DeepEquals, first)
	c.Check(list.Last().Value(), Equals, 2)

	third := list.InsertAfter(3, list.First())

	c.Check(list.Last().Prev(), DeepEquals, third)
	c.Check(third.Prev(), DeepEquals, first)
	c.Check(third.Next(), DeepEquals, second)
}
