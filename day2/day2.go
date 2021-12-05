package day2

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 2 ----\n")
	part1()
	part2()
}

func part1() {
	ins := inputs()
	pos, depth := 0, 0
	for _, v := range ins {
		switch v.command {
		case "forward":
			pos += v.value
		case "up":
			depth -= v.value
		case "down":
			depth += v.value
		}
	}
	fmt.Printf("Part 1: %d\n\n", pos*depth)
}

func part2() {
	ins := inputs()
	pos, depth, aim := 0, 0, 0
	for _, v := range ins {
		switch v.command {
		case "forward":
			pos += v.value
			depth += v.value * aim
		case "up":
			aim -= v.value
		case "down":
			aim += v.value
		}
	}
	fmt.Printf("Part 2: %d\n\n", pos*depth)
}

type instruction struct {
	command string
	value   int
}

func inputs() (o []instruction) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		s := strings.Split(lines[i], " ")
		d, _ := strconv.Atoi(s[1])
		o = append(o, instruction{s[0], d})
	}
	return o
}
