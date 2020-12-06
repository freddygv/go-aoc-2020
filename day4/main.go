package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var (
		count int
		line  string
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		current := s.Text()
		if current != "" {
			line += current + " "
			continue
		}

		id := parseID(strings.TrimSpace(line))
		if isValidID(id) {
			count++
		}
		line = ""
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func parseID(input string) map[string]string {
	resp := make(map[string]string)

	fields := strings.Split(input, " ")

	for _, f := range fields {
		split := strings.Split(f, ":")
		resp[split[0]] = split[1]
	}

	return resp
}

func isValidID(id map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required {
		value, ok := id[field]
		if !ok {
			return false
		}
		if !isValidValue(field, value) {
			return false
		}
	}
	return true
}

func isValidValue(field, value string) bool {
	switch field {
	case "byr":
		return isWithinRange(value, 1920, 2002)

	case "iyr":
		return isWithinRange(value, 2010, 2020)

	case "eyr":
		return isWithinRange(value, 2020, 2030)

	case "hcl":
		hclRE := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return hclRE.MatchString(value)

	case "ecl":
		eclRE := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		return eclRE.MatchString(value)

	case "pid":
		pidRE := regexp.MustCompile(`^[0-9]{9}$`)
		return pidRE.MatchString(value)

	case "hgt":
		qty := value[:len(value)-2]

		if strings.HasSuffix(value, "cm") {
			return isWithinRange(qty, 150, 193)
		}
		if strings.HasSuffix(value, "in") {
			return isWithinRange(qty, 59, 76)
		}
		return false
	}
	return true
}

func isWithinRange(input string, min, max int) bool {
	value, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	if value < min || value > max {
		return false
	}
	return true
}
