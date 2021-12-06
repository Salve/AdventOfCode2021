package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 4 ----\n")
	part1()
	part2()
}

func part1() {
	score := play1(inputs())

	fmt.Printf("Part 1: %d \n\n", score)
}

func part2() {
	score := play2(inputs())

	fmt.Printf("Part 2: %d \n\n", score)
}

func play1(draws []byte, boards []bingoBoard) (score int) {
	for _, draw := range draws {
		for i := range boards {
			if boards[i].check(draw) {
				return boards[i].sum() * int(draw)
			}
		}
	}
	panic("no win")
}

func play2(draws []byte, boards []bingoBoard) (score int) {
	for _, draw := range draws {
		remaining := make([]bingoBoard, 0, len(boards))
		for i := range boards {
			if !boards[i].check(draw) {
				remaining = append(remaining, boards[i])
			} else if len(boards) == 1 {
				return boards[0].sum() * int(draw)
			}
		}
		boards = remaining
	}
	panic("rip")
}

type bingoBoard [5][5]byte

const marked = 1 << 7

func (b *bingoBoard) check(draw byte) (bingo bool) {
	for row := range *b {
		for col, v := range b[row] {
			if v == draw {
				b[row][col] = v | marked
				return b.bingo()
			}
		}
	}
	return false
}

func (b bingoBoard) bingo() bool {
	for i := range b {
		if b[i][0]&b[i][1]&b[i][2]&b[i][3]&b[i][4]&marked == marked {
			return true
		}
	}
	for i := range b {
		if b[0][i]&b[1][i]&b[2][i]&b[3][i]&b[4][i]&marked == marked {
			return true
		}
	}
	return false
}

func (b bingoBoard) sum() (total int) {
	for _, row := range b {
		for _, v := range row {
			if v&marked != marked {
				total += int(v)
			}
		}
	}
	return total
}

func inputs() (draws []byte, boards []bingoBoard) {
	lines := strings.Split(input, "\n")
	for _, v := range strings.Split(lines[0], ",") {
		b, _ := strconv.Atoi(v)
		draws = append(draws, byte(b))
	}
	for i := 2; i < len(lines); i += 6 {
		board := bingoBoard{}
		for j := 0; j < 5; j++ {
			row := [5]byte{}
			for col, v := range strings.Fields(lines[i+j]) {
				b, _ := strconv.Atoi(v)
				row[col] = byte(b)
			}
			board[j] = row
		}
		boards = append(boards, board)
	}
	return
}
