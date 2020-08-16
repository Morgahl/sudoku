package sudoku_test

import (
	"testing"

	"github.com/curlymon/sudoku"
)

func testNewPuzzleFromState(t *testing.T, state sudoku.State) {
	if _, err := sudoku.NewPuzzleFromState(state); err != nil {
		t.Fatal(err)
	}
}

func TestHardPuzzle(t *testing.T) {
	state, err := sudok.LoadStateFromFile("./hard.json")
	if err != nil {
		t.Fatal(err)
	}

	testNewPuzzleFromState(t, state)
}
