package day13

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
	fmt.Printf("\n---- Day 13 ----\n")
	part1()
	part2()
}

func part1() {
	s, ins := inputs()
	fmt.Printf("Part 1 answer: %d\n\n", len(s.fold(ins[0])))
}

func part2() {
	s, ins := inputs()
	for _, f := range ins {
		s = s.fold(f)
	}
	fmt.Printf("Part 2 answer: \n\n%s\n", s)
}

func (s sheet) fold(f int) sheet {
	foldby := "x"
	if f < 0 {
		foldby = "y"
		f = -f
	}

	folded := sheet{}
	switch foldby {
	case "x":
		for p := range s {
			if p.x <= f {
				folded[p] = struct{}{}
				continue
			}
			folded[point{f - (p.x - f), p.y}] = struct{}{}
		}
	case "y":
		for p := range s {
			if p.y <= f {
				folded[p] = struct{}{}
				continue
			}
			folded[point{p.x, f - (p.y - f)}] = struct{}{}
		}
	}
	return folded
}

func (s sheet) String() string {
	sb := strings.Builder{}
	maxX, maxY := s.max()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, ok := s[point{x, y}]; ok {
				sb.WriteString("##")
			} else {
				sb.WriteString("..")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s sheet) max() (x, y int) {
	for p := range s {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}
	return x, y
}

type point struct{ x, y int }
type sheet map[point]struct{}

var rePoint = regexp.MustCompile(`(?P<x>\d+),(?P<y>\d+)`)
var reFold = regexp.MustCompile(`fold along (?P<dir>[xy])=(?P<val>\d+)`)

func inputs() (sheet, []int) {
	o := make(sheet)
	var ins []int
	for _, s := range rePoint.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(s[1])
		y, _ := strconv.Atoi(s[2])
		o[point{x, y}] = struct{}{}
	}
	for _, s := range reFold.FindAllStringSubmatch(input, -1) {
		v, _ := strconv.Atoi(s[2])
		if s[1] == "y" {
			v = -v
		}
		ins = append(ins, v)
	}

	return o, ins
}
