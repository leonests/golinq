package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericReverse(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{6, 5, 4, 3, 2, 1},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []rune{'c', 'i', 'r', 'e', 'n', 'e', 'g'},
		IntList:  []int{1, 2, 3, 4, 5, 6},
		StrList:  []string{"c", "i", "r", "e", "n", "e", "g"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Reverse().ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Reverse().ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Reverse().ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})
}
