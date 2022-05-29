package golinq

func (src Enumerator[K, V]) GroupBy(keySelector func(K, V) any) Enumerator[any, any] {
	return Enumerator[any, any]{
		Enumerate: func() MoveNext[any, any] {
			moveNext := src.Enumerate()
			set := make(map[any][]any)
			for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
				key := keySelector(k, v)
				set[key] = append(set[key], v)
			}

			index, length := 0, len(set)
			keys := make([]any, length)
			for k := range set {
				keys[index] = k
				index++
			}
			index = 0
			return func() (k any, v any, ok bool) {
				ok = index < length
				if ok {
					k = keys[index]
					v = set[k]
					index++
				}
				return
			}
		},
	}
}

func (src Enumerator[K, V]) GroupByKV(keySelector, valueSelector func(K, V) any) Enumerator[any, any] {
	return Enumerator[any, any]{
		Enumerate: func() MoveNext[any, any] {
			moveNext := src.Enumerate()
			set := make(map[any][]any)
			for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
				key := keySelector(k, v)
				set[key] = append(set[key], valueSelector(k, v))
			}

			index, length := 0, len(set)
			keys := make([]any, length)
			for k := range set {
				keys[index] = k
				index++
			}
			index = 0
			return func() (k any, v any, ok bool) {
				ok = index < length
				if ok {
					k = keys[index]
					v = set[k]
					index++
				}
				return
			}
		},
	}
}
