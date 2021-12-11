package day11

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 11 ----\n")
	part1()
	part2()
}

func part1() {
	cave := inputs()
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += cave.step()
	}
	fmt.Printf("Part 1 answer: %d\n\n", flashes)
}

func part2() {
	cave := inputs()
	i := 0
	for cave.step() != 100 {
		i++
	}
	fmt.Printf("Part 2 answer: %d\n\n", i+1)
}

type point struct{ x, y int }
type cave map[point]uint8

func (c cave) step() (flashes int) {
	c.inc()
	return c.count()
}

func (c cave) inc() {
	for p := range c {
		c[p]++
	}
}

func (c cave) count() (flashes int) {
	for _, p := range c.high() {
		flashes++
		c.flash(p)
	}
	if flashes == 0 {
		return 0
	}
	return flashes + c.count()
}

func (c cave) high() (o []point) {
	for p := range c {
		if c[p] > 9 {
			o = append(o, p)
		}
	}
	return o
}

func (c cave) flash(p point) {
	c[p] = 0
	for _, ap := range c.adjacent(p) {
		c[ap]++
	}
}

func (c cave) adjacent(p point) (a []point) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			ap := point{p.x + x, p.y + y}
			if _, ok := c[ap]; ok && ap != p && c[ap] > 0 {
				a = append(a, ap)
			}
		}
	}
	return a
}

func inputs() cave {
	o := make(cave, 100)
	for x, line := range strings.Fields(input) {
		for y, v := range line {
			o[point{x, y}] = uint8(v - '0')
		}
	}
	return o
}
