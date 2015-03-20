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

type LinkedList struct {
	first *Item
}

func (l *LinkedList) Prepend(value interface{}) *Item {
	first := l.first
	l.first = &Item{value, nil, first}
	return l.first
}

func (l *LinkedList) Append(value interface{}) *Item {
	item := NewItem(value)
	last := l.Last()

	if last != nil {
		l.Last().next = item
	} else {
		l.first = item
	}

	return item
}

func (l *LinkedList) First() *Item {
	return l.first
}

func (l *LinkedList) Last() *Item {
	last := l.first

	if last != nil {
		for last.Next() != nil {
			last = last.Next()
		}
	}

	return last
}

func (l *LinkedList) InsertAfter(value interface{}, item *Item) *Item {
	new := NewItem(value)
	new.next = item.next
	item.next = new
	return new
}
