package puzzle

import (
	"testing"
)

func TestEasyPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./fixtures/easy.json")
}
func TestHardPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./fixtures/hard.json")
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

	var finalDepth int
	if !p.IsSolved() {
		t.Log("attempting recursive solve")
		p, finalDepth, err = SolveRecursively(p, 0)
		if err != nil {
			t.Error(err)
		}
		t.Logf("finalDepth: %d", finalDepth)
	}

	afterSolveState, _ := p.State()
	t.Logf("End State:%s", afterSolveState)
}

func BenchmarkEasyyPuzzle(b *testing.B) {
	benchmarkNewPuzzleFromStateFile(b, "./fixtures/easy.json")
}
func BenchmarkHardPuzzle(b *testing.B) {
	benchmarkNewPuzzleFromStateFile(b, "./fixtures/hard.json")
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
			p, _, _ = SolveRecursively(p, 0)
		}
	})
}
