package golinq

import (
	"reflect"
	"testing"
)

func TestFromSliceAndString(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:      "generic",
		IntList:  []int{1, 2, 3, 4, 5, 6},
		StrList:  []string{"g", "e", "n", "e", "r", "i", "c"},
		ArrayAny: [3]any{1, 2, 3},
	}
	wanted := GenericTest{
		RuneList: []int32{'g', 'e', 'n', 'e', 'r', 'i', 'c'},
		IntList:  []int{1, 2, 3, 4, 5, 6},
		StrList:  []string{"g", "e", "n", "e", "r", "i", "c"},
		AnyList:  []any{1, 2, 3},
	}
	t.Run("String", func(t *testing.T) {
		res := FromString(test.Str).ToSlice()
		if !reflect.DeepEqual(res, wanted.RuneList) {
			t.Fail()
		}
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).ToSlice()
		if !reflect.DeepEqual(res, wanted.IntList) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).ToSlice()
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})

	t.Run("Array", func(t *testing.T) {
		res := FromSlice(test.ArrayAny[:]).ToSlice()
		if !reflect.DeepEqual(res, wanted.AnyList) {
			t.Fail()
		}
	})
}

func TestFromMap(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"a", "b", "c"},
		MapIntStr: map[int]string{0: "a", 1: "b", 2: "c"},
		MapStrInt: map[string]int{"a": 1, "b": 2, "c": 3},
	}
	wanted := GenericTest{
		MapIntStr: map[int]string{0: "a", 1: "b", 2: "c"},
		MapStrInt: map[string]int{"a": 1, "b": 2, "c": 3},
		MapIntInt: map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6},
	}

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntInt) {
			t.Fail()
		}
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntStr) {
			t.Fail()
		}
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := FromMap(test.MapIntStr).ToMap()
		if !reflect.DeepEqual(res, wanted.MapIntStr) {
			t.Fail()
		}
	})

	t.Run("MapStrInt", func(t *testing.T) {
		res := FromMap(test.MapStrInt).ToMap()
		if !reflect.DeepEqual(res, wanted.MapStrInt) {
			t.Fail()
		}
	})
}

func TestFromChan(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		ChanFloat: make(chan float64, 3),
	}
	wanted := GenericTest{
		FloatList: []float64{1.0, 2.0, 3.0},
	}
	t.Run("ChanFloat", func(t *testing.T) {
		test.ChanFloat <- 1.0
		test.ChanFloat <- 2.0
		test.ChanFloat <- 3.0
		close(test.ChanFloat)
		res := FromChan(test.ChanFloat).ToSlice()
		if !reflect.DeepEqual(res, wanted.FloatList) {
			t.Fail()
		}
	})
}

func TestFromJustAndRange(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		FloatList: []float64{1.0, 2.0, 3.0},
		ArrayAny:  [3]any{1, 2, 3},
	}
	wanted := GenericTest{
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"generic"},
		FloatList: []float64{1.0, 2.0, 3.0},
		AnyList:   []any{1, 2, 3},
	}
	t.Run("JustString", func(t *testing.T) {
		res := Just(test.Str).ToSlice()
		if !reflect.DeepEqual(res, wanted.StrList) {
			t.Fail()
		}
	})
	t.Run("JustFloat", func(t *testing.T) {
		res := Just(test.FloatList...).ToSlice()
		if !reflect.DeepEqual(res, wanted.FloatList) {
			t.Fail()
		}
	})
	t.Run("JustInt", func(t *testing.T) {
		res := Just(test.IntList...).ToSlice()
		if !reflect.DeepEqual(res, wanted.IntList) {
			t.Fail()
		}
	})
	t.Run("JustAny", func(t *testing.T) {
		res := Just(test.ArrayAny[:]...).ToSlice()
		if !reflect.DeepEqual(res, wanted.AnyList) {
			t.Fail()
		}
	})
	t.Run("JustAny", func(t *testing.T) {
		res := Just(test.ArrayAny[:]).ToSlice()
		if !reflect.DeepEqual(res, [][]any{{1, 2, 3}}) {
			t.Fail()
		}
	})

	t.Run("Range", func(t *testing.T) {
		res := Range(1, 6).ToSlice()
		if !reflect.DeepEqual(res, wanted.IntList) {
			t.Fail()
		}
	})
}
