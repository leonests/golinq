package golinq

func (src IEnumerator[K, V]) Take(count int) IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
			moveNext := src.Enumerate()
			n := count 
			return func()(k K, v V, ok bool){
				if n <= 0 {
					return 
				}
				n--
				return moveNext()
			}
		},
	}
}