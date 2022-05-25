package golinq

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestGenericDestination(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "c"},
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		RuneList:  []int32{'g', 'e', 'n', 'e', 'r', 'i', 'c'},
		IntList:   []int{3, 6},
		StrList:   []string{"c"},
		MapIntStr: map[int]string{-1: "a"},
		MapIntInt: map[int]int{5: 6},
		FloatList: []float64{2.0},
		AnyList:   []interface{}{"1", "2", "3", "4", "5", "6"},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).ToSlice()
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).Where(func(i1, i2 int) bool { return i1 > 4 }).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntInt) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).Where(func(i int, s string) bool { return s < "e" }).ToSlice()
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := FromMap(test.MapIntStr).Where(func(i int, s string) bool { return i < 0 }).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntStr) {
			t.Fail()
		}
	})

	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).Where(func(idx int, v float64) bool { return math.Abs(v/2-1.0) <= 1e-5 }).ToSlice()
		if !reflect.DeepEqual(res, wanted.FloatList) {
			t.Fail()
		}
	})

	t.Run("IntList1", func(t *testing.T) {
		res := make([]interface{}, 0)
		for c := range FromSlice(test.IntList).Select(func(i int, v int) any { return fmt.Sprintf("%d", v) }).ToChannel() {
			res = append(res, c)
		}
		if !reflect.DeepEqual(res, wanted.AnyList) {
			t.Fail()
		}
	})
}

func TestToResult(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{1, 2, 3, 4, 5, 6},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []int32{'g', 'e', 'n', 'e', 'r', 'i', 'c'},
		IntList:  []int{3, 6},
		StrList:  []string{"1", "2", "3", "4", "5", "6"},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Select(func(i int, r rune) any { return r }).Result(&res)
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		var res []string
		FromSlice(test.IntList).Select(func(i1, i2 int) any { return fmt.Sprintf("%d", i2) }).Result(&res)
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		var res []map[string]int
		wanted := []map[string]int{{"g": 0}, {"e": 1}, {"n": 2}, {"e": 3}, {"r": 4}, {"i": 5}, {"c": 6}}
		FromSlice(test.StrList).Select(func(i int, s string) any { return map[string]int{s: i} }).Result(&res)
		if !reflect.DeepEqual(res, wanted) {
			t.Fail()
		}
	})
}

func TestForeach(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:     "generic",
		IntList: []int{1, 3, 5},
		StrList: []string{"g", "e", "n", "e", "r", "i", "c"},
	}
	wanted := GenericTest{
		RuneList: []int32{'g' + 1, 'e' + 1, 'n' + 1, 'e' + 1, 'r' + 1, 'i' + 1, 'c' + 1},
		IntList:  []int{3, 6},
		StrList:  []string{"2", "4", "6"},
	}
	t.Run("String", func(t *testing.T) {
		var res []rune
		FromString(test.Str).Foreach(func(i int, v rune) { res = append(res, v+1) })
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		var res []string
		FromSlice(test.IntList).Foreach(func(i, v int) { res = append(res, fmt.Sprintf("%d", v+1)) })
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		var res []map[string]int
		wanted := []map[string]int{{"g": 0}, {"e": 1}, {"n": 2}, {"e": 3}, {"r": 4}, {"i": 5}, {"c": 6}}
		FromSlice(test.StrList).Foreach(func(i int, s string) { res = append(res, map[string]int{s: i}) })
		if !reflect.DeepEqual(res, wanted) {
			t.Fail()
		}
	})
}
