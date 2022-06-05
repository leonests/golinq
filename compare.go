package golinq

import "reflect"

type sorter[K, V any] struct {
	key   K
	value V
}
type multiSorter[K, V any] struct {
	sorters []sorter[K, V]
	conds   []sortCond[K, V]
}

type sortCond[K, V any] struct {
	selector func(K, V) any
	desc     bool
	compare  compareFunc
}
type compareFunc func(x, y any) int

func (s multiSorter[K, V]) Len() int {
	return len(s.sorters)
}
func (s multiSorter[K, V]) Swap(i, j int) {
	s.sorters[i], s.sorters[j] = s.sorters[j], s.sorters[i]
}

func (s multiSorter[K, V]) Less(i, j int) bool {
	x, y := s.sorters[i], s.sorters[j]
	for _, cond := range s.conds {
		selector := cond.selector
		switch cond.compare(selector(x.key, x.value), selector(y.key, y.value)) {
		case -1:
			return !cond.desc
		case 1:
			return cond.desc
		}
	}
	return false //all the same, keep the original order
}

// Comparable is an interface that self-defined collection element has to implemented
// when using order-wise linq func
type Comparable interface {
	CompareTo(Comparable) int
}

func getCompareFunc(item any) compareFunc {
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

func convert2Int64(number any) func(x any) int64 {
	switch number.(type) {
	case int, int8, int16, int32, int64:
		return func(x any) int64 {
			return reflect.ValueOf(x).Int()
		}
	case uint, uint8, uint16, uint32, uint64:
		return func(x any) int64 {
			return int64(reflect.ValueOf(x).Uint())
		}
	}
	return func(x any) int64 {
		return x.(int64)
	}
}

func convert2Float64(number any) func(number any) float64 {
	switch number.(type) {
	case int, int8, int16, int32, int64:
		return func(x any) float64 {
			return float64(reflect.ValueOf(x).Int())
		}
	case uint, uint8, uint16, uint32, uint64:
		return func(x any) float64 {
			return float64(reflect.ValueOf(x).Uint())
		}
	case float32, float64:
		return func(x any) float64 {
			return reflect.ValueOf(x).Float()
		}
	}
	return func(x any) float64 {
		return x.(float64)
	}
}
