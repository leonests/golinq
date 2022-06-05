package golinq

func (src Enumerator[K, V]) Union(obj Enumerator[K, V]) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext1 := src.Enumerate()
			moveNext2 := obj.Enumerate()
			srcGoing := true

			set := make(map[any]bool)
			return func() (k K, v V, ok bool) {
				if srcGoing {
					for k, v, ok = moveNext1(); ok; k, v, ok = moveNext1() {
						if _, exist := set[v]; !exist {
							set[v] = true
							return
						}
					}
					srcGoing = false
				}
				for k, v, ok = moveNext2(); ok; k, v, ok = moveNext2() {
					if _, exist := set[v]; !exist {
						set[v] = true
						return
					}
				}
				return
			}
		},
	}
}
