package sudoku

import (
	"fmt"
)

type Board struct {
	// Updating vals would remove possible choices
	Vals       [81]uint8  // max is 9
	Candidates [81]uint16 // bitwise of 9 choices (2^10 - 1) = (1024 - 1) = 1023
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
			// b.Candidates[i] = 1 << (b.Vals[i] - 1)
			b.Candidates[i] = 0
		} else {
			b.Vals[i] = 0
			b.Candidates[i] = 1023 // all choices are possible.
		}
	}
	b.CalculateChoices()
	return b
}

// Update a cell with a new value
// @return false if value is not valid.
func (b *Board) Set(position uint8, value uint8) bool {
	candidate := b.Candidates[position] // bitMask of possible values
	valBit := uint16(1) << (value - 1)
	if valBit&candidate > 0 {
		return false
	}
	b.Vals[position] = value
	b.CalculateChoices()
	return true
}

func (b *Board) Validate() error {
	// Validate board if there is any conficts by Rules
	return nil
}

// Compute bitMasks for provided position
// @position position of interest to compute
// @return bitMasks for Row, Column, Block
func (b *Board) bitMasks(position uint8) [3]uint16 {
	out := [3]uint16{
		0,
		0,
		0,
	}
	row := position / 9 // first of the row
	out[0] = createBitMask(b.Vals[row : row+9])

	col := position % 9 // first of each col
	colPicked := make([]uint8, 9)
	for rw := uint8(0); rw < 9; rw++ {
		colPicked[rw] = b.Vals[rw*9+col]
	}
	out[1] = createBitMask(colPicked)

	blkTop := position / 27 * 27 // integer division by 27 = [0, 1, 2] then +27 per each block to the top
	blkLeft := (col / 3) * 3     // integer mod 3 +3 per each block on the left
	blkOffset := blkTop + blkLeft
	blkPicked := make([]uint8, 9)
	for rw := uint8(0); rw < 3; rw++ {
		for cl := uint8(0); cl < 3; cl++ {
			blkPicked[3*rw+cl] = b.Vals[blkOffset+rw*9+cl]
		}
	}
	out[2] = createBitMask(blkPicked)
	return out
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
		top := blk / 3 * 27
		left := (blk % 3) * 3
		offset := top + left
		picked := make([]uint8, 9)
		for rw := 0; rw < 3; rw++ {
			for col := 0; col < 3; col++ {
				picked[rw*3+col] = b.Vals[offset+rw*9+col]
			}
		}
		blkMasks[blk] = createBitMask(picked)
	}

	// Update candidates back
	for i := 0; i < 81; i++ {
		b.Candidates[i] = createCandidate(b.Vals[i], []uint16{rowMasks[i/9], colMasks[i%9], blkMasks[i/27*3+i%9/3]})
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
	fmt.Println("┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐")
	for rw := 0; rw < 9; rw++ {
		for tinyRow := 0; tinyRow < 3; tinyRow++ {
			for cl := 0; cl < 9; cl++ {
				offset := rw*9 + cl
				cand := getCandidates(b.Candidates[offset], uint8(tinyRow*3))
				fmt.Printf("│%s %s %s", cand[0], cand[1], cand[2])
			}
			fmt.Println("│")
		}

		if rw < 8 {
			fmt.Println("├─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤")
		}
	}
	fmt.Println("└─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘")
}
