package puzzle

type Constraint struct {
	constrained []*Cell
	trigger     Trigger
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
	for y := uint(0); y < p.stride; y += p.boxStride {
		for x := uint(0); x < p.stride; x += p.boxStride {
			if constraint, err = BoxConstraint(p, x, y); err != nil {
				return
			}

			constraints = append(constraints, constraint)
		}
	}

	return
}

func BoxConstraint(p *Puzzle, x, y uint) (*Constraint, error) {
	constrained := make([]*Cell, 0, p.stride)
	for sy := y; sy < y+p.boxStride; sy++ {
		for sx := x; sx < x+p.boxStride; sx++ {
			cell, err := p.At(sx, sy)
			if err != nil {
				return nil, err
			}

			constrained = append(constrained, cell)
		}
	}

	return &Constraint{
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func BuildRowConstraints(p *Puzzle) (constraints []*Constraint, err error) {
	constraints = make([]*Constraint, 0, p.stride)
	var constraint *Constraint
	for y := uint(0); y < p.stride; y++ {
		if constraint, err = RowConstraint(p, y); err != nil {
			return
		}

		constraints = append(constraints, constraint)
	}

	return
}

func RowConstraint(p *Puzzle, y uint) (*Constraint, error) {
	constrained := make([]*Cell, p.stride)
	for x := range constrained {
		cell, err := p.At(uint(x), y)
		if err != nil {
			return nil, err
		}

		constrained[x] = cell
	}

	return &Constraint{
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func BuildColumnConstraints(p *Puzzle) (constraints []*Constraint, err error) {
	constraints = make([]*Constraint, 0, p.stride)
	var constraint *Constraint
	for x := uint(0); x < p.stride; x++ {
		if constraint, err = ColumnConstraint(p, x); err != nil {
			return
		}

		constraints = append(constraints, constraint)
	}

	return
}

func ColumnConstraint(p *Puzzle, x uint) (*Constraint, error) {
	constrained := make([]*Cell, p.stride)
	for y := range constrained {
		cell, err := p.At(x, uint(y))
		if err != nil {
			return nil, err
		}

		constrained[y] = cell
	}

	return &Constraint{
		constrained: constrained,
		trigger:     StaticTrigger,
	}, nil
}

func (c *Constraint) Propagate() error {
	valuesToClear, cellsToClear := c.trigger(c.constrained)
	if len(cellsToClear) > 0 {
		for _, cellToClear := range cellsToClear {
			if err := cellToClear.Clear(valuesToClear); err != nil {
				return err
			}
		}
	}

	return nil
}
