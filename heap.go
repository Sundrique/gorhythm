package gorhythm

type Heap struct {
	data []interface{}
}

func (h *Heap) Pop() interface{} {
	if len(h.data) > 0 {
		first := h.data[0]
		h.data[0] = h.data[len(h.data)-1]
		h.data = h.data[:len(h.data)-1]
		pos := 0
		for !h.isLeaf(pos) && h.disordered(pos) {
			child := h.minChildPos(pos)
			h.swap(pos, child)
			pos = child
		}
		return first
	}
	return 0
}

func (h *Heap) Push(node interface{}) {
	h.data = append(h.data, node)
	if len(h.data) > 1 {
		pos := len(h.data) - 1
		for pos > 0 && h.disordered(h.parent(pos)) {
			h.swap(pos, h.parent(pos))
			pos = h.parent(pos)
		}
	}
}

func (s *Heap) parent(pos int) int {
	return (pos - 1) / 2
}

func (h *Heap) left(pos int) int {
	return pos*2 + 1
}

func (h *Heap) right(pos int) int {
	return pos*2 + 2
}

func (h *Heap) isLeaf(pos int) bool {
	return pos >= len(h.data)/2
}

func (h *Heap) disordered(pos int) bool {
	return (h.left(pos) < len(h.data) && h.greater(pos, h.left(pos))) || (h.right(pos) < len(h.data) && h.greater(pos, h.right(pos)))
}

func (h *Heap) minChildPos(pos int) int {
	if h.right(pos) >= len(h.data) || h.greater(h.right(pos), h.left(pos)) {
		return h.left(pos)
	} else {
		return h.right(pos)
	}
}

func (h *Heap) swap(src int, dst int) {
	h.data[src], h.data[dst] = h.data[dst], h.data[src]
}

func (h *Heap) greater(posA int, posB int) bool {
	a := h.data[posA]
	b := h.data[posB]
	switch a.(type) {
	default:
		return false
	case Comparable:
		return a.(Comparable).Compare(b) > 0
	case int:
		return a.(int) > b.(int)
	case int8:
		return a.(int8) > b.(int8)
	case int16:
		return a.(int16) > b.(int16)
	case int32:
		return a.(int32) > b.(int32)
	case int64:
		return a.(int64) > b.(int64)
	case uint:
		return a.(uint) > b.(uint)
	case uint8:
		return a.(uint8) > b.(uint8)
	case uint16:
		return a.(uint16) > b.(uint16)
	case uint32:
		return a.(uint32) > b.(uint32)
	case uint64:
		return a.(uint64) > b.(uint64)
	case float32:
		return a.(float32) > b.(float32)
	case float64:
		return a.(float64) > b.(float64)
	case string:
		return a.(string) > b.(string)
	}
}
