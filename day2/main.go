package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var count int

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := parseLine(s.Text())
		if line.isValidNew() {
			count++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("valid count:", count)
}

type Password struct {
	password string
	rule     Rule
}
type Rule struct {
	lower, upper int
	char         string
}

func (p *Password) isValidInitial() bool {
	var count int
	for _, char := range p.password {
		if string(char) == p.rule.char {
			count++
		}
	}

	if count < p.rule.lower || count > p.rule.upper {
		return false
	}
	return true
}

func (p *Password) isValidNew() bool {
	if p.rule.lower > len(p.password) || p.rule.upper > len(p.password) {
		return false
	}

	var n int
	if string(p.password[p.rule.lower-1]) == p.rule.char {
		n++
	}
	if string(p.password[p.rule.upper-1]) == p.rule.char {
		n++
	}
	return n == 1
}

func (p *Password) String() string {
	return fmt.Sprintf("pw: %s (lower: %d, upper: %d, char: %s)", p.password, p.rule.lower, p.rule.upper, p.rule.char)
}

func parseLine(line string) Password {
	split := strings.Split(line, ":")

	rule := strings.Split(split[0], " ")
	password := strings.TrimSpace(split[1])

	ruleBoundaries := strings.Split(rule[0], "-")
	lower, err := strconv.Atoi(ruleBoundaries[0])
	if err != nil {
		log.Fatal(err)
	}
	upper, err := strconv.Atoi(ruleBoundaries[1])
	if err != nil {
		log.Fatal(err)
	}

	return Password{
		password: password,
		rule: Rule{
			lower: lower,
			upper: upper,
			char:  rule[1],
		},
	}
}
