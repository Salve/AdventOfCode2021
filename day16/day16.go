package day16

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bearmini/bitstream-go"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 16 ----\n")
	part1and2()
}

func part1and2() {
	bitReader := inputs()
	res := bitReader.readPacket()
	fmt.Printf("Part 1 answer: %d\n\nPart 2 answer: %d\n\n", bitReader.versionSum, res)
}

func (r *br) readPacket() (result int) {
	v, t := r.readVersionType()
	r.versionSum += v
	switch t {
	case 4:
		result = r.readLiteral()
	default:
		result = r.readOperator(t)
	}
	return result
}

const (
	flagLastGroup uint8 = 1 << 4
	groupMask           = flagLastGroup - 1
)

func (r *br) readLiteral() (result int) {
	var lastGroup bool
	for !lastGroup {
		group, _ := r.ReadNBitsAsUint8(5)
		result <<= 4
		result |= int(group & groupMask)
		lastGroup = group&flagLastGroup != flagLastGroup
	}
	return result
}

func (r *br) readOperator(typeID int) (result int) {
	var vals []int

	switch r.readLengthTypeID() {
	case 0:
		bitsToRead := r.bitsRead + r.readSubpacketsBitCount()
		for r.bitsRead < bitsToRead {
			vals = append(vals, r.readPacket())
		}
	case 1:
		n := r.readSubpacketsCount()
		for i := 0; i < n; i++ {
			vals = append(vals, r.readPacket())
		}
	}

	return operator(typeID)(vals)
}

func (r *br) readSubpacketsBitCount() int {
	l, _ := r.ReadNBitsAsUint16BE(15)
	return int(l)
}

func (r *br) readSubpacketsCount() int {
	l, _ := r.ReadNBitsAsUint16BE(11)
	return int(l)
}

func (r *br) readLengthTypeID() uint8 {
	b, _ := r.ReadNBitsAsUint8(1)
	return b
}

func (r *br) readVersionType() (version, typeID int) {
	v, _ := r.ReadNBitsAsUint8(3)
	t, _ := r.ReadNBitsAsUint8(3)
	return int(v), int(t)
}

func operator(typeID int) func([]int) int {
	switch typeID {
	case 0: // Sum
		return func(vals []int) int {
			res := 0
			for _, v := range vals {
				res += v
			}
			return res
		}
	case 1: // Product
		return func(vals []int) int {
			res := vals[0]
			for i := 1; i < len(vals); i++ {
				res *= vals[i]
			}
			return res
		}
	case 2: // Minimum
		return func(vals []int) int {
			res := vals[0]
			for i := 1; i < len(vals); i++ {
				if vals[i] < res {
					res = vals[i]
				}
			}
			return res
		}
	case 3: // Maximum
		return func(vals []int) int {
			res := vals[0]
			for i := 1; i < len(vals); i++ {
				if vals[i] > res {
					res = vals[i]
				}
			}
			return res
		}
	case 5: // Greater than
		return func(vals []int) int {
			if vals[0] > vals[1] {
				return 1
			}
			return 0
		}
	case 6: // Less than
		return func(vals []int) int {
			if vals[0] < vals[1] {
				return 1
			}
			return 0
		}
	case 7: // Equal
		return func(vals []int) int {
			if vals[0] == vals[1] {
				return 1
			}
			return 0
		}
	default:
		panic("rip")
	}
}

func (r *br) ReadNBitsAsUint8(nBits uint8) (uint8, error) {
	r.bitsRead += int(nBits)
	return r.Reader.ReadNBitsAsUint8(nBits)
}

func (r *br) ReadNBitsAsUint16BE(nBits uint8) (uint16, error) {
	r.bitsRead += int(nBits)
	return r.Reader.ReadNBitsAsUint16BE(nBits)
}

type br struct {
	bitsRead   int
	versionSum int
	*bitstream.Reader
}

func inputs() *br {
	r := bitstream.NewReader(hex.NewDecoder(strings.NewReader(strings.Fields(input)[0])), nil)
	return &br{Reader: r}
}
