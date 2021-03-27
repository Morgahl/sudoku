package sudoku

import (
	"math"
)

type Puzzle struct {
	stride    uint
	boxstride uint
	cells     []*Cell
}

func New(n uint) *Puzzle {
	cells := make([]*Cell, n*n)
	for y := uint(0); y < n; y++ {
		for x := uint(0); x < n; x++ {
			cells[y*n+x] = newCell(x, y, n)
		}
	}

	return &Puzzle{
		stride:    n,
		boxstride: uint(math.Sqrt(float64(n))),
		cells:     cells,
	}
}

func NewPuzzleFromState(s State) (p *Puzzle, err error) {
	p = New(s.Dim)
	for y := uint(0); y < s.Dim; y++ {
		for x := uint(0); x < s.Dim; x++ {
			if v := s.Puzzle[y][x]; v > 0 {
				v-- // offset down into indexable representation
				p.SolveCell(x, y, v)
			}
		}
	}

	return
}

func (p *Puzzle) SolveCell(x, y, v uint) error {
	return solveCell(p, getCell(p, x, y), v)
}

func (p *Puzzle) IsSolved() bool {
	for _, cell := range p.cells {
		if !cell.solved {
			return false
		}
	}

	return true
}

func (p *Puzzle) copy() *Puzzle {
	cells := make([]*Cell, len(p.cells))
	for i, cell := range p.cells {
		cells[i] = cell.copy()
	}

	return &Puzzle{
		stride:    p.stride,
		boxstride: p.boxstride,
		cells:     cells,
	}
}

func (p *Puzzle) State() (s State, err error) {
	s.Dim = p.stride
	s.Puzzle = make([][]uint, p.stride)
	var c *Cell
	for y := uint(0); y < p.stride; y++ {
		s.Puzzle[y] = make([]uint, p.stride)
		for x := uint(0); x < p.stride; x++ {
			c = getCell(p, x, y)
			val := c.value
			if c.solved {
				// offset up into human representation
				val++
			}
			s.Puzzle[y][x] = val
		}
	}

	return
}
