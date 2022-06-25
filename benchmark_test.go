package golinq

import "testing"

func BenchmarkSum_Generic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(1, 1000000).Where(func(i1, i2 int) bool {
			return i2%2 == 0
		}).Sum2Int()
	}
}

func BenchmarkWhereSelectWhereFirst_Generic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(1, 1000000).SelectV(func(i1, i2 int) int {
			return -i2
		}).Where(func(i1, i2 int) bool {
			return i2 > -1000
		}).First()
	}
}

func BenchmarkSkipTake_Generic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(1, 1000000).Skip(2).Take(5)
	}
}
