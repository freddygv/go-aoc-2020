package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	input = "input.txt"
)

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	preamble := 5
	if input == "input.txt" {
		preamble = 25
	}

	var (
		i      = 0
		bandit = 0
		buffer = ring.New(preamble)
		filled = false
		digits = make([]int, 0)
	)
	s := bufio.NewScanner(f)
	for s.Scan() {
		digit, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, digit)

		if filled {
			// Part 1
			ok := checkTwoSum(digit, buffer)
			if !ok {
				bandit = digit
				break
			}
		}

		buffer.Value = digit
		buffer = buffer.Next()

		i++
		if i >= preamble {
			filled = true
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("invalid number:", bandit)

	// Part 2
	bounds := findContiguousSumBounds(digits, bandit)
	fmt.Println(bounds[0] + bounds[len(bounds)-1])
}

func checkTwoSum(input int, r *ring.Ring) bool {
	memo := make(map[int]struct{}, 0)

	var valid bool
	r.Do(func(digit interface{}) {
		value := digit.(int)

		complement := input - value
		if _, ok := memo[complement]; ok {
			valid = true
		}
		memo[value] = struct{}{}
	})
	return valid
}

func findContiguousSumBounds(digits []int, target int) []int {
	if target == digits[0] {
		return []int{target, target}
	}

	var (
		i   = 0
		j   = 1
		sum = digits[i] + digits[j]
	)
	for i <= j && j < len(digits) {
		if sum < target {
			j++
			sum += digits[j]
		} else if sum > target {
			sum -= digits[i]
			i++
		} else {
			break
		}
	}

	var (
		min = math.MaxInt64
		max = math.MinInt64
	)
	for k := i; k <= j; k++ {
		if digits[k] < min {
			min = digits[k]
		}
		if digits[k] > max {
			max = digits[k]
		}
	}
	return []int{min, max}
}
