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
	fmt.Printf("Part 1 answer: %d\n\n", traverse(1, nil, inputs(), "start"))
}

func part2() {
	fmt.Printf("Part 2 answer: %d\n\n", traverse(2, nil, inputs(), "start"))
}

func traverse(max int, visited []string, connections map[string][]string, current string) (count int) {
	if current == "end" {
		return 1
	}
	if isLower(current) && in(visited, current) {
		if max == 1 || current == "start" {
			return 0
		}
		if max == 2 {
			max = 1
		}
	}
	visited = append(visited, current)
	for _, c := range connections[current] {
		count += traverse(max, visited, connections, c)
	}
	return count
}

func in(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

func inputs() map[string][]string {
	connections := make(map[string][]string)
	for _, line := range strings.Fields(input) {
		cave := strings.Split(line, "-")
		connections[cave[0]] = append(connections[cave[0]], cave[1])
		connections[cave[1]] = append(connections[cave[1]], cave[0])
	}
	return connections
}
