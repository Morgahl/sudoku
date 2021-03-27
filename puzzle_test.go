package sudoku_test

import (
	"testing"

	"github.com/curlymon/sudoku"
)

func TestEasyPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./fixtures/easy.json")
}
func TestHardPuzzle(t *testing.T) {
	testNewPuzzleFromStateFile(t, "./fixtures/hard.json")
}

func testNewPuzzleFromStateFile(t *testing.T, stateFile string) {
	state, err := sudoku.LoadStateFromFile(stateFile)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Start State:\n%s", state)

	p, err := sudoku.NewPuzzleFromState(state)
	if err != nil {
		t.Fatal(err)
	}

	if !p.IsSolved() {
		t.Log("attempting recursive solve")
		p, err = sudoku.Solve(p)
		if err != nil {
			t.Errorf("error solving: %s\n%v", err, p)
		}
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
	state, err := sudoku.LoadStateFromFile(stateFile)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		var p *sudoku.Puzzle
		for pb.Next() {
			p, _ = sudoku.NewPuzzleFromState(state)
			p, _ = sudoku.Solve(p)
		}
	})
}
