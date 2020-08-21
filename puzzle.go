package puzzle

import (
	"errors"
)

type Puzzle struct {
	stride      uint8
	boxStride   uint8
	cells       []*Cell
	constraints []*Constraint
}

func NewPuzzle(boxStride uint8) (p *Puzzle, err error) {
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
		if err = p.ApplyConstraint(constraint); err != nil {
			return
		}
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

			p.Set(uint8(x), uint8(y), val)
		}
	}

	return
}

func (p *Puzzle) State() (s State, err error) {
	s.Dim = p.boxStride
	s.Puzzle = make([][]uint8, p.stride)
	var cell *Cell
	for y := uint8(0); y < p.stride; y++ {
		s.Puzzle[y] = make([]uint8, p.stride)
		for x := uint8(0); x < p.stride; x++ {
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

func (p *Puzzle) ApplyConstraint(constraint *Constraint) error {
	for _, existingConstraint := range p.constraints {
		if existingConstraint == constraint {
			return nil
		}
	}

	p.constraints = append(p.constraints, constraint)

	for _, cell := range constraint.constrained {
		cell.constrainedMemberOf = append(cell.constrainedMemberOf, constraint)
	}

	for _, cell := range constraint.view {
		cell.viewMemberOf = append(cell.viewMemberOf, constraint)
	}

	return nil
}

func (p Puzzle) At(x, y uint8) (*Cell, error) {
	idx := (y * p.stride) + x
	if int(idx) >= len(p.cells) {
		return nil, PuzzleErrorInvalidCell
	}

	return p.cells[idx], nil
}

func (p *Puzzle) Set(x, y, v uint8) (err error) {
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
