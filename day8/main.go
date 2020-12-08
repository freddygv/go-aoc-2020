package main

import (
	"bufio"
	"errors"
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

	instructions := make([]instruction, 0)
	invertible := make([]int, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		instruction := process(s.Text())
		instructions = append(instructions, instruction)

		if instruction.op == jmp || instruction.op == nop {
			invertible = append(invertible, len(instructions)-1)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	accumulator, err := traverse(instructions, true, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accumulator)

	// Part 2
	for _, idx := range invertible {
		accumulator, err := traverse(instructions, false, idx)
		if err != nil {
			continue
		}
		fmt.Println(instructions[idx].String(), accumulator)
	}
}

func traverse(arr []instruction, avoidRepeat bool, invertIdx int) (int, error) {
	var (
		i           int
		maxIter     int
		accumulator int
		seen        = make(map[int]int)
	)
	for i >= 0 && i < len(arr) {
		// The program should terminate before this many iterations
		maxIter++
		if maxIter > 60000 {
			return 0, errors.New("infinite loop")
		}

		// Avoid re-visiting operations in part 1
		current := arr[i]
		if _, ok := seen[i]; ok && avoidRepeat {
			break
		}
		seen[i] += 1

		// Flip a jmp/nop operation when trying to break inf loop in part 2
		if invertIdx == i {
			switch current.op {
			case jmp:
				current.op = nop
			case nop:
				current.op = jmp
			}
		}

		switch current.op {
		case acc:
			accumulator += current.arg
			i++
		case jmp:
			i += current.arg
		case nop:
			i++
		}
	}
	return accumulator, nil
}

func process(input string) instruction {
	split := strings.Split(input, " ")

	code, ok := opCodes[split[0]]
	if !ok {
		log.Fatalf("invalid op code %q", split[0])
	}
	arg, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	return instruction{
		op:  code,
		arg: arg,
	}
}

type instruction struct {
	op  operation
	arg int
}

func (i instruction) String() string {
	return fmt.Sprintf("(op=%s, arg=%d)", i.op, i.arg)
}

type operation int8

const (
	acc operation = iota
	jmp
	nop
)

var opCodes = map[string]operation{
	"acc": acc,
	"jmp": jmp,
	"nop": nop,
}

func (o operation) String() string {
	return [...]string{"acc", "jmp", "nop"}[o]
}
