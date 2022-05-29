package golinq

import (
	"reflect"
)

func (src Enumerator[K, V]) ToSlice() []V {
	res := make([]V, 0)
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = append(res, v)
	}
	return res
}

func (src Enumerator[K, V]) ToChannel() <-chan V {
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

// AsMap stores the K-V pair of IEnumerator[K, V] in the value pointed to by map res.
// If res is nil or not a pointer, it will directly return.
//
// ToMap is not supported because the key contraint of map is comparable, and
// there is no way to convert any to comparable; Also if comparable constaint
// is applied to type K, there would be obtacles for type chain, so there is
// definitely compromise here, hope new golang generic features could help with this.
func (src Enumerator[K, V]) AsMap(res any) {
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

// AsSlice stores the K-V pair of IEnumerator[K, V] in the value pointed to by slice res.
// If res is nil or not a pointer, it will directly return.
func (src Enumerator[K, V]) AsSlice(res any) {
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

func (src Enumerator[K, V]) First() V {
	_, v, _ := src.Enumerate()()
	return v
}

func (src Enumerator[K, V]) Last() V {
	var res V
	moveNext := src.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = v
	}
	return res
}

func (src Enumerator[K, V]) Count() int {
	var res int
	moveNext := src.Enumerate()
	for _, _, ok := moveNext(); ok; _, _, ok = moveNext() {
		res++
	}
	return res
}

func (src Enumerator[K, V]) CountBy(predict func(K, V) bool) int {
	var res int
	moveNext := src.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		if predict(k, v) {
			res++
		}
	}
	return res
}

func (src Enumerator[K, V]) Max() V {
	moveNext := src.Enumerate()
	_, max, ok := moveNext()
	if !ok {
		return max
	}
	compare := getInternalCompare(max)
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		if compare(v, max) == 1 {
			max = v
		}
	}
	return max
}

func (src Enumerator[K, V]) Min() V {
	moveNext := src.Enumerate()
	_, min, ok := moveNext()
	if !ok {
		return min
	}
	compare := getInternalCompare(min)
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		if compare(v, min) == -1 {
			min = v
		}
	}
	return min
}

func (src Enumerator[K, V]) Sum2Int() int64 {
	moveNext := src.Enumerate()
	_, first, ok := moveNext()
	if !ok {
		return 0
	}
	converter := convert2Int64(first)
	sum := converter(first)
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		sum = sum + converter(v)
	}
	return sum
}

func (src Enumerator[K, V]) Sum2Float() float64 {
	moveNext := src.Enumerate()
	_, first, ok := moveNext()
	if !ok {
		return 0
	}
	converter := convert2Float64(first)
	sum := converter(first)
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		sum = sum + converter(v)
	}
	return sum
}
