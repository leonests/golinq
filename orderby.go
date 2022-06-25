package golinq

import (
	"sort"
)

type OrderEnumerator[K, V any] struct {
	Enumerator[K, V]                  // inherit all the methods of Enumerator
	source           Enumerator[K, V] // keep the source
	conditions       []sortCond[K, V] // order conditions
}

func (src Enumerator[K, V]) OrderBy(selector func(K, V) any) OrderEnumerator[K, V] {
	return OrderEnumerator[K, V]{
		source:     src,
		conditions: []sortCond[K, V]{{selector: selector, desc: false}},
		Enumerator: Enumerator[K, V]{
			Enumerate: func() MoveNext[K, V] {
				tmp := src.sort([]sortCond[K, V]{{selector: selector, desc: false}})
				index, length := 0, len(tmp.sorters)
				return func() (k K, v V, ok bool) {
					ok = index < length
					if ok {
						k, v = tmp.sorters[index].key, tmp.sorters[index].value
						index++
					}
					return
				}
			},
		},
	}
}
func (src Enumerator[K, V]) OrderByDescending(selector func(K, V) any) OrderEnumerator[K, V] {
	return OrderEnumerator[K, V]{
		source:     src,
		conditions: []sortCond[K, V]{{selector: selector, desc: true}},
		Enumerator: Enumerator[K, V]{
			Enumerate: func() MoveNext[K, V] {
				tmp := src.sort([]sortCond[K, V]{{selector: selector, desc: true}})
				index, length := 0, len(tmp.sorters)
				return func() (k K, v V, ok bool) {
					ok = index < length
					if ok {
						k, v = tmp.sorters[index].key, tmp.sorters[index].value
						index++
					}
					return
				}
			},
		},
	}
}

func (src OrderEnumerator[K, V]) ThenBy(selector func(K, V) any) OrderEnumerator[K, V] {
	return OrderEnumerator[K, V]{
		source:     src.source,
		conditions: append(src.conditions, sortCond[K, V]{selector: selector, desc: false}),
		Enumerator: Enumerator[K, V]{
			Enumerate: func() MoveNext[K, V] {
				tmp := src.source.sort(append(src.conditions, sortCond[K, V]{selector: selector, desc: false}))
				index, length := 0, len(tmp.sorters)
				return func() (k K, v V, ok bool) {
					ok = index < length
					if ok {
						k, v = tmp.sorters[index].key, tmp.sorters[index].value
						index++
					}
					return
				}
			},
		},
	}
}

func (src Enumerator[K, V]) sort(conds []sortCond[K, V]) (multi multiSorter[K, V]) {
	moveNext := src.Enumerate()
	sorters := make([]sorter[K, V], 0)
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		sorters = append(sorters, sorter[K, V]{k, v})
	}
	if len(sorters) == 0 {
		return
	}

	for i := 0; i < len(conds); i++ {
		merge := conds[i].selector(sorters[0].key, sorters[0].value)
		conds[i].compare = getCompareFunc(merge)
	}

	multi = multiSorter[K, V]{
		sorters: sorters,
		conds:   conds,
	}
	sort.Sort(multi)
	return
}
