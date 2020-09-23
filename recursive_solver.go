package puzzle

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
	var vals []uint8
	for i, cell := range p.cells {
		if cell.solved {
			continue
		}
		if vals = cell.ValuesCount(); len(vals) < lowestCount || lowestCount == 0 {
			idx = i
			lowestCount = len(vals)
			continue
		}
	}

	for _, val := range vals {
		np, err = p.Copy()
		if err != nil {
			return nil, err
		}

		cell := np.cells[idx]
		cell.Solve(val)
		if np, err = Solve(np); err == nil {
			return np, nil
		}
	}

	// return failed to solve error as this cannot be solved on this path
	return nil, PuzzleErrorInvalidSolution
}

var PuzzleErrorInvalidSolution = errors.New("puzzle: Puzzle Solution is Invalid")
