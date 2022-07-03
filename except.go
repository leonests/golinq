package golinq

func (src Enumerator[K, V]) Except(other Enumerator[K, V]) Enumerator[K, V]{
	return Enumerator[K, V] {
		Enumerate: func() MoveNext[K, V] {
			moveNext1 := src.Enumerate()
			moveNext2 := other.Enumerate()
			set := make(map[any]bool)
			for _, v, ok := moveNext2(); ok; _, v, ok = moveNext2() {
				set[v] = true
			} 

			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext1(); ok; k, v, ok = moveNext1() {
					if _, has := set[v]; !has {
						return
					}
				} 
				return
			}
		},
	}
}
