package golinq

func (src Enumerator[K, V]) Reverse() Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()

			items := make([]sorter[K, V], 0)
			for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
				items = append(items, sorter[K, V]{k, v})
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
