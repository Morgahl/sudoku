package puzzle

import (
	"errors"
)

type Puzzle struct {
	stride      uint
	boxStride   uint
	cells       []*Cell
	constraints []*Constraint
}

func NewPuzzle(boxStride uint) (p *Puzzle, err error) {
	p = &Puzzle{
		stride:    boxStride * boxStride,
		boxStride: boxStride,
	}

	p.cells = BuildPuzzleCells(p)

	var constraints []*Constraint
	if constraints, err = BuildPuzzleConstraints(p); err != nil {
		return
	}

	for _, constraint := range constraints {
		p.ApplyConstraint(constraint)
	}

	return
}

func NewPuzzleFromState(state State) (p *Puzzle, err error) {
	if p, err = NewPuzzle(state.Dim); err != nil {
		return
	}

	for y, row := range state.Puzzle {
		for x, val := range row {
			if val == 0 {
				continue
			}
			// offset down into indexable representation
			val--

			p.Set(uint(x), uint(y), val)
		}
	}

	return
}

func (p *Puzzle) State() (s State, err error) {
	s.Dim = p.boxStride
	s.Puzzle = make([][]uint8, p.stride)
	var cell *Cell
	for y := uint(0); y < p.stride; y++ {
		s.Puzzle[y] = make([]uint8, p.stride)
		for x := uint(0); x < p.stride; x++ {
			if cell, err = p.At(x, y); err != nil {
				return
			}
			val := cell.val
			if cell.solved {
				// offset up into human representation
				val++
			}
			s.Puzzle[y][x] = val
		}
	}

	return
}

func (p *Puzzle) ApplyConstraint(constraint *Constraint) {
	for _, existingConstraint := range p.constraints {
		if existingConstraint == constraint {
			return
		}
	}

	p.constraints = append(p.constraints, constraint)

	for _, cell := range constraint.constrained {
		cell.constrainedMemberOf = append(cell.constrainedMemberOf, constraint)
	}

	return
}

func (p Puzzle) At(x, y uint) (*Cell, error) {
	idx := (y * p.stride) + x
	if int(idx) >= len(p.cells) {
		return nil, PuzzleErrorInvalidCell
	}

	return p.cells[idx], nil
}

func (p *Puzzle) Set(x, y uint, v uint8) (err error) {
	var cell *Cell
	if cell, err = p.At(x, y); err != nil {
		return
	}

	return cell.Solve(v)
}

func (p Puzzle) IsSolved() bool {
	for i := range p.cells {
		if !p.cells[i].solved {
			return false
		}
	}

	return true
}

func (p Puzzle) Copy() (np *Puzzle, err error) {
	np, err = NewPuzzle(p.boxStride)
	if err != nil {
		return nil, err
	}
	for i := range np.cells {
		if cell := p.cells[i]; cell.solved {
			np.cells[i].Solve(cell.val)
		}
	}

	return np, nil
}

var PuzzleErrorInvalidCell = errors.New("puzzle: Requested Cell is Invalid")
