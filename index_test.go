package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		FloatList: []float64{1.1, 2.2, 3.3},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
	}
	wanted := GenericTest{
		Int:       1,
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).IndexOf(func(i int, r rune) bool {
			return r == 'e'
		})
		assert.Equal(t, wanted.Int, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).IndexOf(func(i1, i2 int) bool {
			return i2 % 2 == 0
		})
		assert.Equal(t, wanted.Int, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).IndexOf(func(i int, s string) bool {
			return s == "e"
		})
		assert.Equal(t, wanted.Int, res)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := FromMap(test.MapIntStr).OrderBy(func(i int, s string) any {
			return i
		}).IndexOf(func(i int, s string) bool {
			return i % 2 == 0 && s != "c"
		})
		assert.Equal(t, wanted.Int, res)
	})
}
