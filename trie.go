package gorhythm

type Trie struct {
	Key      rune
	Value    interface{}
	Children []*Trie
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
		newNode := &Trie{Key: []rune(key)[i], Value: value}
		node.Children = append(node.Children, newNode)
		node = newNode
		i++
	}
}

func (t *Trie) getChild(r rune) *Trie {
	for _, child := range t.Children {
		if child.Key == r {
			return child
		}
	}

	return nil
}
