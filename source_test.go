package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		res := FromSlice(test.IntList).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := FromSlice(test.StrList).ToSlice()
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("Array", func(t *testing.T) {
		res := FromSlice(test.ArrayAny[:]).ToSlice()
		assert.Equal(t, wanted.AnyList, res)
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
		res := make(map[int]int)
		FromSlice(test.IntList).AsMap(&res)
		assert.Equal(t, wanted.MapIntInt, res)
	})

	t.Run("StringList", func(t *testing.T) {
		res := make(map[int]string)
		FromSlice(test.StrList).AsMap(&res)
		assert.Equal(t, wanted.MapIntStr, res)
	})

	t.Run("MapIntStr", func(t *testing.T) {
		res := make(map[int]string)
		FromMap(test.MapIntStr).AsMap(&res)
		assert.Equal(t, wanted.MapIntStr, res)
	})

	t.Run("MapStrInt", func(t *testing.T) {
		res := make(map[string]int)
		FromMap(test.MapStrInt).AsMap(&res)
		assert.Equal(t, wanted.MapStrInt, res)
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
		assert.Equal(t, wanted.FloatList, res)
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
		assert.Equal(t, wanted.StrList, res)
	})
	t.Run("JustFloat", func(t *testing.T) {
		res := Just(test.FloatList...).ToSlice()
		assert.Equal(t, wanted.FloatList, res)
	})
	t.Run("JustInt", func(t *testing.T) {
		res := Just(test.IntList...).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})
	t.Run("JustAny", func(t *testing.T) {
		res := Just(test.ArrayAny[:]...).ToSlice()
		assert.Equal(t, wanted.AnyList, res)
	})
	t.Run("JustAny", func(t *testing.T) {
		res := Just(test.ArrayAny[:]).ToSlice()
		assert.Equal(t, [][]any{{1, 2, 3}}, res)
	})

	t.Run("Range", func(t *testing.T) {
		res := Range(1, 6).ToSlice()
		assert.Equal(t, wanted.IntList, res)
	})
}

type col struct {
	item0 int
	item1 bool
	item2 string
	item3 float64
	item4 interface{}
}

func (c col) Enumerate() MoveNext[any, any] {
	index := 0
	return func() (k, v any, ok bool) {
		ok = true
		switch index {
		case 0:
			v = c.item0
		case 1:
			v = c.item1
		case 2:
			v = c.item2
		case 3:
			v = c.item3
		case 4:
			v = c.item4
		default:
			ok = false
		}
		k = index
		index++
		return
	}
}

func TestFrom(t *testing.T) {
	t.Parallel()
	t.Run("Slice", func(t *testing.T) {
		ch := make(chan any, 3)
		ch <- -1
		ch <- -2
		ch <- -3
		close(ch)

		tests := []struct {
			input  any
			wanted any
		}{
			{
				[]string{"a", "b", "c"},
				[]any{"a", "b", "c"},
			},
			{
				[]int{1, 2, 3},
				[]any{1, 2, 3},
			},
			{
				[]any{1, "2", []any{1, 2, 3}},
				[]any{1, "2", []any{1, 2, 3}},
			},
			{
				ch,
				[]any{-1, -2, -3},
			},
			{
				"abcde",
				[]any{'a', 'b', 'c', 'd', 'e'},
			},
			{
				col{item0: 1, item1: true, item2: "abc", item3: -3.1415926, item4: map[string]any{"3": []any{1, 2.2, 3.3}}},
				[]any{1, true, "abc", -3.1415926, map[string]any{"3": []any{1, 2.2, 3.3}}},
			},
		}
		for _, test := range tests {
			res := From(test.input).ToSlice()
			assert.Equal(t, test.wanted, res)
		}
	})

	t.Run("MapIntString", func(t *testing.T) {
		input := map[int]string{1: "1", 2: "2", 3: "3"}
		wanted := map[int]string{1: "1", 2: "2", 3: "3"}
		res := make(map[int]string)
		From(input).AsMap(&res)
		assert.Equal(t, wanted, res)
	})

	t.Run("MapStringAny", func(t *testing.T) {
		input := map[string]any{"3": []any{1, 2.2, 3.3}}
		wanted := map[string]any{"3": []any{1, 2.2, 3.3}}
		res := make(map[string]any)
		From(input).AsMap(&res)
		assert.Equal(t, wanted, res)
	})
}
