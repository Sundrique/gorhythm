package gorhythm

type Compare interface {
	Compare(value interface{}) int
}

type Heap struct {
	data []int
}

func (h *Heap) Pop() int {
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

func (h *Heap) Push(node int) {
	h.data = append(h.data, node)
	if len(h.data) > 1 {
		pos := len(h.data) - 1
		for pos > 0 && h.data[pos] < h.data[h.parent(pos)] {
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
	return h.data[pos] > h.data[h.left(pos)] || (h.right(pos) < len(h.data) && h.data[pos] > h.data[h.right(pos)])
}

func (h *Heap) minChildPos(pos int) int {
	if h.right(pos) >= len(h.data) || h.data[h.left(pos)] < h.data[h.right(pos)] {
		return h.left(pos)
	} else {
		return h.right(pos)
	}
}

func (h *Heap) swap(src int, dst int) {
	h.data[src], h.data[dst] = h.data[dst], h.data[src]
}
