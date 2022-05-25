package golinq

import (
	"reflect"
	"testing"
)

func TestGenericSelect(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		RuneList:  []rune{'g', 'n', 'r', 'i'},
		MapIntInt: map[int]int{0: -1, 1: 2, 2: 3},
		AnyList:   []any{2.0, 4.0, 6.0},
		BoolList:  []bool{false, false, true, false, false, true},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Where(func(i int, v rune) bool { return v > 'e' }).Select(func(i int, v rune) any { return v }).Result(&res)
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		var res []bool
		FromSlice(test.IntList).Select(func(i int, v int) any { return v%3 == 0 }).Result(&res)
		if !reflect.DeepEqual(res, wanted.BoolList) {
			t.Fail()
		}
	})

	t.Run("MapIntStr", func(t *testing.T) {
		var res map[int]int
		FromMap(test.MapIntStr).Select(func(k int, v string) any { return k }).Result(&res)
		if !reflect.DeepEqual(res, wanted.MapIntInt) {
			t.Fail()
		}
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).Select(func(k int, v float64) any { return 2 * v }).ToSlice()
		if !reflect.DeepEqual(res, wanted.AnyList) {
			t.Fail()
		}
	})
}



