package gorhythm

import (
	. "gopkg.in/check.v1"
)

type TrieSuite struct{}

var _ = Suite(&TrieSuite{})

func (s *TrieSuite) TestInsert(c *C) {
	var trie Trie

	trie.Insert("ab")

	trie.Insert("ac")

	trie.Insert("de")

	c.Assert(len(trie.Children), Equals, 2)
	c.Check(string(trie.Children[0].Value), Equals, "a")
	c.Check(string(trie.Children[1].Value), Equals, "d")

	c.Assert(len(trie.Children[0].Children), Equals, 2)
	c.Check(string(trie.Children[0].Children[0].Value), Equals, "b")
	c.Check(string(trie.Children[0].Children[1].Value), Equals, "c")
}
