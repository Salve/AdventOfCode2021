package day9

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 9 ----\n")
	part1()
	part2()
}

func part1() {
	risk := 0
	cave := inputs()
	for y := range cave {
		for x, v := range cave[y] {
			if cave.isLow(point{x, y}) {
				risk += int(v) + 1
			}
		}
	}
	fmt.Printf("Part 1 answer: %d\n\n", risk)
}

func part2() {
	var basins []int
	cave := inputs()
	for y := range cave {
		for x := range cave[y] {
			basins = append(basins, cave.spelunk(point{x, y}))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Printf("Part 2 answer: %d\n\n", basins[0]*basins[1]*basins[2])
}

func (c *cave) spelunk(start point) (size int) {
	todo := []point{start}
	for len(todo) > 0 {
		p := todo[0]
		v := &c[p.y][p.x]
		todo = todo[1:]

		if *v&visited == visited || *v == 9 {
			continue
		}
		*v = *v | visited
		size++
		todo = append(todo, p.adjacent()...)
	}
	return size
}

func (p point) adjacent() (o []point) {
	for _, a := range []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}} {
		if a.x < 0 || a.y < 0 || a.x > 99 || a.y > 99 {
			continue
		}
		o = append(o, a)
	}
	return o
}

func (c cave) isLow(p point) bool {
	for _, pa := range p.adjacent() {
		if c[pa.y][pa.x] <= c[p.y][p.x] {
			return false
		}
	}
	return true
}

const visited uint8 = 1 << 7

type cave [100][100]uint8
type point struct{ x, y int }

func inputs() (o cave) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		for j, r := range lines[i] {
			o[i][j] = uint8(r - '0')
		}
	}
	return o
}
