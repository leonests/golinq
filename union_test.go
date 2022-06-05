package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
	}
	wanted := GenericTest{
		RuneList:  []int32{'g', 'e', 'n', 'r', 'i', 'c'},
		IntList:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Union(FromString("generic")).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Union(Just(1, 2, 3, 7, 8, 9)).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})
}
