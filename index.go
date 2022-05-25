package golinq

func (src IEnumerator[K, V]) IndexOf(predict func(K, V) bool) int {
	moveNext := src.Enumerate()
	index := 0
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		if predict(k, v) {
			return index
		}
		index++
	}
	return -1
}
