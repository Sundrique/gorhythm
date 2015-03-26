package gorhythm

type Trie struct {
	Value    rune
	Children []*Trie
}

func (t *Trie) Insert(s string) {
	node := t

	i := 0
	n := len(s)

	for i < n {
		child := node.getChild([]rune(s)[i])
		if child != nil {
			node = child
			i++
		} else {
			break
		}
	}

	for i < n {
		newNode := &Trie{Value: []rune(s)[i]}
		node.Children = append(node.Children, newNode)
		node = newNode
		i++
	}
}

func (t *Trie) getChild(r rune) *Trie {
	for _, child := range t.Children {
		if child.Value == r {
			return child
		}
	}

	return nil
}
