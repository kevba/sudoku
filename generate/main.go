package main

import (
	"log"
	"math/rand"
	"sudoku"
)

func main() {
	s := generateRandom()
	sudoku.Print(s)
}

// generateRandom generates a "random" sudoku using the solver
// Its not really all that random at all.
func generateRandom() *sudoku.Sudoku {
	s := sudoku.NewSudoku()

	randomCellValue := &sudoku.Cell{
		Value:      rand.Intn(8) + 1,
		FixedValue: true,
	}

	randomRegion := s.Regions[rand.Intn(9)]
	// Fill a cell with a randomly picked value, this will act as a "seed" for our sudoku
	randomRegion.Cells[rand.Intn(9)] = randomCellValue

	solver := sudoku.NewSolver(s)
	if solver.Solve() != nil {
		log.Printf("cannot be generated")
	}

	v := sudoku.Validator{}
	if !v.Validate(s) {
		log.Printf("generated sudoku is invalid")
	}

	return s
}
