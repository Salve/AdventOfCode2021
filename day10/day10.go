package day10

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 10 ----\n")
	part1()
	part2()
}

func part1() {
	total := 0
	for _, line := range inputs() {
		if score, c := parse(line); c {
			total += score
		}
	}
	fmt.Printf("Part 1 answer: %d\n\n", total)
}

func part2() {
	var scores []int
	for _, line := range inputs() {
		if score, c := parse(line); !c {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Printf("Part 2 answer: %d\n\n", scores[(len(scores)-1)/2])
}

func parse(line string) (score int, corrupt bool) {
	var stack []rune
	for _, r := range line {
		if _, opener := pairs[r]; opener {
			stack = append(stack, r)
			continue
		}
		n := len(stack) - 1
		opener := stack[n]
		stack = stack[:n]
		if pairs[opener] != r {
			return scores1[r], true
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		score = score*5 + scores2[pairs[stack[i]]]
	}
	return score, false
}

var pairs = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
var scores1 = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var scores2 = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

func inputs() (o []string) {
	return strings.Fields(input)
}
