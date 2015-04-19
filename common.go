package gorhythm

type Comparable interface {
	Compare(value interface{}) int
}

type Comparator struct {
}

func (c Comparator) Greater(one interface{}, another interface{}) bool {
	switch one.(type) {
	default:
		return false
	case Comparable:
		return one.(Comparable).Compare(another) > 0
	case int:
		return one.(int) > another.(int)
	case int8:
		return one.(int8) > another.(int8)
	case int16:
		return one.(int16) > another.(int16)
	case int32:
		return one.(int32) > another.(int32)
	case int64:
		return one.(int64) > another.(int64)
	case uint:
		return one.(uint) > another.(uint)
	case uint8:
		return one.(uint8) > another.(uint8)
	case uint16:
		return one.(uint16) > another.(uint16)
	case uint32:
		return one.(uint32) > another.(uint32)
	case uint64:
		return one.(uint64) > another.(uint64)
	case float32:
		return one.(float32) > another.(float32)
	case float64:
		return one.(float64) > another.(float64)
	case string:
		return one.(string) > another.(string)
	}
}
