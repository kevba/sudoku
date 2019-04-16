package sudoku

import "encoding/json"

func FromJSON(data []byte) (*Sudoku, error) {
	jsonData := sudokuJSON{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	regions := [9]*Region{}

	for i, region := range jsonData.Regions {
		cells := [9]*Cell{}
		for j, cellValue := range region.Cells {
			if cellValue != 0 {
				cells[j] = &Cell{cellValue, true}
			}
			cells[j] = &Cell{cellValue, false}
		}

		regions[i] = &Region{cells}
	}

	return &Sudoku{regions}, nil
}

type sudokuJSON struct {
	Regions []regionsJSON `json:"regions"`
}

type regionsJSON struct {
	Cells []int `json:"cells"`
}
