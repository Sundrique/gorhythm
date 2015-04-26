package gorhythm

func NewItem(value interface{}) *Item {
	return &Item{value: value}
}

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

func (n *Item) Prev() *Item {
	return n.prev
}

func (n *Item) Next() *Item {
	return n.next
}

func (n *Item) Value() interface{} {
	return n.value
}

type DoubleLinkedList struct {
	first *Item
}

func (l *DoubleLinkedList) Prepend(value interface{}) *Item {
	first := l.first
	new := &Item{value, nil, first}
	if first != nil {
		first.prev = new
	}
	l.first = new
	return l.first
}

func (l *DoubleLinkedList) Append(value interface{}) *Item {
	new := NewItem(value)
	last := l.Last()

	if last != nil {
		new.prev = l.Last()
		l.Last().next = new
	} else {
		l.first = new
	}

	return new
}

func (l *DoubleLinkedList) First() *Item {
	return l.first
}

func (l *DoubleLinkedList) Last() *Item {
	last := l.first

	if last != nil {
		for last.Next() != nil {
			last = last.Next()
		}
	}

	return last
}

func (l *DoubleLinkedList) InsertAfter(value interface{}, item *Item) *Item {
	new := NewItem(value)
	new.next = item.next
	new.prev = item
	item.next = new
	if new.next != nil {
		new.next.prev = new
	}
	return new
}
