package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericTake(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{0, 1, 2, 3, 4, 5},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []rune{'e', 'r', 'i'},
		IntList:  []int{3},
		StrList:  []string{},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Skip(3).Take(3).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Skip(3).Take(1).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Skip(3).Take(0).ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})
}
