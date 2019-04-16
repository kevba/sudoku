package sudoku

import (
	"fmt"
)

func Print(s *Sudoku) {
	for i := 0; i < SIZE; i++ {
		values := s.getRowValues(i)
		for _, v := range values {
			fmt.Printf("|%v", v)
		}
		fmt.Printf("|\n")
	}

}
