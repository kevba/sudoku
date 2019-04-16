package main

import (
	"flag"
	"io/ioutil"
	"log"
	"sudoku"
)

func main() {
	path := flag.String("i", "sudoku.json", "json file containing the sudoku")
	flag.Parse()

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	s, err := sudoku.FromJSON(data)
	if err != nil {
		log.Fatal(err)
	}

	solver := sudoku.NewSolver(s)
	err = solver.Solve()
	if err != nil {
		log.Fatal(err)
	}

	sudoku.Print(s)
}
