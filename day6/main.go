package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(raw)
	groups := strings.Split(input, "\n\n")

	var (
		unique    int
		unanimous int
	)
	for _, group := range groups {
		unique += numUnique(group)
		unanimous += numUnanimous(group)
	}
	
	fmt.Println(unique)
	fmt.Println(unanimous)
}

func numUnique(group string) int {
	seen := make(map[string]struct{})

	for _, resp := range group {
		if string(resp) == "\n" {
			continue
		}
		seen[string(resp)] = struct{}{}
	}

	return len(seen)
}

func numUnanimous(group string) int {
	counts := make(map[string]int)
	numMembers := len(strings.Split(group, "\n"))

	for _, resp := range group {
		if string(resp) == "\n" {
			continue
		}
		counts[string(resp)] += 1
	}

	var n int
	for _, count := range counts {
		if count == numMembers {
			n++
		}
	}
	return n
}
