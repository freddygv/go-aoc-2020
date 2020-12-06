package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	matrix := make([][]int, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := parseLine(s.Text())
		matrix = append(matrix, line)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	steps := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	resp := 1
	for _, step := range steps {
		resp *= numTreesEncountered(matrix, step[0], step[1])
	}
	fmt.Println(resp)
}

func numTreesEncountered(mat [][]int, rightStep int, downStep int) int {
	var (
		n int
		x int
		y int
	)

	rowLen := len(mat[0])
	for y < len(mat) {
		if mat[y][x%rowLen] == 1 {
			n++
		}

		y += downStep
		x += rightStep
	}
	return n
}

func parseLine(input string) []int {
	resp := make([]int, len(input))

	for i, char := range input {
		switch string(char) {
		case ".":
			resp[i] = 0

		case "#":
			resp[i] = 1
		}
	}

	return resp
}
