package gorhythm

import (
	. "gopkg.in/check.v1"
)

type BTreeSuite struct{}

var _ = Suite(&BTreeSuite{})

func (s *BTreeSuite) TestNodeConstructor(c *C) {
	node := NewNode(1)

	c.Check(node.Value, Equals, 1)
	c.Check(node.Right, IsNil)
	c.Check(node.Left, IsNil)
}

func (s *BTreeSuite) TestInsert(c *C) {
	tree := NewNode(3)

	c.Check(tree.Value, Equals, 3)
	c.Check(tree.Right, IsNil)
	c.Check(tree.Left, IsNil)

	tree.Insert(2)

	c.Assert(tree.Left, Not(IsNil))
	c.Check(tree.Left.Value, Equals, 2)

	tree.Insert(4)

	c.Assert(tree.Right, Not(IsNil))
	c.Check(tree.Right.Value, Equals, 4)
}
