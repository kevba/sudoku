package sudoku

import "testing"

func TestIsUniqueList(t *testing.T) {
	tests := []struct {
		list               []int
		expectedToBeUnique bool
	}{
		{[]int{0}, true},
		{[]int{0, 0}, false},
		{[]int{0, 1, 3, 4, 5, 6}, true},
		{[]int{0, 1, 1, 4, 5, 6}, false},
		{[]int{}, true},
	}

	for _, test := range tests {
		isUnique := isUniqueList(test.list)
		if isUnique != test.expectedToBeUnique {
			t.Errorf("expected list to be %v. list is %v", test.expectedToBeUnique, test.list)
		}
	}
}
