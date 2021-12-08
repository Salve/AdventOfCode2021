package day5

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 5 ----\n")
	part1()
	part2()
}

func part1() {
	var d diagram
	for _, line := range inputs() {
		if !line.straight() {
			continue
		}
		d.chart(line)
	}
	fmt.Printf("Part 1: %d\n\n", d.count())
}

func part2() {
	var d diagram
	for _, line := range inputs() {
		d.chart(line)
	}
	fmt.Printf("Part 2: %d\n\n", d.count())
}

type point struct {
	x, y int
}

type line struct {
	a, b point
}

func (l line) straight() bool {
	return l.a.x == l.b.x || l.a.y == l.b.y
}

type diagram [1000][1000]int

func (d *diagram) chart(l line) {
	d[l.a.x][l.a.y]++
	for l.a != l.b {
		switch {
		case l.a.x < l.b.x:
			l.a.x++
		case l.a.x > l.b.x:
			l.a.x--
		}

		switch {
		case l.a.y < l.b.y:
			l.a.y++
		case l.a.y > l.b.y:
			l.a.y--
		}

		d[l.a.x][l.a.y]++
	}
}

func (d diagram) count() (c int) {
	for x := range d {
		for y := range d[x] {
			if d[x][y] > 1 {
				c++
			}
		}
	}
	return c
}

var re = regexp.MustCompile(`,| -> `)

func inputs() (o []line) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		s := re.Split(lines[i], -1)
		x1, _ := strconv.Atoi(s[0])
		y1, _ := strconv.Atoi(s[1])
		x2, _ := strconv.Atoi(s[2])
		y2, _ := strconv.Atoi(s[3])
		o = append(o, line{
			point{x1, y1},
			point{x2, y2},
		})
	}
	return o
}
