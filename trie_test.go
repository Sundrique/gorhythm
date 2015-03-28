package gorhythm

import (
	. "gopkg.in/check.v1"
)

type TrieSuite struct{}

var _ = Suite(&TrieSuite{})

func (s *TrieSuite) TestInsert(c *C) {
	var trie Trie

	trie.Insert("ab", 1)

	trie.Insert("ac", 2)

	trie.Insert("de", 3)

	c.Assert(len(trie.Children), Equals, 2)
	c.Assert(len(trie.Children[0].Children), Equals, 2)
	c.Check(trie.Children[0].Children[0].Value.(int), Equals, 1)
	c.Check(trie.Children[0].Children[1].Value.(int), Equals, 2)

	c.Assert(len(trie.Children[0].Children), Equals, 2)
	c.Check(trie.Children[1].Children[0].Value.(int), Equals, 3)
}
