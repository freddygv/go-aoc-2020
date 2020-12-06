package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const targetSum = 2020

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var input []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var num int

		_, err := fmt.Sscanf(s.Text(), "%d", &num)
		if err != nil {
			log.Fatalf("failed to scan: %v", err)
		}
		input = append(input, num)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(twoSumProduct(input, targetSum))
	fmt.Println(threeSumProduct(input, targetSum))
}

func twoSumProduct(input []int, sum int) int {
	memo := make(map[int]struct{})

	for _, value := range input {
		complement := sum - value
		if _, ok := memo[complement]; ok {
			return value * complement
		}

		memo[value] = struct{}{}
	}
	return -1
}

func threeSumProduct(input []int, sum int) int {
	for i := 0; i < len(input); i++ {
		complement := sum - input[i]
		if product := twoSumProduct(input[i+1:], complement); product != -1 {
			return input[i] * product
		}
	}
	return -1
}
