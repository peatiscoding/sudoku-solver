package main

import (
	"fmt"
	"io"
	"os"

	"github.com/peatiscoding/sudoku-solver/sudoku"
	solvers_wfc "github.com/peatiscoding/sudoku-solver/sudoku/solvers"
)

func main() {
	fmt.Println("Sudoku Solver")
	// Check if the correct number of arguments is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file-path>")
		return
	}
	// Get the file path from the first argument
	filePath := os.Args[1]
	fmt.Println(fmt.Sprintf("Reading puzzle from file... %s", filePath))

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	board := sudoku.New(string(content))

	board.Print()
	board.PrintCandidates()

	// Solve
	solvers_wfc.Solve(board)
	board.Print()
}
