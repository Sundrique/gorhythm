package gorhythm

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

func (n *Node) Insert(value int) {
	if value-n.Value <= 0 {
		if n.Left == nil {
			n.Left = NewNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(value)
		} else {
			n.Right.Insert(value)
		}
	}
}
