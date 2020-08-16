package sudoku

type Constraint struct {
	view        []*Cell
	constrained []*Cell
	trigger     Trigger
	solved      bool
}

func BuildPuzzleConstraints(p *Puzzle) (allConstraints []*Constraint, err error) {
	allConstraints = make([]*Constraint, 0, p.stride*3)
	var constraints []*Constraint
	if constraints, err = BuildRowConstraints(p); err != nil {
		return
	}
	allConstraints = append(allConstraints, constraints...)

	if constraints, err = BuildColumnConstraints(p); err != nil {
		return
	}
	allConstraints = append(allConstraints, constraints...)

	if constraints, err = BuildBoxConstraints(p); err != nil {
		return
	}
	allConstraints = append(allConstraints, constraints...)

	return
}

func BuildBoxConstraints(p *Puzzle) (constraints []*Constraint, err error) {
	constraints = make([]*Constraint, 0, p.stride)
	var constraint *Constraint
	for y := uint8(0); y < p.stride; y += p.boxStride {
		for x := uint8(0); x < p.stride; x += p.boxStride {
			if constraint, err = BoxConstraint(p, x, y); err != nil {
				return
			}

			constraints = append(constraints, constraint)
		}
	}

	return
}

func BoxConstraint(p *Puzzle, x, y uint8) (*Constraint, error) {
	constrained := make([]*Cell, 0, p.stride)
	for sy := y; sy < y+p.boxStride; sy++ {
		for sx := x + (sy - y); sx < x+p.boxStride; x++ {
			cell, err := p.At(uint8(sx), sy)
			if err != nil {
				return nil, err
			}

			constrained = append(constrained, cell)
		}
	}

	return &Constraint{
		view:        constrained,
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func BuildRowConstraints(p *Puzzle) (constraints []*Constraint, err error) {
	constraints = make([]*Constraint, 0, p.stride)
	var constraint *Constraint
	for y := uint8(0); y < p.stride; y++ {
		if constraint, err = RowConstraint(p, y); err != nil {
			return
		}

		constraints = append(constraints, constraint)
	}

	return
}

func RowConstraint(p *Puzzle, y uint8) (*Constraint, error) {
	constrained := make([]*Cell, p.stride)
	for x := range constrained {
		cell, err := p.At(uint8(x), y)
		if err != nil {
			return nil, err
		}

		constrained[x] = cell
	}

	return &Constraint{
		view:        constrained,
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func BuildColumnConstraints(p *Puzzle) (constraints []*Constraint, err error) {
	constraints = make([]*Constraint, 0, p.stride)
	var constraint *Constraint
	for x := uint8(0); x < p.stride; x++ {
		if constraint, err = ColumnConstraint(p, uint8(x)); err != nil {
			return
		}

		constraints = append(constraints, constraint)
	}

	return
}

func ColumnConstraint(p *Puzzle, x uint8) (*Constraint, error) {
	constrained := make([]*Cell, p.stride)
	for y := range constrained {
		cell, err := p.At(x, uint8(y))
		if err != nil {
			return nil, err
		}

		constrained[y] = cell
	}

	return &Constraint{
		view:        constrained,
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func (c *Constraint) Solved(cell *Cell) (changed bool, _ error) {
	if c.solved {
		return changed, nil // no change
	}

	for _, constrained := range c.constrained {
		if cell == constrained || constrained.solved {
			continue
		}

		cellChanged, err := cell.Clear([]uint8{cell.val})
		changed = changed || cellChanged
		if err != nil {
			return changed, err
		}
	}

	valuesToClear, cellsToClear := c.trigger(c.view, c.constrained)
	if len(cellsToClear) > 0 {
		for _, cellToClear := range cellsToClear {
			cellChanged, err := cellToClear.Clear(valuesToClear)
			changed = changed || cellChanged
			if err != nil {
				return changed, err
			}
		}
	}

	if changed {
		count := 0
		for _, constrainedCell := range c.constrained {
			if constrainedCell.solved {
				count++
			}
		}

		if count == len(c.constrained) {
			c.solved = true
		}
	}

	return changed, nil
}
