package gorhythm

type Trie struct {
	Value    interface{}
	children map[rune]*Trie
}

func (t *Trie) Insert(key string, value interface{}) {
	node := t

	i := 0
	n := len(key)

	for i < n {
		letter := []rune(key)[i]
		child := node.children[letter]
		if child != nil {
			node = child
			i++
		} else {
			break
		}
	}

	for i < n {
		newNode := &Trie{Value: value}
		letter := []rune(key)[i]
		if node.children == nil {
			node.children = make(map[rune]*Trie)
		}
		node.children[letter] = newNode
		node = newNode
		i++
	}
}

func (t *Trie) Get(key string) interface{} {
	node := t
	for _, letter := range key {
		node = node.children[letter]
		if node == nil {
			return nil
		}
	}
	return node.Value
}
