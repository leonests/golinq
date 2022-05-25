package golinq

import (
	"sort"
)

func (src IEnumerator[K, V]) OrderBy(selector func(K, V) any) IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
			tmp := src.sort(selector)
			index, length := 0, len(tmp)
			return func() (k K, v V, ok bool) {
				ok = index < length
				if ok {
					k, v = tmp[index].key, tmp[index].value
					index++
				}
				return
			}
		},
	}
}

func (src IEnumerator[K, V]) sort(selector func(K, V) any) (items sorters[K, V]) {
	moveNext := src.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		items = append(items, sorter[K, V]{k, v, selector(k, v)})
	}
	if len(items) == 0 {
		return
	}
	sort.Sort(items)
	return
}
