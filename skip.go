package golinq

func (src Enumerator[K, V]) Skip(span int) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()
			n := span
			return func() (k K, v V, ok bool) {
				for ; n > 0; n-- {
					_, _, ok = moveNext()
					if !ok {
						return
					}
				}
				return moveNext()
			}
		},
	}
}
