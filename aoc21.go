package main

import (
	"AdventOfCode2021/day1"
	"AdventOfCode2021/day2"
	"AdventOfCode2021/day3"
)

func main() {
	days := [...]func(){
		day1.Run,
		day2.Run,
		day3.Run,
	}

	days[len(days)-1]()
}
