package sudoku

import (
	"testing"
)

func testNewPuzzleFromStateFile(t *testing.T, stateFile string) {
	state, err := LoadStateFromFile(stateFile)
	if err != nil {
		t.Fatal(err)
	}

	p, err := NewPuzzleFromState(state)
	if err != nil {
		t.Fatal(err)
	}
	defer p.Destroy()

	if !p.IsSolved() {
		t.Fatal("Puzzle was not solved...")
	}

	afterSolveState, _ := p.State()
	t.Logf("%#v", afterSolveState)
}

func TestEasyyPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./easy.json")
}
func TestHardPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./hard.json")
}
