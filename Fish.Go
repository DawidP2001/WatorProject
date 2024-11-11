package main

type Fish struct {
}

func newFish() *Fish {
	return &Fish{}
}
func (f *Fish) checkAvailablePositions(g *Game, xPosition int, yPosition int, maxX int, maxY int) (availablePositions []Position) {
	// Checks West
	print("Chronos ")
	print(g.chronon)
	print(" ")
	if xPosition-1 >= 0 {
		if g.grid.locations[xPosition-1][yPosition].species != 1 {
			if g.grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition - 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
				print("W")
			}
		}
	}
	// Checks East
	if xPosition+1 < maxX {
		if g.grid.locations[xPosition+1][yPosition].species != 1 {
			if g.grid.locations[xPosition+1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition + 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
				print("E")
			}
		}
	}
	// Checks North
	if yPosition-1 >= 0 {
		if g.grid.locations[xPosition][yPosition-1].species != 1 {
			if g.grid.locations[xPosition][yPosition-1].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition - 1)}
				availablePositions = append(availablePositions, newPosition)
				print("N")
			}
		}
	}
	// Cheks South
	if yPosition+1 < maxY {
		if g.grid.locations[xPosition][yPosition+1].species != 1 {
			if g.grid.locations[xPosition][yPosition+1].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition + 1)}
				availablePositions = append(availablePositions, newPosition)
				print("S")
			}
		}
	}
	print("\n")
	return availablePositions
}
func (f *Fish) setNewPosition(g *Game, xPosition int, yPosition int, maxX int, maxY int) {
	availablePositions := f.checkAvailablePositions(g, xPosition, yPosition, maxX, maxY)
	if len(availablePositions) != 0 {
		newPosition := genRadomPosition(availablePositions)
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(g, newX, newY, xPosition, yPosition)
	} else if len(availablePositions) == 1 {
		newPosition := availablePositions[0]
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(g, newX, newY, xPosition, yPosition)
	}
}

// Updates the location of a fish
func (f *Fish) updateFishLocation(g *Game, newX int, newY int, oldX int, oldY int) {
	g.grid.locations[oldX][oldY].moved = true
	g.grid.locations[newX][newY] = g.grid.locations[oldX][oldY]
	// If its breeding time don't erase fish
	if g.chronon%g.fishBreed != 0 {
		g.grid.locations[oldX][oldY].species = 0
	}
}
