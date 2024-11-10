package main

func (f *Fish) checkAvailablePositions(grid Grid) (availablePositions []Position) {
	xPosition := fish.position.xPosition
	yPosition := fish.position.yPosition
	// Checks North
	if yPosition-1 >= 0 {
		if grid.locations[xPosition][yPosition-1] == 0 {
			newPosition := Position{xPosition, yPosition - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Checks East
	if xPosition+1 < 320 {
		if grid.locations[xPosition+1][yPosition] == 0 {
			newPosition := Position{xPosition + 1, yPosition}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Checks South
	if yPosition+1 < 230 {
		if grid.locations[xPosition][yPosition+1] == 0 {
			newPosition := Position{xPosition, yPosition + 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	// Cheks West
	if xPosition-1 >= 0 {
		if grid.locations[xPosition-1][yPosition] == 0 {
			newPosition := Position{xPosition - 1, yPosition}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	return availablePositions
}

func fishNextTurn() {
	availablePositions := checkAvailablePositions()

}
