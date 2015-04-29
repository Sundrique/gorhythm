package gorhythm

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

type LinkedList interface {
	First() *Item
	Prepend(value interface{}) *Item
	Append(value interface{}) *Item
	InsertAfter(value interface{}, item *Item) *Item
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
	new := &Item{value: value}

	if l.first != nil {
		new.prev = l.first.prev
		l.first.prev = new
		new.next = l.first
	}
	l.first = new

	return new
}

func (l *DoubleLinkedList) Append(value interface{}) *Item {
	return appendItem(l, value)
}

func (l *DoubleLinkedList) InsertAfter(value interface{}, item *Item) *Item {
	new := &Item{value: value}
	new.next = item.next
	new.prev = item
	item.next = new
	if new.next != nil {
		new.next.prev = new
	}

	return new
}

func (l *DoubleLinkedList) First() *Item {
	return l.first
}

func (l *DoubleLinkedList) last() *Item {
	last := l.first

	if last != nil {
		for last.Next() != nil {
			last = last.Next()
		}
	}

	return last
}

type CircularLinkedList struct {
	DoubleLinkedList
}

func (l *CircularLinkedList) Prepend(value interface{}) *Item {
	new := l.DoubleLinkedList.Prepend(value)

	if new.next == nil {
		new.next = new
		new.prev = new
	}

	return new
}

func (l *CircularLinkedList) Append(value interface{}) *Item {
	return appendItem(l, value)
}

func appendItem(l LinkedList, value interface{}) *Item {
	var new *Item
	if l.First() == nil {
		new = l.Prepend(value)
	} else {
		new = l.InsertAfter(value, lastItem(l))
	}

	return new
}

func lastItem(l LinkedList) *Item {
	last := l.First()

	if last != nil {
		for last.Next() != nil {
			last = last.Next()
		}
	}

	return last
}
