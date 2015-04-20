package gorhythm

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(e ...interface{}) {
	s.data = append(s.data, e...)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		maxIndex := len(s.data) - 1
		last := s.data[maxIndex]
		s.data = s.data[:maxIndex]
		return last
	}
	return nil
}
