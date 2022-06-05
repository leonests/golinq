package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{1, 2, 3, 4, 5, 6},
	}
	wanted := GenericTest{
		RuneList: []int32{'g', 'e', 'n'},
		IntList:  []int{1, 2, 3},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Intersect(FromString("gene")).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Intersect(Just(1, 2, 3, 7, 8, 9)).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})
}

func TestSupersect(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{1, 2, 3, 3, 3, 5},
	}
	wanted := GenericTest{
		RuneList: []int32{'g', 'e', 'n', 'e'},
		IntList:  []int{1, 2, 3, 3},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Supersect(FromString("gene")).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Supersect(Just(1, 2, 3, 7, 3)).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})
}
