package gorhythm

type Trie struct {
	key      rune
	Value    interface{}
	children []*Trie
}

func (t *Trie) Insert(key string, value interface{}) {
	node := t

	i := 0
	n := len(key)

	for i < n {
		child := node.getChild([]rune(key)[i])
		if child != nil {
			node = child
			i++
		} else {
			break
		}
	}

	for i < n {
		newNode := &Trie{key: []rune(key)[i], Value: value}
		node.children = append(node.children, newNode)
		node = newNode
		i++
	}
}

func (t *Trie) getChild(r rune) *Trie {
	for _, child := range t.children {
		if child.key == r {
			return child
		}
	}

	return nil
}
