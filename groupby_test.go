package golinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()
	test := GenericTest{
		Str:       "generic",
		IntList:   []int{1, 2, 3, 4, 5, 6},
		StrList:   []string{"g", "e", "n", "e", "r", "i", "c"},
		MapIntStr: map[int]string{-1: "a", 2: "b", 3: "a"},
	}
	t.Run("IntList1", func(t *testing.T) {
		res := make(map[int][]any)
		wanted := map[int][]any{0: {2, 4, 6}, 1: {1, 3, 5}}
		FromSlice(test.IntList).GroupBy(
			func(i1, i2 int) any { return i2 % 2 }).AsMap(&res)
		assert.Equal(t, wanted, res)
	})

	t.Run("IntList2", func(t *testing.T) {
		res := make(map[int][]any)
		wanted := map[int][]any{0: {2, 4, 6}, 1: {1, 3, 5}}
		FromSlice(test.IntList).GroupByKV(
			func(i1, i2 int) any { return i2 % 2 },
			func(i1, i2 int) any { return i2 }).AsMap(&res)
		assert.Equal(t, wanted, res)
	})

	t.Run("MapIntStr1", func(t *testing.T) {
		res := make(map[string][]any)
		wanted := map[string][]any{"a": {"a", "a"}, "b": {"b"}}
		FromMap(test.MapIntStr).GroupBy(func(i int, s string) any {
			return s
		}).AsMap(&res)
		assert.Equal(t, wanted, res)
	})

	t.Run("MapIntStr2", func(t *testing.T) {
		res := make(map[string][]any)
		wanted := map[string][]any{"b": {2}}
		FromMap(test.MapIntStr).GroupByKV(
			func(i int, s string) any { return s },
			func(i int, s string) any { return i },
		).Where(func(a1, a2 any) bool { return a1.(string) == "b" }).AsMap(&res)
		assert.Equal(t, wanted, res)
	})

}
