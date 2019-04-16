package sudoku

type Validator struct {
	allowEmptyValues bool
}

func (v *Validator) Validate(sudoku *Sudoku) bool {
	if !v.validateRows(sudoku) {
		return false
	}
	if !v.validateColumns(sudoku) {
		return false
	}
	if !v.validateRegions(sudoku) {
		return false
	}

	return true
}

func (v *Validator) validateRows(sudoku *Sudoku) bool {
	for i := 0; i < SIZE; i++ {
		if !v.validateList(sudoku.getRowValues(i)) {
			return false
		}
	}

	return true
}

func (v *Validator) validateColumns(sudoku *Sudoku) bool {
	for i := 0; i < SIZE; i++ {
		if !v.validateList(sudoku.getColumnValues(i)) {
			return false
		}
	}

	return true
}

func (v *Validator) validateRegions(sudoku *Sudoku) bool {
	for _, r := range sudoku.Regions {
		if !v.validateList(r.getCellValues()) {
			return false
		}
	}

	return true
}

func (v *Validator) validateList(list []int) bool {
	if v.allowEmptyValues {
		list = removeFromList(list, 0)
	}
	return isUniqueList(list)
}
