package golinq

import (
	"reflect"
	"testing"
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
		IntList:  []int{3, 4, 5},
		StrList:  []string{"e", "r", "i", "c"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).Skip(3).ToSlice()
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Skip(3).ToSlice()
		if !reflect.DeepEqual(res, wanted.IntList) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Skip(3).ToSlice()
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})
}