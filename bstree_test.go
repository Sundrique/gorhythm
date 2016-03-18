package gorhythm

import (
	. "gopkg.in/check.v1"
)

type BTreeSuite struct{}

var _ = Suite(&BTreeSuite{})

func (s *BTreeSuite) TestInsert(c *C) {
	var tree Node

	tree.Insert(3)

	c.Check(tree.Value, Equals, 3)
	c.Check(tree.Right(), IsNil)
	c.Check(tree.Left(), IsNil)

	tree.Insert(2)

	c.Assert(tree.Left(), Not(IsNil))
	c.Check(tree.Left().Value, Equals, 2)

	tree.Insert(4)

	c.Assert(tree.Right(), Not(IsNil))
	c.Check(tree.Right().Value, Equals, 4)
}

func (s *BTreeSuite) TestInsertMultiple(c *C) {
	var tree Node

	tree.Insert(3, 2, 4)

	c.Check(tree.Value, Equals, 3)
	c.Assert(tree.Left(), Not(IsNil))
	c.Check(tree.Left().Value, Equals, 2)
	c.Assert(tree.Right(), Not(IsNil))
	c.Check(tree.Right().Value, Equals, 4)
}

func (s *BTreeSuite) TestInsertSlice(c *C) {
	var tree Node

	tree.Insert([]interface{}{3, 2, 4}...)

	c.Check(tree.Value, Equals, 3)
	c.Assert(tree.Left(), Not(IsNil))
	c.Check(tree.Left().Value, Equals, 2)
	c.Assert(tree.Right(), Not(IsNil))
	c.Check(tree.Right().Value, Equals, 4)
}

func (s *BTreeSuite) TestInsertSliceAVL(c *C) {
	var tree AVLNode

	tree.Insert([]interface{}{3, 2, 4}...)

	c.Check(tree.Value, Equals, 3)
	c.Assert(tree.Left(), Not(IsNil))
	c.Check(tree.Left().Value, Equals, 2)
	c.Assert(tree.Right(), Not(IsNil))
	c.Check(tree.Right().Value, Equals, 4)
}
