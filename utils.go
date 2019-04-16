package sudoku

func removeFromList(list []int, removeable int) []int {
	newList := []int{}
	for _, value := range list {
		if value != removeable {
			newList = append(newList, value)
		}
	}

	return newList
}

func isUniqueList(list []int) bool {
	found := map[int]bool{}
	for _, value := range list {
		if _, ok := found[value]; ok {
			return false
		}

		found[value] = true
	}

	return true
}
