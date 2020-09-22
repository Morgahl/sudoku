package puzzle

type Trigger func(constrained []*Cell) (valuesToClear []uint8, cellsToClear []*Cell)

func StaticTrigger(constrained []*Cell) (valuesToClear []uint8, cellsToClear []*Cell) {
	for _, constrainedCell := range constrained {
		if constrainedCell.solved {
			valuesToClear = append(valuesToClear, constrainedCell.val)
			continue
		}

		cellsToClear = append(cellsToClear, constrainedCell)
	}

	return
}
