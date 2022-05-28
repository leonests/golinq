package golinq

func (src Enumerator[K, V]) Contains(predict func(K, V) bool) bool {
	moveNext := src.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		if predict(k, v) {
			return true
		}
	}
	return false
}
