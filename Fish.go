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
	g.fishSlice = append(g.fishSlice, &fish)

	return &fish
}

// Removes a fish from the fishSlice
func removeFish(g *Game, fish *Fish) {
	var newSlice []*Fish
	g.semChannel <- true
	for i := 0; i < len(g.fishSlice); i++ {
		if !(g.fishSlice[i].x == fish.x && g.fishSlice[i].y == fish.y) {
			newSlice = append(newSlice, g.fishSlice[i])
		}
	}
	g.fishSlice = newSlice
	<-g.semChannel
}
func (f *Fish) checkAvailablePositions(g *Game, maxX int, maxY int) (availablePositions []Position) {
	// Checks North
	if f.y-1 >= 0 {
		if g.grid.locations[f.x][f.y-1].species == 0 {
			newPosition := Position{xPosition: f.x, yPosition: f.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[f.x][maxY-1].species == 0 {
			newPosition := Position{xPosition: f.x, yPosition: maxY - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks East
	if f.x+1 < maxX {
		if g.grid.locations[f.x+1][f.y].species == 0 {
			newPosition := Position{xPosition: f.x + 1, yPosition: f.y}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[0][f.y].species == 0 {
			newPosition := Position{xPosition: 0, yPosition: f.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks South
	if f.y+1 < maxY {
		if g.grid.locations[f.x][f.y+1].species == 0 {
			newPosition := Position{xPosition: f.x, yPosition: f.y + 1}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[f.x][0].species == 0 {
			newPosition := Position{xPosition: f.x, yPosition: 0}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks West
	if f.x-1 >= 0 {
		if g.grid.locations[f.x-1][f.y].species == 0 {
			newPosition := Position{xPosition: f.x - 1, yPosition: f.y}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[maxX-1][f.y].species == 0 {
			newPosition := Position{xPosition: maxX - 1, yPosition: f.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	return availablePositions
}

func (f *Fish) updateFishPosition(g *Game, maxX int, maxY int) {
	availablePositions := f.checkAvailablePositions(g, maxX, maxY)
	if len(availablePositions) != 0 {
		newPosition := genRadomPosition(availablePositions)
		newX := newPosition.xPosition
		newY := newPosition.yPosition
		f.setNewPosition(g, newX, newY)
	}
}

// Updates the location of a fish

func (f *Fish) setNewPosition(g *Game, newX int, newY int) {
	g.grid.locations[newX][newY] = g.grid.locations[f.x][f.y]
	f.x = newX
	f.y = newY
	// If its breeding time don't erase fish
	if g.chronon%g.fishBreed != 0 {
		g.grid.locations[f.x][f.y] = *newSeacreatureEmpty()
	} else {
		newFish := newFish(g, f.x, f.y)
		g.fishSlice = append(g.fishSlice, newFish)
		g.grid.locations[newFish.x][newFish.y] = *newSeacreatureFish(newFish)
	}

}
