package gorhythm

type Heap struct {
	Comparator
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

func (h *Heap) Push(nodes ...interface{}) {
	for _, node := range nodes {
		h.pushSingle(node)
	}
}

func (h *Heap) pushSingle(node interface{}) {
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
	return h.Greater(h.data[posA], h.data[posB])
}
