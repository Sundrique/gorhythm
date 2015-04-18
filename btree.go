package gorhythm

type Node struct {
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

func (n *Node) Insert(value interface{}) *Node {
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
	switch n.Value.(type) {
	default:
		return false
	case Comparable:
		return n.Value.(Comparable).Compare(value) > 0
	case int:
		return n.Value.(int) > value.(int)
	case int8:
		return n.Value.(int8) > value.(int8)
	case int16:
		return n.Value.(int16) > value.(int16)
	case int32:
		return n.Value.(int32) > value.(int32)
	case int64:
		return n.Value.(int64) > value.(int64)
	case uint:
		return n.Value.(uint) > value.(uint)
	case uint8:
		return n.Value.(uint8) > value.(uint8)
	case uint16:
		return n.Value.(uint16) > value.(uint16)
	case uint32:
		return n.Value.(uint32) > value.(uint32)
	case uint64:
		return n.Value.(uint64) > value.(uint64)
	case float32:
		return n.Value.(float32) > value.(float32)
	case float64:
		return n.Value.(float64) > value.(float64)
	case string:
		return n.Value.(string) > value.(string)
	}
}
