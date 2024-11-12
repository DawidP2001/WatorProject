package main

type Fish struct {
	x int
	y int
}

func newFish(g *Game, x int, y int) *Fish {
	fish := Fish{
		x: x,
		y: y,
	}
	g.fishSlice = append(g.fishSlice, fish)

	return &fish
}

func (f *Fish) checkAvailablePositions(g *Game, xPosition int, yPosition int, maxX int, maxY int) (availablePositions []Position) {
	// Checks West
	if xPosition-1 >= 0 {
		if g.grid.locations[xPosition-1][yPosition].species != 1 {
			if g.grid.locations[xPosition-1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition - 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Checks East
	if xPosition+1 < maxX {
		if g.grid.locations[xPosition+1][yPosition].species != 1 {
			if g.grid.locations[xPosition+1][yPosition].species != 2 {
				newPosition := Position{xPosition: float32(xPosition + 1), yPosition: float32(yPosition)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Checks North
	if yPosition-1 >= 0 {
		if g.grid.locations[xPosition][yPosition-1].species != 1 {
			if g.grid.locations[xPosition][yPosition-1].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition - 1)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	// Checks South
	if yPosition+1 < maxY {
		if g.grid.locations[xPosition][yPosition+1].species != 1 {
			if g.grid.locations[xPosition][yPosition+1].species != 2 {
				newPosition := Position{xPosition: float32(xPosition), yPosition: float32(yPosition + 1)}
				availablePositions = append(availablePositions, newPosition)
			}
		}
	}
	return availablePositions
}

func (f *Fish) setNewPosition(g *Game, xPosition int, yPosition int, maxX int, maxY int) {
	availablePositions := f.checkAvailablePositions(g, xPosition, yPosition, maxX, maxY)
	if len(availablePositions) != 0 {
		newPosition := genRadomPosition(availablePositions)
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(g, newX, newY)
	} else if len(availablePositions) == 1 {
		newPosition := availablePositions[0]
		newX := int(newPosition.xPosition)
		newY := int(newPosition.yPosition)
		f.updateFishLocation(g, newX, newY)
	}
}

// Updates the location of a fish

func (f *Fish) updateFishLocation(g *Game, newX int, newY int) {
	g.grid.locations[newX][newY] = g.grid.locations[f.x][f.y]
	f.x = newX
	f.y = newY
	// If its breeding time don't erase fish
	if g.chronon%g.fishBreed != 0 {
		g.grid.locations[f.x][f.y] = *newSeacreatureEmpty()
	} else {
		newFish := newFish(g, f.x, f.y)
		g.fishSlice = append(g.fishSlice, *newFish)
		g.grid.locations[newFish.x][newFish.y] = *newSeacreatureFish(newFish)
	}

}
