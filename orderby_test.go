package golinq

import (
	"reflect"
	"testing"
)

func TestGenericOrderBy(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{6, 5, 4, 3, 2, 1},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{2: "b", 3: "c", -1: "a"},
		MapStrInt: map[string]int{"c": 3, "a": 2, "b": 1},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		RuneList:  []int32{'c', 'e', 'e', 'g', 'i', 'n', 'r'},
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"c", "e", "e", "g", "i", "n", "r"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		AnyList:   []any{"a", "b", "c"},
		FloatList: []float64{6.0, 5.0, 4.0, 3.0, 2.0, 1.0},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).OrderBy(func(i int, r rune) any { return r }).ToSlice()
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).OrderBy(func(i1, i2 int) any { return i2 }).ToSlice()
		if !reflect.DeepEqual(res, wanted.IntList) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).OrderBy(func(i int, s string) any { return s }).ToSlice()
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := FromMap(test.MapIntStr).OrderBy(func(i int, s string) any { return i }).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntStr) {
			t.Fail()
		}
	})
	t.Run("MapStrInt", func(t *testing.T) {
		res := FromMap(test.MapStrInt).OrderBy(func(s string, i int) any { return s }).
			Select(func(s string, i int) any { return s }).ToSlice()
		if !reflect.DeepEqual(res, wanted.AnyList) {
			t.Fail()
		}
	})

	t.Run("IntList1", func(t *testing.T) {
		res := make([]float64, 0)
		FromSlice(test.IntList).OrderBy(func(i1, i2 int) any { return -i2 }).
			Select(func(i1, i2 int) any { return float64(i2) }).Result(&res)
		if !reflect.DeepEqual(res, wanted.FloatList) {
			t.Fail()
		}
	})
}