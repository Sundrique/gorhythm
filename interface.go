package gorhythm

type Comparable interface {
	Compare(value interface{}) int
}
