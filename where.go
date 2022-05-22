package golinq

func (src IEnumerator[K, V]) Where(predict func(K, V) bool) IEnumerator[K, V] {
	return IEnumerator[K, V] {
		Enumerate: func() IMoveNext[K,V] {
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