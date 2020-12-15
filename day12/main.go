package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	north   string = "N"
	south   string = "S"
	east    string = "E"
	west    string = "W"
	forward string = "F"
	left    string = "L"
	right   string = "R"
)

type instruction struct {
	action string
	value  int
}

const input = "input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	instructions := make([]instruction, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		input := s.Text()

		valueRaw := input[1:]
		val, err := strconv.Atoi(valueRaw)
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, instruction{
			action: string(input[0]),
			value:  val,
		})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	fmt.Println(navigatePart1(instructions))
	fmt.Println(navigatePart2(instructions))
}

func navigatePart1(instructions []instruction) int {
	heading := east
	position := []int{0, 0}

	for _, inst := range instructions {
		currHeading := heading

		// First get the direction we're going to move in for the current instruction
		switch {
		case inst.action == left || inst.action == right:
			if inst.value == 90 {
				heading = turn90(heading, inst.action)
			} else if inst.value == 180 {
				heading = turn180(heading)
			} else if inst.value == 270 {
				heading = turn270(heading, inst.action)
			}
			currHeading = heading

			// For L/R there's no movement only rotation
			continue

		// N/S/E/W headings will not affect the overall heading, only the current move's
		case inst.action == north || inst.action == south || inst.action == east || inst.action == west:
			currHeading = inst.action
		}

		switch currHeading {
		case north:
			position[1] += inst.value
		case south:
			position[1] -= inst.value
		case east:
			position[0] += inst.value
		case west:
			position[0] -= inst.value
		}
	}

	return distance(position[0], position[1])
}

func navigatePart2(instructions []instruction) int {
	var (
		shipLoc     = []int{0, 0}
		waypointLoc = []int{10, 1}
	)

	for _, inst := range instructions {
		if inst.action == left || inst.action == right {
			waypointLoc = rotateAroundOrigin(waypointLoc, inst.action, inst.value)
			continue
		}

		switch inst.action {
		case north:
			waypointLoc[1] += inst.value
		case south:
			waypointLoc[1] -= inst.value
		case east:
			waypointLoc[0] += inst.value
		case west:
			waypointLoc[0] -= inst.value
		case forward:
			shipLoc[0] += inst.value * waypointLoc[0]
			shipLoc[1] += inst.value * waypointLoc[1]
		}
	}
	return distance(shipLoc[0], shipLoc[1])
}

func rotateAroundOrigin(loc []int, direction string, degrees int) []int {
	center := []int{0, 0}
	rotatedVector := make([]int, 2)

	for degrees > 0 {
		rotatingVector := []int{loc[0] - center[0], loc[1] - center[1]}

		switch direction {
		case "L":
			rotatedVector = []int{-rotatingVector[1], rotatingVector[0]}
		case "R":
			rotatedVector = []int{rotatingVector[1], -rotatingVector[0]}
		}
		loc = []int{rotatedVector[0] + center[0], rotatedVector[1] + center[1]}

		degrees -= 90
	}
	return loc
}

func turn270(input string, direction string) string {
	firstTurn := turn180(input)
	return turn90(firstTurn, direction)
}

func turn90(input string, direction string) string {
	if direction == "L" {
		switch input {
		case north:
			return west
		case east:
			return north
		case south:
			return east
		case west:
			return south
		}
	}
	if direction == "R" {
		switch input {
		case north:
			return east
		case east:
			return south
		case south:
			return west
		case west:
			return north
		}
	}
	return ""
}

func turn180(input string) string {
	switch input {
	case north:
		return south
	case east:
		return west
	case south:
		return north
	case west:
		return east
	}
	return ""
}

// manhattan distance
func distance(x, y int) int {
	return abs(x) + abs(y)
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
