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
		MapIntInt: map[int]int{-1: -1},
		AnyList:   []any{2.0, 4.0, 6.0},
		BoolList:  []bool{false, false, true, false, false, true},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Where(func(i int, v rune) bool { return v > 'e' }).Select(func(i int, v rune) any { return v }).AsSlice(&res)
		assert.Equal(t, res, wanted.RuneList)
	})

	t.Run("IntList", func(t *testing.T) {
		var res []bool
		FromSlice(test.IntList).Select(func(i int, v int) any { return v%3 == 0 }).AsSlice(&res)
		assert.Equal(t, res, wanted.BoolList)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]int)
		FromMap(test.MapIntStr).Where(func(i int, s string) bool { return i < 0 }).Select(func(k int, v string) any { return k }).AsMap(&res)
		assert.Equal(t, res, wanted.MapIntInt)
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).Select(func(k int, v float64) any { return 2 * v }).ToSlice()
		assert.Equal(t, res, wanted.AnyList)
	})
}

func TestGenericSelectMany(t *testing.T) {
	t.Parallel()
	t.Run("IntList", func(t *testing.T) {
		input := [][]int{{1, 2, 3}, {5, 6, 7}}
		wanted := []any{1, 2, 3, 5, 6, 7}
		res := FromSlice(input).SelectMany(func(i1 int, i2 []int) any { return i2 }).ToSlice()
		assert.Equal(t, res, wanted)
	})

	t.Run("StringList", func(t *testing.T) {
		input := []string{"select", "many"}
		wanted := []any{'s', 'e', 'l', 'e', 'c', 't', 'm', 'a', 'n', 'y'}
		res := FromSlice(input).SelectMany(func(i int, s string) any { return FromString(s).ToSlice() }).ToSlice()
		assert.Equal(t, res, wanted)
	})

	t.Run("Map2Slice", func(t *testing.T) {
		input := map[string][]int{"a": {1, 2, 3}, "b": {-1, -2, -3}}
		wanted := []any{-3, -2, -1, 1, 2, 3}
		res := FromMap(input).SelectMany(func(s string, i []int) any { return i }).
			OrderBy(func(i int, a any) any { return a }).ToSlice()
		assert.Equal(t, res, wanted)
	})

	t.Run("Map2Map", func(t *testing.T) {
		input := map[string][]int{"a": {1, 2, 3}, "b": {-1, -2, -3}}
		wanted := []any{-3, -2, -1, 1, 2, 3}
		res := FromMap(input).SelectMany(func(s string, i []int) any { return i }).
			OrderBy(func(i int, a any) any { return a }).ToSlice()
		assert.EqualValues(t, res, wanted)
	})
}
