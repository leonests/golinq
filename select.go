package golinq

// current golang version 1.18.2 does not support generic in
// anonymous function, nor generic type parameter in method
// so here just return any, it will break the type chain
func (src IEnumerator[K, V]) Select(selector func(K, V) any) IEnumerator[int, any] {
	return IEnumerator[int, any]{
		Enumerate: func() IMoveNext[int, any] {
			index := 0
			moveNext := src.Enumerate()
			return func() (k int, v any, ok bool) {
				key, value, ok := moveNext()
				if ok {
					k = index
					v = selector(key, value)
					index++
				}
				return
			}
		},
	}
}

