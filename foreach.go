package golinq

func (src IEnumerator[K, V]) Foreach(action func(K, V)) {
	moveNext := src.Enumerate()

	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		action(k, v)
	}
}
