package sudoku

import (
	"errors"
)

func Solve(p *Puzzle) (np *Puzzle, err error) {
	if p.IsSolved() {
		return p, nil
	}

	// find first cell with lowest count to solve
	idx := 0
	lowestCount := 0
	var notes []uint
	for i, c := range p.cells {
		if c.solved {
			continue
		}
		if notes = notesCount(c); len(notes) < lowestCount || lowestCount == 0 {
			idx = i
			lowestCount = len(notes)
			continue
		}
	}

	for _, note := range notes {
		np = p.copy()
		solveCell(np, np.cells[idx], note)
		if np, err = Solve(np); err == nil {
			return np, nil
		}
	}

	// return failed to solve error as this cannot be solved on this path
	return nil, errors.New("puzzle: Puzzle Solution is Invalid")
}
