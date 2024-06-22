package solvers_wfc

import "github.com/peatiscoding/sudoku-solver/sudoku"

// Solve by set a new answer
func Solve(b *sudoku.Board) {
	b.Set(0, 1)
}
