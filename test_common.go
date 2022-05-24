package golinq

type GenericTest struct {
	Str       string
	Int       int
	IntList   []int
	StrList   []string
	AnyList   []interface{}
	MapStrInt map[string]int
	MapIntStr map[int]string
	RuneList  []rune
	ChanFloat chan float64
	FloatList []float64
}

func RebuildSlice[K comparable, V any](e IEnumerator[K, V]) []V {
	res := make([]V, 0)
	moveNext := e.Enumerate()
	for _, v, ok := moveNext(); ok; _, v, ok = moveNext() {
		res = append(res, v)
	}
	return res
}

func RebuildMap[K comparable, V any](e IEnumerator[K, V]) map[K]V {
	res := make(map[K]V)
	moveNext := e.Enumerate()
	for k, v, ok := moveNext(); ok; k, v, ok = moveNext() {
		res[k] = v
	}
	return res
}
