package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var (
		min = math.MaxInt64
		max = math.MinInt64

		filled = map[int]struct{}{}
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		input := s.Text()
		row := binarySearch(input[:7], 127, "F", "B")
		seat := binarySearch(input[len(input)-3:], 7, "L", "R")

		seatID := row * 8 + seat
		if seatID > max {
			max = seatID
		}
		if seatID < min {
			min = seatID
		}

		filled[seatID] = struct{}{}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var seat int
	for i := min; i <= max; i++ {
		if _, ok := filled[i]; !ok {
			seat = i
			break
		}
	}

	fmt.Println("max:", max)
	fmt.Println("seat:", seat)
}

func binarySearch(input string, maxIdx int, lowerPartition, upperPartition string) int {
	var (
		min = 0
		max = maxIdx
	)
	for _, partition := range input {
		mid := min + (max-min)/2

		switch string(partition) {
		case upperPartition:
			min = mid+1
		case lowerPartition:
			max = mid
		default:
			return -1
		}
	}
	if min != max {
		return -1
	}
	return max
}


