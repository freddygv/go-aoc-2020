package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	input = "input.txt"
)

type part int

const (
	ONE part = iota
	TWO
)

var emptyThreshold = map[part]int{
	ONE: 4,
	TWO: 5,
}

type seat struct {
	x, y int
}

func main() {
	matrix := readInput()
	fmt.Println(countOccupied(matrix, ONE), "occupied seats in part 1")

	matrix = readInput()
	fmt.Println(countOccupied(matrix, TWO), "occupied seats in part 2")
}

func countOccupied(matrix [][]string, part part) int {
	for {
		toFlip := make([]seat, 0)
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				var flipSeat bool
				switch matrix[i][j] {
				case ".":
					continue

				case "#":
					switch part {
					case ONE:
						if numNeighbors(matrix, i, j) >= emptyThreshold[part] {
							flipSeat = true
						}
					case TWO:
						if numVisible(matrix, i, j) >= emptyThreshold[part] {
							flipSeat = true
						}
					}

				case "L":
					switch part {
					case ONE:
						if numNeighbors(matrix, i, j) == 0 {
							flipSeat = true
						}
					case TWO:
						if numVisible(matrix, i, j) == 0 {
							flipSeat = true
						}
					}
				}
				if flipSeat {
					toFlip = append(toFlip, seat{y: i, x: j})
				}
			}
		}
		if len(toFlip) == 0 {
			break
		}
		for _, seat := range toFlip {
			matrix[seat.y][seat.x] = flip(matrix[seat.y][seat.x])
		}
	}
	var numOccupied int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "#" {
				numOccupied++
			}
		}
	}
	return numOccupied
}

func numVisible(matrix [][]string, y, x int) int {
	var (
		n          int
		directions = make([][]int, 0)
	)

	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			directions = append(directions, []int{dy, dx})
		}
	}
	for _, dir := range directions {
		if fullSeatInDirection(matrix, dir, y, x) {
			n++
		}
	}
	return n
}

func fullSeatInDirection(matrix [][]string, dir []int, y, x int) bool {
	for {
		y += dir[0]
		if y < 0 || y >= len(matrix) {
			return false
		}
		x += dir[1]
		if x < 0 || x >= len(matrix[0]) {
			return false
		}

		switch matrix[y][x] {
		case "#":
			return true
		case "L":
			return false
		default:
			// floor space, continue
		}
	}
}

func numNeighbors(matrix [][]string, y, x int) int {
	var n int
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dy == 0 && dx == 0 {
				continue
			}

			newY := y + dy
			if newY < 0 || newY >= len(matrix) {
				continue
			}
			newX := x + dx
			if newX < 0 || newX >= len(matrix[0]) {
				continue
			}

			if matrix[newY][newX] == "#" {
				n++
			}
		}
	}
	return n
}

func flip(input string) string {
	if input == "L" {
		return "#"
	}
	return "L"
}

func printMatrix(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseLine(input string) []string {
	resp := make([]string, len(input))

	for i, char := range input {
		resp[i] = string(char)
	}

	return resp
}

func readInput() [][]string {
	matrix := make([][]string, 0)

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := parseLine(s.Text())
		matrix = append(matrix, line)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return matrix
}
