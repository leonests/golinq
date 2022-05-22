package golinq

import "reflect"

// like C# MoveNext，iterate one item as index-value or key-value pair, check if out of boundary
type IMoveNext[K comparable, V any] func() (key K, value V, ok bool)

// generic enumerable interface
type IEnumerable[K comparable, V any] interface {
	Enumerate() IMoveNext[K, V]
}

// generic enumerable object, as a connector of different manipulations
type IEnumerator[K comparable, V any] struct {
	Enumerate func() IMoveNext[K, V]
}

type MoveNext func() (key, value any, ok bool)

// non-generic enumerable interface
type Enumerable interface {
	Enumerate() MoveNext
}

// non-generic enumerable object, as a connector of different manipulations
type Enumerator struct {
	Enumerate func() MoveNext
}

// if object is array, just convert it to slice, for example:
//
// a := [3]int{1, 2, 3}, e := FromSlice(a[:])
func FromSlice[T any](object []T) IEnumerator[int, T] {
	length := len(object)
	return IEnumerator[int, T]{
		Enumerate: func() IMoveNext[int, T] {
			index := 0
			return func() (key int, value T, ok bool) {
				ok = index < length
				if ok {
					key = index
					value = object[index]
					index++
				}
				return
			}
		},
	}
}

func FromMap[K comparable, V any](object map[K]V) IEnumerator[K, V] {
	length := len(object)
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
			index := 0
			keys := make([]K, 0, length)
			for k := range object {
				keys = append(keys, k)
			}

			return func() (key K, value V, ok bool) {
				ok = index < length
				if ok {
					key = keys[index]
					value = object[key]
					index++
				}
				return
			}
		},
	}
}

func FromChan[T any](object <-chan T) IEnumerator[int, T] {
	return IEnumerator[int, T]{
		Enumerate: func() IMoveNext[int, T] {
			index := 0
			return func() (key int, value T, ok bool) {
				key = index
				value, ok = <-object
				index++
				return
			}
		},
	}
}

func FromString(object string) IEnumerator[int, rune] {
	runes := []rune(object)
	return FromSlice(runes)
}

func Just[T any](items ...T) IEnumerator[int, T] {
	return FromSlice(items)
}

func Range(begin, count int) IEnumerator[int, int] {
	return IEnumerator[int, int]{
		Enumerate: func() IMoveNext[int, int] {
			index := 0
			return func() (key int, value int, ok bool) {
				ok = index < count
				if ok {
					key = index
					value, ok = begin, true
					index++
					begin++
				}
				return
			}
		},
	}
}

// non-generic version, there is cost penalty
// the only 1 entry for non-generic version
func From(object any) Enumerator {
	obj := reflect.ValueOf(object)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}
	switch obj.Kind() {
	case reflect.Slice, reflect.Array:
		length := obj.Len()
		return Enumerator{
			Enumerate: func() MoveNext {
				index := 0

				return func() (k, v any, ok bool) {
					ok = index < length
					if ok {
						k = index
						v = obj.Index(index).Interface()
						index++
					}
					return
				}
			},
		}
	case reflect.Map:
		length := obj.Len()
		return Enumerator{
			Enumerate: func() MoveNext {
				index := 0
				keys := obj.MapKeys() // sequence may change

				return func() (k, v any, ok bool) {
					ok = index < length
					if ok {
						k = keys[index].Interface()
						v = obj.MapIndex(keys[index]).Interface()
						index++
					}
					return
				}
			},
		}
	case reflect.Chan:
		return Enumerator{
			Enumerate: func() MoveNext {
				index := 0
				return func() (k, v any, ok bool) {
					k = index
					v, ok = <-obj.Interface().(chan any)
					index++
					return
				}
			},
		}
	case reflect.String:
		runes := []rune(object.(string))
		length := len(runes)
		return Enumerator{
			Enumerate: func() MoveNext {
				index := 0
				return func() (k, v any, ok bool) {
					ok = index < length
					if ok {
						k = index
						v = runes[index]
						index++
					}
					return
				}
			},
		}

	default:
		e := object.(Enumerable)
		return Enumerator{
			Enumerate: e.Enumerate,
		}
	}
}
