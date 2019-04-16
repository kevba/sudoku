package sudoku

import (
	"io/ioutil"
	"testing"
)

func TestFromJSON(t *testing.T) {
	data, _ := ioutil.ReadFile("json.golden")

	_, err := FromJSON(data)
	if err != nil {
		t.Errorf("%v", err)
	}
}
