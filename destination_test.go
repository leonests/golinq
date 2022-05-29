package golinq

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericDestination(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		FloatList: []float64{1.1, 2.2, 3.3},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		Int:       21,
		Int64:     21,
		Float64:   21,
		RuneList:  []int32{'g', 'e', 'n', 'e', 'r', 'i', 'c'},
		IntList:   []int{3, 6},
		StrList:   []string{"c"},
		MapIntStr: map[int]string{-1: "a"},
		MapIntInt: map[int]int{5: 6},
		FloatList: []float64{2.0},
		AnyList:   []any{"1", "2", "3", "4", "5", "6"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := make(map[int]int)
		FromSlice(test.IntList).Where(func(i1, i2 int) bool { return i1 > 4 }).AsMap(&res)
		assert.Equal(t, wanted.MapIntInt, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Where(func(i int, s string) bool { return s < "e" }).ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]string)
		FromMap(test.MapIntStr).Where(func(i int, s string) bool { return i < 0 }).AsMap(&res)
		assert.Equal(t, wanted.MapIntStr, res)
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).Where(func(idx int, v float64) bool { return math.Abs(v/2-1.0) <= 1e-5 }).ToSlice()
		assert.Equal(t, wanted.FloatList, res)
	})

	t.Run("IntList1", func(t *testing.T) {
		res := make([]any, 0)
		for c := range FromSlice(test.IntList).Select(func(i int, v int) any { return fmt.Sprintf("%d", v) }).ToChannel() {
			res = append(res, c)
		}
		assert.Equal(t, wanted.AnyList, res)
	})

	t.Run("IntList2", func(t *testing.T) {
		res := FromSlice(test.IntList).Select(func(i int, v int) any { return v }).Sum2Int()
		assert.Equal(t, wanted.Int64, res)
	})
	t.Run("IntList3", func(t *testing.T) {
		res := FromSlice(test.IntList).Select(func(i int, v int) any { return v }).Sum2Float()
		assert.Equal(t, wanted.Float64, res)
	})

	t.Run("IntList3", func(t *testing.T) {
		res := FromSlice(test.IntList).Select(func(i int, v int) any { return v }).Sum2Float()
		assert.Equal(t, wanted.Float64, res)
	})

	t.Run("FloatList", func(t *testing.T) {
		res := FromSlice(test.FloatList).Select(func(i int, v float64) any { return v }).Sum2Float()
		assert.Equal(t, 6.6, res)
	})

	t.Run("IntList4", func(t *testing.T) {
		res := FromSlice(test.IntList).Select(func(i int, v int) any { return v }).Max()
		assert.Equal(t, 6, res)
	})

	t.Run("IntList5", func(t *testing.T) {
		res := FromSlice(test.IntList).Select(func(i int, v int) any { return v }).Min()
		assert.Equal(t, 1, res)
	})
}

func TestResult(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{1, 2, 3, 4, 5, 6},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []int32{'g', 'e', 'n', 'e', 'r', 'i', 'c'},
		IntList:  []int{3, 6},
		StrList:  []string{"1", "2", "3", "4", "5", "6"},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Select(func(i int, r rune) any { return r }).AsSlice(&res)
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		var res []string
		FromSlice(test.IntList).Select(func(i1, i2 int) any { return fmt.Sprintf("%d", i2) }).AsSlice(&res)
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		var res []map[string]int
		wanted := []map[string]int{{"g": 0}, {"e": 1}, {"n": 2}, {"e": 3}, {"r": 4}, {"i": 5}, {"c": 6}}
		FromSlice(test.StrList).Select(func(i int, s string) any { return map[string]int{s: i} }).AsSlice(&res)
		assert.Equal(t, wanted, res)
	})
}
