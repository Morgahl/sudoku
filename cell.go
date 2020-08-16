package sudoku

import (
	"errors"
)

type Cell struct {
	constrainedMemberOf []*Constraint
	viewMemberOf        []*Constraint
	val                 uint8
	solved              bool
	vals                []bool
}

func BuildPuzzleCells(p *Puzzle) (cells []*Cell) {
	cells = make([]*Cell, p.stride*p.stride)
	for i := range cells {
		cells[i] = NewCell(p.stride)
	}

	return
}

func NewCell(valCount uint8) *Cell {
	return &Cell{
		vals: NewValueSet(valCount),
	}
}

func (c *Cell) Solve(val uint8) error {
	if c.solved {
		return nil // no change
	}

	c.solved = true
	c.val = val
	for i := range c.vals {
		c.vals[i] = false
	}

	c.vals[val] = true
	for _, constraint := range c.constrainedMemberOf {
		if _, err := constraint.Solved(c); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cell) Clear(valuesToClear []uint8) (changed bool, _ error) {
	if c.solved {
		return changed, nil // no change
	}

	for _, val := range valuesToClear {
		if c.vals[val] {
			c.vals[val] = false
			changed = true
		}
	}

	if changed {
		count := 0
		idx := -1
		for i, val := range c.vals {
			if val {
				count++
				idx = i
			}
		}

		switch count {
		case 0:
			return changed, CellErrorAllValuesCleared
		case 1:
			c.solved = true
			c.val = uint8(idx)
			for _, constraint := range c.constrainedMemberOf {
				constraintChanged, err := constraint.Solved(c)
				changed = changed || constraintChanged
				if err != nil {
					return changed, err
				}
			}
		}
	}

	return changed, nil
}

var CellErrorAllValuesCleared = errors.New("cell: All Values Cleared: Likely indicates an invalid Puzzle.")
