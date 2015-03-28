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
	c.Assert(len(trie.children['a'].children), Equals, 2)
	c.Assert(len(trie.children['d'].children), Equals, 1)

	c.Check(trie.children['a'].children['b'].Value.(int), Equals, 1)
	c.Check(trie.children['a'].children['c'].Value.(int), Equals, 2)
	c.Check(trie.children['d'].children['e'].Value.(int), Equals, 3)
}

func (s *TrieSuite) TestGet(c *C) {
	var trie Trie

	trie.Insert("ab", 1)
	trie.Insert("ac", 2)
	trie.Insert("de", 3)

	c.Assert(trie.Get("ab"), Not(IsNil))
	c.Assert(trie.Get("ac"), Not(IsNil))
	c.Assert(trie.Get("de"), Not(IsNil))

	c.Check(trie.Get("abc"), IsNil)

	c.Check(trie.Get("ab").(int), Equals, 1)
	c.Check(trie.Get("ac").(int), Equals, 2)
	c.Check(trie.Get("de").(int), Equals, 3)
}
