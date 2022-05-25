package golinq

func (src IEnumerator[K, V]) Distinct() IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
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

func (src IEnumerator[K, V]) DistinctBy(selector func(K, V) any) IEnumerator[K, V] {
	return IEnumerator[K, V]{
		Enumerate: func() IMoveNext[K, V] {
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
