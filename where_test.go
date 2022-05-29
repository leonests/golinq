package golinq

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericWhere(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		RuneList:  []int32{'g', 'n', 'r', 'i'},
		IntList:   []int{3, 6},
		StrList:   []string{"c"},
		MapIntStr: map[int]string{3: "c"},
		FloatList: []float64{2.0},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Where(func(i int, v rune) bool { return v > 'e' }).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Where(func(i int, v int) bool { return v%3 == 0 }).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Where(func(i int, s string) bool { return s < "e" }).ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]string)
		FromMap(test.MapIntStr).Where(func(i int, s string) bool { return i == 3 }).AsMap(&res)
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
}
