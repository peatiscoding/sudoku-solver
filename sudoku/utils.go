package sudoku

import (
	"fmt"
)

func toChar(i uint8) string {
	if i >= 1 && i <= 9 {
		return fmt.Sprintf("%d", i)
	}
	return " "
}

// Write out in block of 3
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
