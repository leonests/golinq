package golinq

// current golang version 1.18.2 does not support generic in
// anonymous function, nor generic type parameter in method
// so here just return any, it will break the type chain
func (src Enumerator[K, V]) Select(selector func(K, V) any) Enumerator[int, any] {
	return Enumerator[int, any]{
		Enumerate: func() MoveNext[int, any] {
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

func (src Enumerator[K, V]) SelectMany(selector func(K, V) any) Enumerator[int, any] {
	return Enumerator[int, any]{
		Enumerate: func() MoveNext[int, any] {
			outMoveNext := src.Enumerate()
			var (
				index      int
				inMoveNext MoveNext[any, any]
				outKey     K
				outVal     V
				inVal      any
				outGoing   bool
			)

			return func() (k int, v any, ok bool) {
				for !ok {
					if !outGoing {
						outKey, outVal, ok = outMoveNext()
						if !ok {
							return
						}
						outGoing = true
						item := selector(outKey, outVal)
						inMoveNext = From(item).Enumerate() // rely on non-generic version
					}

					_, inVal, ok = inMoveNext()
					if !ok {
						outGoing = false
					} else {
						k = index
						v = inVal
						index++
					}
				}
				return
			}
		},
	}
}
