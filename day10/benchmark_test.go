package main

import (
	"testing"
)

var (
	ratings      []int
	permutations int
)

func BenchmarkPart2(b *testing.B) {
	ratings = loadRatings()
	b.Run("full input", benchPermutations)
}

func benchPermutations(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		memo := make(map[int]int)
		permutations = countPermutations(memo, ratings, 0, len(ratings)-1)
	}
}
