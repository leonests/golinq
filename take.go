package golinq

func (src Enumerator[K, V]) Take(count int) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()
			n := count
			return func() (k K, v V, ok bool) {
				if n <= 0 {
					return
				}
				n--
				return moveNext()
			}
		},
	}
}
