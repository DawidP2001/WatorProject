package main

type Grid struct {
	// Value 0 in position = empty
	// Value 1 in poistion = fish
	// Value 2 in position = shark
	locations [320][240]Seacreature
}

func (grid *Grid) resetMovedPositions() {
	for i := 0; i < len(grid.locations); i++ {
		for j := 0; j < len(grid.locations[i]); j++ {
			if grid.locations[i][j].species == 1 {
				grid.locations[i][j].moved = false
			}
		}
	}
}
