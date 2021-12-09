package day8

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 8 ----\n")
	part1()
	part2()
}

func part1() {
	c := 0
	for _, v := range inputs() {
		for _, digit := range v.output {
			if l := bitsSet(digit); l == 2 || l == 4 || l == 3 || l == 7 {
				c++
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n\n", c)
}

func part2() {
	sum := 0
	for _, entry := range inputs() {
		val, _ := strconv.Atoi(decode(entry))
		sum += val
	}
	fmt.Printf("Part 2 answer: %d\n\n", sum)
}

func decode(entry entry) (o string) {
	m := findMappings(entry)
	for _, digit := range entry.output {
		o += fmt.Sprintf("%d", mapDigit(digit, m))
	}
	return o
}

func mapDigit(d sevenseg, m mappings) int {
	var mapped sevenseg
	for i := a; i <= g; i <<= 1 {
		if d&i == i {
			mapped |= m[i]
		}
	}
	for i, digit := range digits {
		if mapped == digit {
			return i
		}
	}
	fmt.Printf("failed on: %8b\n\n", mapped)
	panic("rip")
}

func findMappings(entry entry) mappings {
	// Count how many times each scrambled segment occurs, across all ten input digits
	counts := make(map[sevenseg]int)
	for _, v := range entry.input {
		for i := 0; i < 7; i++ {
			seg := sevenseg(1 << i)
			if v&seg == seg {
				counts[seg]++
			}
		}
	}

	// For each scrambled segment, set the real segment, based on the known number of occurrences in the real input.
	m := make(mappings)
	for k, v := range counts {
		switch v {
		case 9:
			m[k] = f
		case 8:
			m[k] = a | c
		case 7:
			m[k] = d | g
		case 6:
			m[k] = b
		case 4:
			m[k] = e
		}
	}

	// Two pairs of segments are still uncertain. The digits 1, 4, 7 and 8 can be identified based on the number of
	// segments alone. From these, the number 1 contains segment C but not A and the number 4 contains segment D but not G.
	// Use this to lock down the remaining uncertain mappings.
	for _, digit := range entry.input {
		switch bitsSet(digit) {
		case 2:
			for seg, mapping := range m {
				if seg&digit == seg && mapping == a|c {
					m[seg] = c
				} else if mapping == a|c {
					m[seg] = a
				}
			}
		case 4:
			for seg, mapping := range m {
				if seg&digit == seg && mapping == d|g {
					m[seg] = d
				} else if mapping == d|g {
					m[seg] = g
				}
			}
		}
	}

	return m
}

type mappings map[sevenseg]sevenseg
type sevenseg uint8

const (
	a sevenseg = 0b00000001
	b sevenseg = 0b00000010
	c sevenseg = 0b00000100
	d sevenseg = 0b00001000
	e sevenseg = 0b00010000
	f sevenseg = 0b00100000
	g sevenseg = 0b01000000
)

var digits = [10]sevenseg{
	a | b | c | e | f | g,     // 0
	c | f,                     // 1
	a | c | d | e | g,         // 2
	a | c | d | f | g,         // 3
	b | c | d | f,             // 4
	a | b | d | f | g,         // 5
	a | b | d | e | f | g,     // 6
	a | c | f,                 // 7
	a | b | c | d | e | f | g, // 8
	a | b | c | d | f | g,     // 9
}

type entry struct {
	input  [10]sevenseg
	output [4]sevenseg
}

func bitsSet(s sevenseg) (o int) {
	for s != 0 {
		s &= s - 1
		o++
	}
	return o
}

func inputs() (o []entry) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		inout := strings.Split(lines[i], " | ")
		inputs := strings.Fields(inout[0])
		outputs := strings.Fields(inout[1])
		entry := entry{}
		for i := 0; i < 10; i++ {
			entry.input[i] = stringSegment(inputs[i])
		}
		for i := 0; i < 4; i++ {
			entry.output[i] = stringSegment(outputs[i])
		}
		o = append(o, entry)
	}
	return o
}

func stringSegment(s string) (o sevenseg) {
	for _, r := range s {
		o |= runeSegment(r)
	}
	return o
}

func runeSegment(r rune) sevenseg {
	switch r {
	case 'a':
		return a
	case 'b':
		return b
	case 'c':
		return c
	case 'd':
		return d
	case 'e':
		return e
	case 'f':
		return f
	case 'g':
		return g
	}
	panic("rip")
}
