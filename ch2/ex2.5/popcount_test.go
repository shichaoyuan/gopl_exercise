package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountTrick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTrick(uint64(i))
	}
}
