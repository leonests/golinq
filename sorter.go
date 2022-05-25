package golinq

import "reflect"

type sorter[K comparable, V any] struct {
	key   K
	value V
	merge any
}
type sorters[K comparable, V any] []sorter[K, V]

func (s sorters[K, V]) Len() int {
	return len(s)
}

func (s sorters[K, V]) Less(i, j int) bool {
	x, y := s[i].merge, s[j].merge
	compare := getInternalCompare(x)
	switch compare(x, y) {
	case -1:
		return true
	case 1:
		return false
	default:
		return false
	}

}
func (s sorters[K, V]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Comparable is an interface that self-defined collection element has to implemented
// when using order-wise linq func
type Comparable interface {
	CompareTo(Comparable) int
}

func getInternalCompare(item any) func(any, any) int {
	switch item.(type) {
	case int, int8, int16, int32, int64:
		return func(x, y any) int {
			vx, vy := reflect.ValueOf(x).Int(), reflect.ValueOf(y).Int()
			switch {
			case vx > vy:
				return 1
			case vx < vy:
				return -1
			default:
				return 0
			}
		}
	case uint, uint8, uint16, uint32, uint64:
		return func(x, y any) int {
			vx, vy := reflect.ValueOf(x).Uint(), reflect.ValueOf(y).Uint()
			switch {
			case vx > vy:
				return 1
			case vx < vy:
				return -1
			default:
				return 0
			}
		}
	case float32, float64:
		return func(x, y any) int {
			vx, vy := reflect.ValueOf(x).Float(), reflect.ValueOf(y).Float()
			switch {
			case vx > vy:
				return 1
			case vx < vy:
				return -1
			default:
				return 0
			}
		}
	case string:
		return func(x, y any) int {
			vx, vy := x.(string), y.(string)
			switch {
			case vx > vy:
				return 1
			case vx < vy:
				return -1
			default:
				return 0
			}
		}
	case bool:
		return func(x, y any) int {
			vx, vy := x.(bool), y.(bool)
			switch {
			case vx == vy:
				return 0
			case vx:
				return 1
			default:
				return -1
			}
		}
	default:
		return func(x, y any) int {
			vx, vy := x.(Comparable), y.(Comparable)
			return vx.CompareTo(vy)
		}
	}
}
