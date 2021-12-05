package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 1 ----\n")
	part1()
	part2()
}

func part1() {
	vals := inputs()
	c := 0
	for i := 1; i < len(vals); i++ {
		if vals[i] > vals[i-1] {
			c++
		}
	}
	fmt.Printf("Part 1: %d\n\n", c)
}

func part2() {
	vals := inputs()
	c := 0
	for i := 0; i < len(vals)-3; i++ {
		a := vals[i] + vals[i+1] + vals[i+2]
		b := vals[i+1] + vals[i+2] + vals[i+3]
		if b > a {
			c++
		}
	}
	fmt.Printf("Part 2: %d\n\n", c)
}

func inputs() []int {
	inputs := strings.Split(input, "\n")
	vals := make([]int, len(inputs))
	for i := range inputs {
		vals[i], _ = strconv.Atoi(inputs[i])
	}
	return vals
}
