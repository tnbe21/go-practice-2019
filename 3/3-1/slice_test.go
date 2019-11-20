package test

import "testing"

func BenchmarkInitIntSlice_DefinitionOnly(b *testing.B) {
	SIZE := b.N
	var base []int
	b.ResetTimer()
	for i := 0; i < SIZE; i++ {
		base = append(base, i)
	}
}

func BenchmarkInitIntSlice_DefinitionAndAllocate(b *testing.B) {
	SIZE := b.N
	base := make([]int, 0, SIZE)
	b.ResetTimer()
	for i := 0; i < SIZE; i++ {
		base = append(base, i)
	}
}
