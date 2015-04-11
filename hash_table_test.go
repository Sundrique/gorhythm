package gorhythm

import (
	. "gopkg.in/check.v1"
	"strconv"
)

type KeyedInt int

func (k KeyedInt) Key() string {
	return strconv.Itoa(int(k))
}

type HashTableSuite struct{}

var _ = Suite(&HashTableSuite{})

func (s *HashTableSuite) TestHashTable(c *C) {
	var t HashTable

	c.Assert(t.Get("1"), IsNil)

	t.Insert(KeyedInt(1))
	t.Insert(KeyedInt(2))
	t.Insert(KeyedInt(3))

	c.Check(int(t.Get("1").(KeyedInt)), Equals, 1)
	c.Check(int(t.Get("2").(KeyedInt)), Equals, 2)
	c.Check(int(t.Get("3").(KeyedInt)), Equals, 3)
	c.Check(t.Get("4"), IsNil)

}
