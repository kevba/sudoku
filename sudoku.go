package sudoku

const SIZE = 9

type Sudoku struct {
	Regions [SIZE]*Region
}

func NewSudoku() *Sudoku {
	regions := [SIZE]*Region{}

	for i := 0; i < SIZE; i++ {
		regions[i] = NewRegion()
	}

	return &Sudoku{regions}
}

func (s *Sudoku) getRowValues(row int) []int {
	rowValues := []int{}

	regionOffset := 0
	if row > 3 {
		regionOffset = 1
	}

	if row >= 6 {
		regionOffset = 2
	}

	rowOffset := row - regionOffset

	for _, region := range s.Regions[0+regionOffset : 3+regionOffset] {
		rowValues = append(
			rowValues,
			region.Cells[0+rowOffset].Value,
			region.Cells[1+rowOffset].Value,
			region.Cells[2+rowOffset].Value,
		)
	}

	return rowValues
}

func (s *Sudoku) getColumnValues(column int) []int {
	columnValues := []int{}

	regionOffset := 0
	if column >= 3 {
		regionOffset = 1
	}

	if column >= 6 {
		regionOffset = 2
	}

	columnOffset := column - (regionOffset * 3)

	columnRegions := []*Region{
		s.Regions[0+regionOffset],
		s.Regions[3+regionOffset],
		s.Regions[6+regionOffset],
	}

	for _, region := range columnRegions {
		columnValues = append(
			columnValues,
			region.Cells[0+columnOffset].Value,
			region.Cells[3+columnOffset].Value,
			region.Cells[6+columnOffset].Value,
		)
	}

	return columnValues
}

func (s *Sudoku) getCellByNumber(number int) *Cell {
	region := number / SIZE
	cellNumber := number % SIZE

	cell := s.Regions[region].Cells[cellNumber]

	return cell
}

type Region struct {
	Cells [9]*Cell
}

func NewRegion() *Region {
	cells := [SIZE]*Cell{}
	for i := 0; i < SIZE; i++ {
		cells[i] = NewCell()
	}

	return &Region{cells}
}

func (r *Region) getCellValues() []int {
	cellValues := []int{}

	for _, cell := range r.Cells {
		cellValues = append(cellValues, cell.Value)
	}

	return cellValues
}

type Cell struct {
	// A zero means no value has been set
	Value      int
	FixedValue bool
}

func NewCell() *Cell {
	return &Cell{0, false}
}
