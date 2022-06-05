package golinq

// Intersect returns the set intersection of two collections, without duplicates
func (src Enumerator[K, V]) Intersect(obj Enumerator[K, V]) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext1 := src.Enumerate()
			moveNext2 := obj.Enumerate()

			set := make(map[any]bool)
			for _, v, ok := moveNext2(); ok; _, v, ok = moveNext2() {
				if _, exist := set[v]; !exist {
					set[v] = true
				}
			}
			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext1(); ok; k, v, ok = moveNext1() {
					if _, exist := set[v]; exist {
						delete(set, v)
						return
					}
				}
				return
			}
		},
	}
}

// Supersect returns the set intersection of two collections, with duplicates
func (src Enumerator[K, V]) Supersect(obj Enumerator[K, V]) Enumerator[K, V] {
	return Enumerator[K, V]{
		Enumerate: func() MoveNext[K, V] {
			moveNext1 := src.Enumerate()
			moveNext2 := obj.Enumerate()

			set := make(map[any]int)
			for _, v, ok := moveNext2(); ok; _, v, ok = moveNext2() {
				if _, exist := set[v]; !exist {
					set[v] = 1
				} else {
					set[v]++
				}
			}
			return func() (k K, v V, ok bool) {
				for k, v, ok = moveNext1(); ok; k, v, ok = moveNext1() {
					if _, exist := set[v]; exist {
						if set[v] > 0 {
							set[v]--
							if set[v] == 0 {
								delete(set, v)
							}
							return
						}
					}
				}
				return
			}
		},
	}
}
