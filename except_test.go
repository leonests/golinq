package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericExcept(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
	}
	wanted := GenericTest{
		RuneList:  []rune{'g', 'n', 'r', 'i'},
		IntList:   []int{1, 2, 3},
		MapIntStr: map[int]string{-1: "a"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Except(FromSlice([]rune{'e', 'c'})).ToSlice()
		assert.Equal(t, res, wanted.RuneList)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Except(FromSlice([]int{4, 5, 6})).ToSlice()
		assert.Equal(t, res, wanted.IntList)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]string)
		FromMap(test.MapIntStr).Except(FromSlice([]string{"b", "c"})).AsMap(&res)
		assert.Equal(t, res, wanted.MapIntStr)
	})
}
