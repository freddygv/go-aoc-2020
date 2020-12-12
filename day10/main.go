package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	input = "input.txt"

	outletRating       = 0
	deviceDifferential = 3
)

func main() {
	ratings := loadRatings()
	sort.Ints(ratings)

	// Part 1
	var (
		current, prev int
		diffCounts    = make(map[int]int)
	)
	for i := 0; i < len(ratings); i++ {
		current = ratings[i]

		diff := current - prev
		if diff > 3 {
			fmt.Println(i, current, prev)
			panic("this shouldn't happen")
		}
		diffCounts[diff] += 1
		prev = current
	}
	fmt.Printf("%d * %d = %d\n", diffCounts[1], diffCounts[3], diffCounts[1]*diffCounts[3])

	// Part 2
	memo := make(map[int]int)
	fmt.Println(countPermutations(memo, ratings, 0, len(ratings)-1))
}

func countPermutations(memo map[int]int, input []int, start, end int) int {
	if start == end {
		return 1
	}
	if _, ok := memo[start]; ok {
		return memo[start]
	}

	var count int
	for i := start + 1; i <= end; i++ {
		diff := input[i] - input[start]
		if diff > 3 {
			break
		}
		count += countPermutations(memo, input, i, end)
	}
	memo[start] = count
	return count
}

func loadRatings() []int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var (
		ratings = []int{outletRating}
		max     = math.MinInt64
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		rating, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		if rating > max {
			max = rating
		}
		ratings = append(ratings, rating)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// Also append our device
	deviceRating := max + deviceDifferential
	ratings = append(ratings, deviceRating)

	return ratings
}
