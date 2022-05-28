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

func (src IEnumerator[K, V]) AsMap(res any) {
	val := reflect.ValueOf(res)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return
	}
	m := reflect.Indirect(val)
	moveNext := src.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}
	val.Elem().Set(m)
}

// AsSlice stores the K-V pair of IEnumerable[K, V] in the value pointed to by res.
// If res is nil or not a pointer, AsResult will directly return.
func (src IEnumerator[K, V]) AsSlice(res any) {
	val := reflect.ValueOf(res)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return
	}
	slice := reflect.Indirect(val)
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		slice = reflect.Append(slice, reflect.ValueOf(v))
	}
	val.Elem().Set(slice)
}

func (src IEnumerator[K, V]) First() V {
	_, v, _ := src.Enumerate()()
	return v
}

func (src IEnumerator[K, V]) Last() V {
	var res V
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = v
	}
	return res
}
