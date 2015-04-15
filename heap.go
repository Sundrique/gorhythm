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
		for pos < len(h.data)/2 &&
			((h.data[pos] > h.data[h.left(pos)]) ||
				(h.right(pos) < len(h.data) && h.data[pos] > h.data[h.right(pos)])) {

			var newPos int
			if h.right(pos) >= len(h.data) || h.data[h.left(pos)] < h.data[h.right(pos)] {
				newPos = h.left(pos)
			} else {
				newPos = h.right(pos)
			}

			h.data[pos], h.data[newPos], pos = h.data[newPos], h.data[pos], newPos
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
			h.data[pos], h.data[h.parent(pos)] = h.data[h.parent(pos)], h.data[pos]
		}
	}
}

func (s *Heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) left(index int) int {
	return index*2 + 1
}

func (h *Heap) right(index int) int {
	return index*2 + 2
}

func (h *Heap) size() int {
	return len(h.data) - 1
}
