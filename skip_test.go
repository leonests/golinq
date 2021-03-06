package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericSkip(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{0, 1, 2, 3, 4, 5},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []rune{'e', 'r', 'i', 'c'},
		IntList:  []int{},
		StrList:  []string{"e", "r", "i", "c"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Skip(3).ToSlice()
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Skip(6).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Skip(3).ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})
}
