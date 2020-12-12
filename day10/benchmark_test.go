package main

import (
	"sort"
	"testing"
)

var (
	ratings      []int
	permutations int
)

func BenchmarkPart2(b *testing.B) {
	ratings = loadRatings()
	b.ResetTimer()

	b.Run("full input", benchPermutations)
}

func benchPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memo := make(map[int]int)

		sort.Ints(ratings)
		permutations = countPermutations(memo, ratings, 0, len(ratings)-1)
	}
}
