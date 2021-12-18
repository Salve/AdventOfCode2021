package day14

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 14 ----\n")
	part1()
	part2()
}

func part1() {
	template, rules := inputs()
	fmt.Printf("Part 1 answer: %d\n\n", solve(template, rules, 10))
}

func part2() {
	template, rules := inputs()
	fmt.Printf("Part 2 answer: %d\n\n", solve(template, rules, 40))
}

func solve(template string, r rules, steps int) int {
	count := make(pcount)
	for i := 0; i < len(template)-1; i++ {
		count[pair{template[i], template[i+1]}]++
	}
	o := counts(template, react(count, r, steps))
	return o[len(o)-1] - o[0]
}

func react(c pcount, r rules, step int) pcount {
	if step == 0 {
		return c
	}
	nc := make(pcount)
	for p, v := range c {
		nc[pair{p[0], r[p]}] += v
		nc[pair{r[p], p[1]}] += v
	}
	return react(nc, r, step-1)
}

func counts(template string, c pcount) (o []int) {
	count := make(map[byte]int)
	for p, v := range c {
		count[p[0]] += v
	}
	count[template[len(template)-1]]++
	for _, v := range count {
		o = append(o, v)
	}
	sort.Ints(o)
	return o
}

type pair [2]byte
type pcount map[pair]int
type rules map[pair]byte

var reRule = regexp.MustCompile(`(?P<pair>\w+) -> (?P<insert>\w)`)

func inputs() (template string, r rules) {
	template = strings.SplitN(input, "\n", 2)[0]
	r = make(rules)
	for _, s := range reRule.FindAllStringSubmatch(input, -1) {
		r[pair{s[1][0], s[1][1]}] = s[2][0]
	}
	return
}
