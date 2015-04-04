package gorhythm

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(e interface{}) {
	q.data = append(q.data, e)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.data) > 0 {
		first := q.data[0]
		q.data = q.data[1:]
		return first
	}
	return nil
}
