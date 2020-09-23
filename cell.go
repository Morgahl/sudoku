package puzzle

import (
	"errors"
)

type Cell struct {
	val    uint8
	solved bool
	vals   []bool

	constrainedMemberOf []*Constraint
}

func BuildPuzzleCells(p *Puzzle) (cells []*Cell) {
	cells = make([]*Cell, p.stride*p.stride)
	for y := uint(0); y < p.stride; y++ {
		for x := uint(0); x < p.stride; x++ {
			cells[(y*p.stride)+x] = NewCell(p.stride)
		}
	}

	return
}

func NewCell(valCount uint) *Cell {
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
	c.vals = nil

	for _, constraint := range c.constrainedMemberOf {
		if err := constraint.Propagate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cell) Clear(valuesToClear []uint8) error {
	if c.solved {
		return nil // no change
	}

	for _, val := range valuesToClear {
		if c.vals[val] {
			c.vals[val] = false
		}
	}

	count := 0
	idx := uint8(0)
	for i, val := range c.vals {
		if val {
			count++
			idx = uint8(i)
			if count > 1 {
				return nil
			}
		}
	}

	switch count {
	case 0:
		return CellErrorAllValuesCleared
	case 1:
		return c.Solve(idx)
	}

	return nil
}

func (c Cell) ValuesCount() []uint8 {
	vals := make([]uint8, 0, len(c.vals))
	for i, exists := range c.vals {
		if exists {
			vals = append(vals, uint8(i))
		}
	}

	return vals
}

var CellErrorAllValuesCleared = errors.New("cell: All Values Cleared: Likely indicates an invalid Puzzle.")
