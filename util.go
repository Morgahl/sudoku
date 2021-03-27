package sudoku

import "errors"

func solveCell(p *Puzzle, c *Cell, v uint) error {
	c.value = v
	c.solved = true
	for i := range c.notes {
		c.notes[i] = false
	}
	f := clearNote(p, v)
	if err := row(p, c, f); err != nil {
		return err
	}
	if err := col(p, c, f); err != nil {
		return err
	}
	if err := box(p, c, f); err != nil {
		return err
	}

	return nil
}

func clearNote(p *Puzzle, v uint) func(*Cell) error {
	return func(c *Cell) error {
		if c.solved {
			return nil
		}
		c.notes[v] = false
		n := 0
		idx := -1
		for i, note := range c.notes {
			if note {
				idx = i
				n += 1
			}
		}
		switch n {
		case 0:
			return errors.New("cell has all values cleared")
		case 1:
			return solveCell(p, c, uint(idx))
		default:
			return nil
		}
	}
}

func notesCount(c *Cell) []uint {
	notes := make([]uint, 0, len(c.notes))
	for i, note := range c.notes {
		if note {
			notes = append(notes, uint(i))
		}
	}

	return notes
}

func getCell(p *Puzzle, x, y uint) *Cell {
	if y*p.stride+x > uint(len(p.cells))-1 {
		return nil
	}
	return p.cells[y*p.stride+x]
}

func row(p *Puzzle, c *Cell, f func(*Cell) error) error {
	for x := uint(0); x < p.stride; x++ {
		if err := f(getCell(p, x, c.y)); err != nil {
			return err
		}
	}
	return nil
}

func col(p *Puzzle, c *Cell, f func(*Cell) error) error {
	for y := uint(0); y < p.stride; y++ {
		if err := f(getCell(p, c.x, y)); err != nil {
			return err
		}
	}
	return nil
}

func box(p *Puzzle, c *Cell, f func(*Cell) error) error {
	xmin := (c.x / p.boxstride) * p.boxstride
	ymin := (c.y / p.boxstride) * p.boxstride
	xmax := xmin + p.boxstride
	ymax := ymin + p.boxstride
	for y0 := ymin * p.boxstride; y0 < ymax; y0++ {
		for x0 := xmin * p.boxstride; x0 < xmax; x0++ {
			if err := f(getCell(p, x0, y0)); err != nil {
				return err
			}
		}
	}
	return nil
}
