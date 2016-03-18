package gorhythm

type Node1 map[string]Node

type Node struct {
	Comparator
	Value interface{}
	left  *Node
	right *Node
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Insert(values ...interface{}) *Node {
	var newNode *Node
	for _, value := range values {
		newNode = n.insertSingle(value)
	}

	return newNode
}

func (n *Node) insertSingle(value interface{}) *Node {
	var newNode *Node

	if n != nil {
		if n.Value == nil {
			n.Value = value
			newNode = n
		} else {
			if n.greater(value) {
				n.left = n.left.Insert(value)
				newNode = n.left
			} else {
				n.right = n.right.Insert(value)
				newNode = n.right
			}
		}
	} else {
		newNode = &Node{Value: value}
	}

	return newNode
}

func (n *Node) greater(value interface{}) bool {
	return n.Greater(n.Value, value)
}

type AVLNode struct {
	Node
}

func (n *AVLNode) Insert(values ...interface{}) *Node {
	return n.Node.Insert(values)
}
