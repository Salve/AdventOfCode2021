package day3

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 3 ----\n")
	part1()
	part2()
}

func part1() {
	vals := inputs()
	var gamma, epsilon uint16
	for i := uint16(0); i < 12; i++ {
		most, least := count(vals, i)
		gamma |= most << (11 - i)
		epsilon |= least << (11 - i)
	}

	fmt.Printf("Part 1: gamma: %d, epsilon: %d, power consumption: %d\n\n", gamma, epsilon, int(gamma)*int(epsilon))
}

func part2() {

	o2filter := func(most, _ uint16) uint16 {
		return most
	}
	co2filter := func(_, least uint16) uint16 {
		return least
	}

	o2rating := process(o2filter, inputs())
	co2rating := process(co2filter, inputs())

	fmt.Printf("Part 1: o2 rating: %d, co2 rating: %d, life-support rating: %d\n\n", o2rating, co2rating, int(o2rating)*int(co2rating))
}

type bitCriteria func(most, least uint16) uint16

func process(f bitCriteria, remaining []uint16) uint16 {
	for i := uint16(0); len(remaining) > 1; i++ {
		remaining = discard(f, remaining, i)
	}
	return remaining[0]
}

func discard(f bitCriteria, vals []uint16, idx uint16) (remaining []uint16) {
	common := f(count(vals, idx))
	for _, v := range vals {
		if v>>(11-idx)&1 == common {
			remaining = append(remaining, v)
		}
	}
	return remaining
}

func count(vals []uint16, idx uint16) (most, least uint16) {
	var zeroes int
	for _, v := range vals {
		if v&(1<<(11-idx)) == 0 {
			zeroes++
		}
	}
	if zeroes*2 > len(vals) {
		return 0, 1
	}
	return 1, 0
}

func inputs() (o []uint16) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		v, _ := strconv.ParseUint(lines[i], 2, 12)
		o = append(o, uint16(v))
	}
	return o
}
