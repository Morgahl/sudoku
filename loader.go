package sudoku

import (
	"encoding/json"
	"os"
)

// State ...
type State struct {
	Dim    uint8     `json:"dim"`
	Puzzle [][]uint8 `json:"puzzle"`
}

func LoadStateFromFile(path string) (s State, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&s)
	return
}
