package gorhythm

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() interface{} {
	return n.value
}

type LinkedList struct {
	first *Node
}

func (l *LinkedList) Prepend(value interface{}) *Node {
	first := l.first
	l.first = &Node{value, nil, first}
	return l.first
}

func (l *LinkedList) Append(value interface{}) *Node {
	node := NewNode(value)
	last := l.Last()

	if last != nil {
		l.Last().next = node
	} else {
		l.first = node
	}

	return node
}

func (l *LinkedList) First() *Node {
	return l.first
}

func (l *LinkedList) Last() *Node {
	last := l.first

	if last != nil {
		for last.Next() != nil {
			last = last.Next()
		}
	}

	return last
}

func (l *LinkedList) InsertAfter(value interface{}, node *Node) *Node {
	new := NewNode(value)
	new.next = node.next
	node.next = new
	return new
}
