package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericDistinct(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 3, 4, 5, 3, 3, 2, 3, 3, 4, 2, 2},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "a", 3: "a"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		Int:       1,
		RuneList:  []int32{'g', 'e', 'n', 'r', 'i', 'c'},
		IntList:   []int{1, 3, 4, 5, 2},
		StrList:   []string{"g", "e", "n", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a"},
		FloatList: []float64{1.0},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Distinct().ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Distinct().ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Distinct().ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]string)
		FromMap(test.MapIntStr).DistinctBy(func(i int, s string) any { return s }).AsMap(&res)
		assert.Equal(t, wanted.Int, len(res))
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).DistinctBy(func(i int, f float64) any { return f - float64(i) }).ToSlice()
		assert.Equal(t, wanted.FloatList, res)
	})
}
