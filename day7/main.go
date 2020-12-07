package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const target = "shiny gold"

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bags := make(map[string]*bag)

	s := bufio.NewScanner(f)
	for s.Scan() {
		processRule(s.Text(), bags)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	fmt.Println(countParents(bags, target))

	// Part 2
	fmt.Println(countChildren(bags, target))
}

func countParents(bags map[string]*bag, color string) int {
	var (
		q    = queue{contents: []string{color}}
		seen = make(map[string]struct{})
	)

	var parentCount int
	for !q.empty() {
		current := bags[q.dequeue()]

		for _, parent := range current.containedBy {
			if _, ok := seen[parent.color]; !ok {
				q.enqueue(parent.color)

				seen[parent.color] = struct{}{}
				parentCount++
			}
		}
	}
	return parentCount
}

func countChildren(bags map[string]*bag, color string) int {
	bag := bags[color]
	if bag == nil {
		return 0
	}
	if len(bag.contains) == 0 {
		return 0
	}

	var sum int
	for _, child := range bag.contains {
		children := countChildren(bags, child.color)
		sum += child.count + (child.count * children)
	}
	return sum
}

func processRule(line string, bags map[string]*bag) {
	parentRE := regexp.MustCompile(`^([a-z]+\s[a-z]+)`)
	color := parentRE.FindString(line)

	parent := bags[color]
	if parent == nil {
		parent = &bag{
			color: color,
		}
	}

	childrenRE := regexp.MustCompile(`([0-9])(?:\s)?([a-z]+\s[a-z]+)`)
	childrenMatches := childrenRE.FindAllSubmatch([]byte(line), -1)
	if len(childrenMatches) == 0 {
		bags[parent.color] = parent
		return
	}

	for _, contained := range childrenMatches {
		color := string(contained[2])

		child := bags[color]
		if child == nil {
			child = &bag{color: color}
		}
		if child.containedBy == nil {
			child.containedBy = make(map[string]*bag)
		}
		child.containedBy[parent.color] = parent

		count, err := strconv.Atoi(string(contained[1]))
		if err != nil {
			log.Fatal(err)
		}
		bc := &bagCount{
			color: child.color,
			count: count,
		}
		if parent.contains == nil {
			parent.contains = make(map[string]*bagCount)
		}
		parent.contains[child.color] = bc

		bags[child.color] = child
	}
	bags[parent.color] = parent

}

type bag struct {
	color       string
	contains    map[string]*bagCount
	containedBy map[string]*bag
}

type bagCount struct {
	color string
	count int
}

type queue struct {
	contents []string
}

func (q *queue) empty() bool {
	return len(q.contents) == 0
}

func (q *queue) enqueue(value string) {
	q.contents = append(q.contents, value)
}

func (q *queue) dequeue() string {
	if len(q.contents) == 0 {
		return ""
	}

	resp := q.contents[0]
	q.contents = q.contents[1:]

	return resp
}
