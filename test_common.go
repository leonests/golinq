package golinq

type GenericTest struct {
	Str       string
	Int       int
	IntList   []int
	StrList   []string
	AnyList   []any
	MapStrInt map[string]int
	MapIntStr map[int]string
	MapIntInt map[int]int
	RuneList  []rune
	ChanFloat chan float64
	FloatList []float64
	ArrayAny  [3]any
	BoolList  []bool
}
