package sudoku

import (
	"fmt"
)

type Board struct {
	// Updating vals would remove possible choices
	Vals       [81]uint8  // max is 9
	Candidates [81]uint16 // bitwise of 9 choices (2^10 - 1) = (1024 - 1) = 1023
}

func toChar(i uint8) string {
	if i >= 1 && i <= 9 {
		return fmt.Sprintf("%d", i)
	}
	return " "
}

// Create a disable mask for a list of numbers
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

func New(input string) *Board {
	b := &Board{
		Vals:       [81]uint8{},
		Candidates: [81]uint16{},
	}

	if len(input) > 81 {
		fmt.Printf("input too long: %d\n", len(input))
	}

	for i := 0; i < 81; i++ {
		if i < len(input) && input[i] != ' ' {
			b.Vals[i] = uint8(input[i] - '0')
			b.Candidates[i] = 1 << (b.Vals[i] - 1)
		} else {
			b.Vals[i] = 0
			b.Candidates[i] = 1023 // all choices are possible.
		}
	}
	b.CalculateChoices()
	return b
}

func (b *Board) CalculateChoices() {
	// Calculate choices
	rowMasks := [9]uint16{}
	colMasks := [9]uint16{}
	blkMasks := [9]uint16{} // 0, 1, 2, 3, 4, 5, 6, 7, 8
	// (a) Check same row
	for rw := 0; rw < 9; rw++ {
		picked := b.Vals[rw*9 : rw*9+9]
		rowMasks[rw] = createBitMask(picked)
	}
	// (b) Check same column
	for col := 0; col < 9; col++ {
		picked := make([]uint8, 9)
		for rw := 0; rw < 9; rw++ {
			picked[rw] = b.Vals[rw*9+col]
		}
		colMasks[col] = createBitMask(picked)
	}
	// (c) Check same block
	for blk := 0; blk < 9; blk++ {
		picked := make([]uint8, 9)
		for rw := 0; rw < 3; rw++ {
			for col := 0; col < 3; col++ {
				picked[rw*3+col] = b.Vals[blk*9+rw*3+col]
			}
		}
		blkMasks[blk] = createBitMask(picked)
	}

	// Update candidates back
	for i := 0; i < 81; i++ {
		b.Candidates[i] = 1023 & ^(rowMasks[i/9] | colMasks[i%9] | blkMasks[i/27*3+i%9/3])
	}
}

func (b *Board) Print() {
	fmt.Println("Board:")
	fmt.Println("┌─────┬─────┬─────┐")
	for rw := 0; rw < 9; rw++ {
		for rep := 0; rep < 3; rep++ {
			offset := rw*9 + rep*3
			fmt.Printf("│%s %s %s", toChar(b.Vals[offset+0]), toChar(b.Vals[offset+1]),
				toChar(b.Vals[offset+2]))
		}
		fmt.Println("│")
		if (rw+1)%3 == 0 && rw < 8 {
			fmt.Println("├─────┼─────┼─────┤")
		}
	}
	fmt.Println("└─────┴─────┴─────┘")
}

func (b *Board) PrintCandidates() {
	// Print candidates per cell
	fmt.Println("Candidates:")
	fmt.Println("┌─────┬─────┬─────┐")
	for rw := 0; rw < 9; rw++ {
		for rep := 0; rep < 3; rep++ {
			offset := rw*9 + rep*3
			fmt.Printf("│%s %s %s", toChar(b.Vals[offset+0]), toChar(b.Vals[offset+1]),
				toChar(b.Vals[offset+2]))
		}
		fmt.Println("│")
		if (rw+1)%3 == 0 && rw < 8 {
			fmt.Println("├─────┼─────┼─────┤")
		}
	}
	fmt.Println("└─────┴─────┴─────┘")
}
