package sudoku

import (
	"fmt"
)

// Create a bit mask for a list of numbers in given group. (Row, Column, Block)
//
// Example
// @maskPostions = [1, 2, 3, 5]
// @return maskNumber = 0000010111
func createBitMask(maskPositions []uint8) uint16 {
	mask := uint16(0)
	for _, n := range maskPositions {
		mask |= 1 << (n - 1)
	}
	return mask
}

// Merge all bitMasks into one. And revert it to build up candidates
//
// Example
// @value - the current value of the cell.
// @bitMasks - list of bitMasks of surrounding (from the Row, from the Column, from the Block relative to this cell).
// @return posible value of this cell.
func createCandidate(value uint8, bitMasks []uint16) uint16 {
	mask := uint16(0)
	// value already set
	if value > 0 {
		return mask
	}
	for _, n := range bitMasks {
		mask |= n
	}
	return 1023 & ^mask
}

func toChar(i uint8) string {
	if i >= 1 && i <= 9 {
		return fmt.Sprintf("%d", i)
	}
	return " "
}

// Write out in block of 3
// Print candidates as string in bulk of "3"
//
// @c - candidate
// @offset - the offset bit of the candidates (0, 3, 6)
// @return string of candidates
func getCandidates(c uint16, offset uint8) [3]string {
	out := [3]string{"", "", ""}
	bit := uint16(1) << offset
	for i := 0; i < 3; i++ {
		if bit&c > 0 {
			out[i] = toChar(uint8(uint8(i) + offset + 1))
		} else {
			out[i] = toChar(0)
		}
		bit = bit << 1
	}
	return out
}
