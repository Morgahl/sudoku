package sudoku

import (
	"testing"
)

func TestEasyyPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./easy.json")
}
func TestHardPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./hard.json")
}

func testNewPuzzleFromStateFile(t *testing.T, stateFile string) {
	state, err := LoadStateFromFile(stateFile)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Start State:%s", state)

	p, err := NewPuzzleFromState(state)
	if err != nil {
		t.Fatal(err)
	}

	if !p.IsSolved() {
		t.Error("Puzzle was not solved...")
	}

	afterSolveState, _ := p.State()
	t.Logf("End State:%s", afterSolveState)
}

func BenchmarkEasyyPuzzle(b *testing.B) {
	benchmarkNewPuzzleFromStateFile(b, "./easy.json")
}
func BenchmarkHardPuzzle(b *testing.B) {
	benchmarkNewPuzzleFromStateFile(b, "./hard.json")
}

func benchmarkNewPuzzleFromStateFile(b *testing.B, stateFile string) {
	state, err := LoadStateFromFile(stateFile)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		var p *Puzzle
		for pb.Next() {
			p, _ = NewPuzzleFromState(state)
		}
		p = p
	})
}
