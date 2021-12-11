package day12

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 12 ----\n")
	part1()
	part2()
}

func part1() {
	fmt.Printf("Part 1 answer: \n\n")
}

func part2() {
	fmt.Printf("Part 2 answer: \n\n")
}

func inputs() (o []string) {
	for _, line := range strings.Fields(input) {
		for _, v := range strings.Fields(line) {
			o = append(o, v)
		}
	}
	return o
}
