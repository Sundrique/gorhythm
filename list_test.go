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
	list := LinkedList{}

	first := list.Prepend(1)

	c.Check(list.First().Value(), Equals, 1)
	c.Check(list.First().Next(), IsNil)

	list.Prepend(2)

	c.Check(list.First().Value(), Equals, 2)
	c.Check(list.First().Next(), DeepEquals, first)
}

func (s *ListSuite) TestAppend(c *C) {
	list := LinkedList{}

	list.Append(1)

	c.Check(list.Last().Value(), Equals, 1)
	c.Check(list.Last().Next(), IsNil)

	list.Append(2)

	c.Check(list.Last().Value(), Equals, 2)
	c.Check(list.Last().Next(), IsNil)
}

func (s *ListSuite) TestInsertAfter(c *C) {
	list := LinkedList{}

	list.Append(1)
	new := list.InsertAfter(2, list.First())

	c.Check(list.First().Next(), DeepEquals, new)
	c.Check(list.Last().Value(), Equals, 2)
}
