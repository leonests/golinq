package golinq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericForeach(t *testing.T) {
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
		assert.Equal(t, wanted.RuneList, res)
	})

	t.Run("IntList", func(t *testing.T) {
		var res []string
		FromSlice(test.IntList).Foreach(func(i, v int) { res = append(res, fmt.Sprintf("%d", v+1)) })
		assert.Equal(t, wanted.StrList, res)
	})

	t.Run("StringList", func(t *testing.T) {
		var res []map[string]int
		wanted := []map[string]int{{"g": 0}, {"e": 1}, {"n": 2}, {"e": 3}, {"r": 4}, {"i": 5}, {"c": 6}}
		FromSlice(test.StrList).Foreach(func(i int, s string) { res = append(res, map[string]int{s: i}) })
		assert.Equal(t, wanted, res)
	})
}
