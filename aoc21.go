package main

import (
	"AdventOfCode2021/day1"
	"AdventOfCode2021/day2"
	"AdventOfCode2021/day3"
	"AdventOfCode2021/day4"
	"AdventOfCode2021/day5"
	"AdventOfCode2021/day6"
	"AdventOfCode2021/day7"
)

func main() {
	days := [...]func(){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
	}

	days[len(days)-1]()
}
