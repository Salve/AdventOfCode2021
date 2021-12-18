package main

import (
	"AdventOfCode2021/day1"
	"AdventOfCode2021/day10"
	"AdventOfCode2021/day11"
	"AdventOfCode2021/day12"
	"AdventOfCode2021/day13"
	"AdventOfCode2021/day14"
	"AdventOfCode2021/day2"
	"AdventOfCode2021/day3"
	"AdventOfCode2021/day4"
	"AdventOfCode2021/day5"
	"AdventOfCode2021/day6"
	"AdventOfCode2021/day7"
	"AdventOfCode2021/day8"
	"AdventOfCode2021/day9"

	"fmt"
	"time"
)

var days = [...]func(){
	day1.Run,
	day2.Run,
	day3.Run,
	day4.Run,
	day5.Run,
	day6.Run,
	day7.Run,
	day8.Run,
	day9.Run,
	day10.Run,
	day11.Run,
	day12.Run,
	day13.Run,
	day14.Run,
}

func main() {
	timeFunc(latest())
	// timeAll()
}

func latest() func() {
	return days[len(days)-1]
}

func timeAll() {
	var d time.Duration
	for _, f := range days {
		d += timeFunc(f)
	}
	fmt.Printf("\nTotal execution time: %s\n\n", d)
}

func timeFunc(f func()) time.Duration {
	start := time.Now()
	f()
	d := time.Now().Sub(start)
	fmt.Printf("Execution time: %s\n", d)
	return d
}
