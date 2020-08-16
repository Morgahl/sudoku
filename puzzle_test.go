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

	p, err := NewPuzzleFromState(state)
	if err != nil {
		t.Fatal(err)
	}
	defer p.IsSolved()

	// if !p.IsSolved() {
	// 	t.Fatal("Puzzle was not solved...")
	// }

	afterSolveState, _ := p.State()
	t.Logf("%#v", afterSolveState)
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
