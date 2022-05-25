package golinq

import (
	"reflect"
)

func (src IEnumerator[K, V]) ToSlice() []V {
	res := make([]V, 0)
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = append(res, v)
	}
	return res
}

func (src IEnumerator[K, V]) ToMap() map[K]V {
	res := make(map[K]V, 0)
	moveNext := src.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		res[k] = v
	}
	return res
}

func (src IEnumerator[K, V]) ToChannel() <-chan V {
	res := make(chan V)
	moveNext := src.Enumerate()
	go func() {
		for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
			res <- v
		}
		close(res)
	}()
	return res
}

// Result stores the kv pair IEnumerable[K, V] in the value pointed to by res.
// If res is nil or not a pointer, Result will directly return.
func (src IEnumerator[K, V]) Result(res any) {
	val := reflect.ValueOf(res)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return
	}
	switch val.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		moveNext := src.Enumerate()
		slice := reflect.Indirect(val)
		for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
			slice = reflect.Append(slice, reflect.ValueOf(v))
		}
		val.Elem().Set(slice)
	case reflect.Map:
		moveNext := src.Enumerate()
		m := reflect.MakeMap(
			val.Elem().Type(),
		)
		for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
			m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		val.Elem().Set(m)
	default:
		return
	}
}

func (src Enumerator) ToSlice() []any {
	res := make([]any, 0)
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = append(res, v)
	}
	return res
}
