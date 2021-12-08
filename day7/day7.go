package day7

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 7 ----\n")
	part1()
	part2()
}

func part1() {
	pos := inputs()
	min, max := minMax(pos)
	minFuel := 0
	for i := min; i <= max; i++ {
		fuel := cost(pos, i)
		if fuel < minFuel || i == min {
			minFuel = fuel
		}
	}
	fmt.Printf("Part 1 answer: %d\n\n", minFuel)
}

func part2() {
	pos := inputs()
	min, max := minMax(pos)
	minFuel := 0
	for i := min; i <= max; i++ {
		fuel := cost2(pos, i)
		if fuel < minFuel || i == min {
			minFuel = fuel
		}
	}
	fmt.Printf("Part 2 answer: %d\n\n", minFuel)
}

func cost(vals []int, target int) (total int) {
	for _, v := range vals {
		total += abs(target - v)
	}
	return total
}

func cost2(vals []int, target int) (total int) {
	for _, v := range vals {
		distance := abs(target - v)
		total += (distance * (distance + 1)) / 2
	}
	return total
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func minMax(vals []int) (min, max int) {
	min, max = vals[0], vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func inputs() (o []int) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		for _, s := range strings.Split(lines[i], ",") {
			v, _ := strconv.Atoi(s)
			o = append(o, v)
		}

	}
	return o
}
