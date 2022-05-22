package golinq

func (src IEnumerator[K, V]) Skip(span int) IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
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
