package main

type Fish struct {
	position Position
}

func newFish(position Position, grid *Grid) *Fish {
	grid.locations[poistion.xPosition][poistion.yPosition] = 1
	return &Fish{
		poistion: poistion
	}
}
func (f *Fish) checkAvailablePositions(grid *Grid) (availablePositions []Position) {
	xPosition := f.position.xPosition
	yPosition := f.position.yPosition
	// Checks North
	if yPosition-1 >= 0 {
		if grid.locations[int(xPosition)][int(yPosition-1)] == nil {
			newPosition := Position{xPosition, yPosition - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Checks East
	if xPosition+1 < 320 {
		if grid.locations[int(xPosition+1)][int(yPosition)] == nil {
			newPosition := Position{xPosition + 1, yPosition}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Checks South
	if yPosition+1 < 230 {
		if grid.locations[int(xPosition)][int(Position+1)] == nil {
			newPosition := Position{xPosition, yPosition + 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Cheks West
	if xPosition-1 >= 0 {
		if grid.locations[int(xPosition-1)][int(yPosition)] == nil {
			newPosition := Position{xPosition - 1, yPosition}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	return availablePositions
}

func (f *Fish) fishNextTurn() {
	availablePositions := checkAvailablePositions()
	newPosition := genRadomPosition(availablePositions)
	f.position = newPosition
}
