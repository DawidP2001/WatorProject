package main

type Fish struct {
}

func newFish() *Fish {
	return &Fish{}
}
func (f *Fish) checkAvailablePositions(grid *Grid, xPosition int, yPosition int, maxX int, maxY int) (availablePositions []Position) {
	// Checks West
	if xPosition-1 >= 0 {
		if grid.locations[xPosition-1][yPosition].species != 1 {
			if grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition - 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Checks East
	if xPosition+1 < maxX {
		if grid.locations[xPosition+1][yPosition].species != 1 {
			if grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition + 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Checks North
	if yPosition-1 >= 0 {
		if grid.locations[xPosition][yPosition-1].species != 1 {
			if grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition - 1)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Cheks South
	if yPosition+1 < maxY {
		if grid.locations[xPosition-1][yPosition+1].species != 1 {
			if grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition + 1)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	return availablePositions
}
func (f *Fish) setNewPosition(grid *Grid, xPosition int, yPosition int, maxX int, maxY int) {
	availablePositions := f.checkAvailablePositions(grid, xPosition, yPosition, maxX, maxY)
	if len(availablePositions) == 0 {

	} else if len(availablePositions) == 1 {
		newPosition := availablePositions[0]
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(grid, newX, newY, xPosition, yPosition)
	} else {
		newPosition := genRadomPosition(availablePositions)
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(grid, newX, newY, xPosition, yPosition)
	}
}
func (f *Fish) fishNextPositionTest(xPosition int, yPoistion int) (int, int) {
	//availablePositions := checkAvailablePositions()
	//newPosition := genRadomPosition(availablePositions)
	if xPosition < 319 {
		if yPoistion > 0 {
			return xPosition, yPoistion - 1
		}
	}
	return xPosition, yPoistion
}

// Updates the location of a fish
func (f *Fish) updateFishLocation(grid *Grid, newX int, newY int, oldX int, oldY int) {
	grid.locations[oldX][oldY].moved = true
	grid.locations[newX][newY] = grid.locations[oldX][oldY]
	grid.locations[oldX][oldY].species = 0
}
