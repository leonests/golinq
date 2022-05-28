package golinq

func (src Enumerator[K, V]) Distinct() Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()
			set := make(map[any]struct{})
			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext(); ok; k, v, ok = moveNext() {
					if _, exist := set[v]; !exist {
						set[v] = struct{}{}
						return
					}
				}
				return
			}
		},
	}
}

func (src Enumerator[K, V]) DistinctBy(selector func(K, V) any) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext := src.Enumerate()
			set := make(map[any]struct{})
			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext(); ok; k, v, ok = moveNext() {
					item := selector(k, v)
					if _, exist := set[item]; !exist {
						set[item] = struct{}{}
						return
					}
				}
				return
			}
		},
	}
}
