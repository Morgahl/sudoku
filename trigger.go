package puzzle

type Trigger func(view, constrained []*Cell) (valuesToClear []uint8, cellsToClear []*Cell)

func StaticTrigger(view, constrained []*Cell) (valuesToClear []uint8, cellsToClear []*Cell) {
	for _, constrainedCell := range constrained {
		if constrainedCell.solved {
			valuesToClear = append(valuesToClear, constrainedCell.val)
			continue
		}

		cellsToClear = append(cellsToClear, constrainedCell)
	}

	return
}
