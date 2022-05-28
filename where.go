package golinq

func (src Enumerator[K, V]) Where(predict func(K, V) bool) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()
			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext(); ok; k, v, ok = moveNext() {
					if predict(k, v) {
						return
					}
				}
				return
			}
		},
	}
}
