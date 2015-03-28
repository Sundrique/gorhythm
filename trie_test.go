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

	c.Assert(len(trie.children), Equals, 2)
	c.Assert(len(trie.children[0].children), Equals, 2)
	c.Assert(len(trie.children[1].children), Equals, 1)

	c.Check(trie.children[0].children[0].Value.(int), Equals, 1)
	c.Check(trie.children[0].children[1].Value.(int), Equals, 2)
	c.Check(trie.children[1].children[0].Value.(int), Equals, 3)
}
