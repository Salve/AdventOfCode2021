package day6

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 6 ----\n")
	part1()
	part2()
}

func part1() {
	example := school{3, 4, 3, 1, 2}
	fmt.Printf("Part 1 example: %d\n\n", example.simulate(18))

	school := inputs()
	fmt.Printf("Part 1 answer: %d\n\n", school.simulate(80))

	alt := inputs()
	var d days
	for _, f := range alt {
		d[f]++
	}
	fmt.Printf("Part 1 alt answer: %d\n\n", d.simulate(80))
}

func part2() {
	school := inputs()
	var d days
	for _, f := range school {
		d[f]++
	}
	fmt.Printf("Part 2 answer: %d\n\n", d.simulate(256))
}

type days [9]uint64

func (d days) simulate(days int) (sum uint64) {
	for i := 0; i < days; i++ {
		c := d
		for j := range d {
			switch j {
			case 0:
				d[0] -= c[0]
				d[6] += c[0]
				d[8] += c[0]
			default:
				d[j] -= c[j]
				d[j-1] += c[j]
			}
		}
	}
	for _, c := range d {
		sum += c
	}
	return sum
}

func (s *school) simulate(days int) int {
	var spawn school
	for i := range *s {
		switch (*s)[i] {
		case 0:
			spawn = append(spawn, fish(8))
			(*s)[i] = 6
		default:
			(*s)[i]--
		}
	}
	*s = append(*s, spawn...)

	days--
	if days == 0 {
		return len(*s)
	}
	return s.simulate(days)
}

type school []fish
type fish uint8

func inputs() (s school) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		vals := strings.Split(lines[i], ",")
		for _, v := range vals {
			f, _ := strconv.Atoi(v)
			s = append(s, fish(f))
		}
	}
	return s
}
