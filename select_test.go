package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericSelect(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		RuneList:  []rune{'g', 'n', 'r', 'i'},
		MapIntInt: map[int]int{0: -1},
		AnyList:   []any{2.0, 4.0, 6.0},
		BoolList:  []bool{false, false, true, false, false, true},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Where(func(i int, v rune) bool { return v > 'e' }).Select(func(i int, v rune) any { return v }).AsSlice(&res)
		assert.EqualValues(t, res, wanted.RuneList)
	})

	t.Run("IntList", func(t *testing.T) {
		var res []bool
		FromSlice(test.IntList).Select(func(i int, v int) any { return v%3 == 0 }).AsSlice(&res)
		assert.EqualValues(t, res, wanted.BoolList)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]int)
		FromMap(test.MapIntStr).Where(func(i int, s string) bool { return i < 0 }).Select(func(k int, v string) any { return k }).AsMap(&res)
		assert.EqualValues(t, res, wanted.MapIntInt)
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).Select(func(k int, v float64) any { return 2 * v }).ToSlice()
		assert.EqualValues(t, res, wanted.AnyList)
	})
}
