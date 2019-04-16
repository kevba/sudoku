package sudoku

import (
	"fmt"
)

type Solver struct {
	sudoku    *Sudoku
	validator *Validator
}

func NewSolver(s *Sudoku) Solver {
	return Solver{
		s,
		&Validator{true},
	}
}

func (s *Solver) Solve() error {
	if !s.solveCell(0) {
		return fmt.Errorf("Cannot be solved")
	}

	return nil
}

func (s *Solver) solveCell(cellNumber int) (solved bool) {
	if cellNumber >= (SIZE * SIZE) {
		return true
	}

	cell := s.sudoku.getCellByNumber(cellNumber)
	nextCell := cellNumber + 1

	for i := 1; i <= SIZE; i++ {
		if !cell.FixedValue {
			cell.Value = i
		}

		if s.validator.Validate(s.sudoku) {
			if s.solveCell(nextCell) {
				return true
			}
		}
	}

	if !cell.FixedValue {
		cell.Value = 0
	}

	return false
}
