package day17

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 17 ----\n")
	part1and2()
}

func part1and2() {
	target := inputs()
	var maxY int
	var maxV xy
	hits := make(map[xy]struct{})
	for x := 1; x <= target.xmax; x++ {
		for y := target.ymin; y <= -target.ymin; y++ {
			v := xy{x, y}
			p := probe{v: v}
			if hit, max := p.simulate(target); hit {
				hits[v] = struct{}{}
				if max > maxY {
					maxY, maxV = max, v
				}
			}
		}
	}
	fmt.Printf("Part 1 answer: %d (%d,%d)\n\nPart 2 answer: %d\n\n", maxY, maxV.x, maxV.y, len(hits))
}

func (p *probe) simulate(t target) (hit bool, maxY int) {
	for p.pos.y > t.ymin && p.pos.x < t.xmax {

		p.pos.x += p.v.x
		p.pos.y += p.v.y
		if p.pos.y > p.maxY {
			p.maxY = p.pos.y
		}
		if p.v.x < 0 {
			p.v.x++
		}
		if p.v.x > 0 {
			p.v.x--
		}
		p.v.y--
		if p.pos.x >= t.xmin && p.pos.x <= t.xmax && p.pos.y >= t.ymin && p.pos.y <= t.ymax {
			return true, p.maxY
		}
	}
	return false, p.maxY
}

var reRule = regexp.MustCompile(`target area: x=(?P<xmin>-?\d+)..(?P<xmax>-?\d+), y=(?P<ymin>-?\d+)..(?P<ymax>-?\d+)`)

type probe struct {
	pos  xy
	v    xy
	maxY int
}
type xy struct{ x, y int }
type target struct{ xmin, xmax, ymin, ymax int }

func inputs() (t target) {
	vals := reRule.FindStringSubmatch(input)
	t.xmin, _ = strconv.Atoi(vals[1])
	t.xmax, _ = strconv.Atoi(vals[2])
	t.ymin, _ = strconv.Atoi(vals[3])
	t.ymax, _ = strconv.Atoi(vals[4])
	return t
}
