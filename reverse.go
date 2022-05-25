package golinq

func (src IEnumerator[K, V]) Reverse() IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
			moveNext := src.Enumerate()

			items := make(sorters[K, V], 0)
			for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
				items = append(items, sorter[K, V]{k, v, nil})
			}
			index := len(items)
			return func() (k K, v V, ok bool) {
				index--
				if index < 0 {
					return
				}
				k, v, ok = items[index].key, items[index].value, true
				return
			}
		},
	}
}
