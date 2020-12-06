package main

import (
	"fmt"
	"testing"
)

func TestPasswordWithRule_isValidNew(t *testing.T) {
	tt := []struct {
		input  string
		expect bool
	}{
		{
			input:  "1-3 a: abcde",
			expect: true,
		},
		{
			input:  "1-3 b: cdefg",
			expect: false,
		},
		{
			input:  "2-9 c: ccccccccc",
			expect: false,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			parsed := parseLine(tc.input)

			valid := parsed.isValidNew()
			if valid != tc.expect {
				t.Errorf(`(%q) expected "%t", got "%t"`, tc.input, tc.expect, valid)
			}
		})
	}
}
