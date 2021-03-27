package sudoku

type Cell struct {
	value  uint
	x      uint
	y      uint
	notes  []bool
	solved bool
}

func newCell(x, y, n uint) *Cell {
	notes := make([]bool, n)
	for i := range notes {
		notes[i] = true
	}

	return &Cell{x: x, y: y, notes: notes}
}

func (c *Cell) copy() *Cell {
	notes := make([]bool, len(c.notes))
	for i, note := range c.notes {
		notes[i] = note
	}

	return &Cell{
		value:  c.value,
		x:      c.x,
		y:      c.y,
		notes:  notes,
		solved: c.solved,
	}
}
