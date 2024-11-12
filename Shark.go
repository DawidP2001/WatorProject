package main

type Shark struct {
	energyLeft int // Stores the amount of energy a shark has left
	x          int
	y          int
}

func newShark(g *Game, x int, y int) *Shark {
	shark := Shark{
		energyLeft: g.starve,
		x:          x,
		y:          y,
	}
	g.sharkSlice = append(g.sharkSlice, &shark)
	return &shark
}
func removeShark(g *Game, shark *Shark) {
	var newSlice []*Shark
	g.semChannel <- true
	for _, value := range g.sharkSlice {
		if !(value.x == shark.x && value.y == shark.y) {
			newSlice = append(newSlice, value)
		}
	}
	g.sharkSlice = newSlice
	<-g.semChannel
}

// Checks which positions are free (this is when none surrounding locations have food)
func (s *Shark) checkAvailablePositions(g *Game, maxX int, maxY int) (availablePositions []Position) {
	// Checks North
	if s.y-1 >= 0 {
		if g.grid.locations[s.x][s.y-1].species == 0 {
			newPosition := Position{xPosition: s.x, yPosition: s.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[s.x][maxY-1].species == 0 {
			newPosition := Position{xPosition: s.x, yPosition: maxY - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks East
	if s.x+1 < maxX {
		if g.grid.locations[s.x+1][s.y].species == 0 {
			newPosition := Position{xPosition: s.x + 1, yPosition: s.y}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[0][s.y].species == 0 {
			newPosition := Position{xPosition: 0, yPosition: s.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks South
	if s.y+1 < maxY {
		if g.grid.locations[s.x][s.y+1].species == 0 {
			newPosition := Position{xPosition: s.x, yPosition: s.y + 1}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[s.x][0].species == 0 {
			newPosition := Position{xPosition: s.x, yPosition: 0}
			availablePositions = append(availablePositions, newPosition)
		}
	}

	// Checks West
	if s.x-1 >= 0 {
		if g.grid.locations[s.x-1][s.y].species == 0 {
			newPosition := Position{xPosition: s.x - 1, yPosition: s.y}
			availablePositions = append(availablePositions, newPosition)
		}
	} else {
		if g.grid.locations[maxX-1][s.y].species == 0 {
			newPosition := Position{xPosition: maxX - 1, yPosition: s.y - 1}
			availablePositions = append(availablePositions, newPosition)
		}
	}
	return availablePositions
}

// Checks surrouding area for food
func (s *Shark) checkAvailableFood(g *Game, maxX int, maxY int) (availableFood []Position) {
	// Checks North
	if s.y-1 >= 0 {
		if g.grid.locations[s.x][s.y-1].species == 1 {
			newPosition := Position{xPosition: s.x, yPosition: s.y - 1}
			availableFood = append(availableFood, newPosition)
		}
	} else {
		if g.grid.locations[s.x][maxY-1].species == 1 {
			newPosition := Position{xPosition: s.x, yPosition: maxY - 1}
			availableFood = append(availableFood, newPosition)
		}
	}

	// Checks East
	if s.x+1 < maxX {
		if g.grid.locations[s.x+1][s.y].species == 1 {
			newPosition := Position{xPosition: s.x + 1, yPosition: s.y}
			availableFood = append(availableFood, newPosition)
		}
	} else {
		if g.grid.locations[0][s.y].species == 1 {
			newPosition := Position{xPosition: 0, yPosition: s.y - 1}
			availableFood = append(availableFood, newPosition)
		}
	}

	// Checks South
	if s.y+1 < maxY {
		if g.grid.locations[s.x][s.y+1].species == 1 {
			newPosition := Position{xPosition: s.x, yPosition: s.y + 1}
			availableFood = append(availableFood, newPosition)
		}
	} else {
		if g.grid.locations[s.x][0].species == 1 {
			newPosition := Position{xPosition: s.x, yPosition: 0}
			availableFood = append(availableFood, newPosition)
		}
	}

	// Checks West
	if s.x-1 >= 0 {
		if g.grid.locations[s.x-1][s.y].species == 1 {
			newPosition := Position{xPosition: s.x - 1, yPosition: s.y}
			availableFood = append(availableFood, newPosition)
		}
	} else {
		if g.grid.locations[maxX-1][s.y].species == 1 {
			newPosition := Position{xPosition: maxX - 1, yPosition: s.y - 1}
			availableFood = append(availableFood, newPosition)
		}
	}
	return availableFood
}

// Eats
func (s *Shark) eat(g *Game, newX int, newY int) {
	s.energyLeft = g.starve // Increase sharks energy
	fish := g.grid.locations[newX][newY].fish
	removeFish(g, fish)
	g.grid.locations[newX][newY] = g.grid.locations[s.x][s.y]
	s.x = newX
	s.y = newY
	if g.chronon%g.fishBreed != 0 {
		g.grid.locations[s.x][s.y] = *newSeacreatureEmpty()
	} else {
		s.makeNewShark(g)
	}
}

// Updates shark position
func (s *Shark) updateSharkPosition(g *Game, maxX int, maxY int) {
	availableFood := s.checkAvailableFood(g, maxX, maxY)
	if len(availableFood) != 0 {
		newPosition := genRadomPosition(availableFood)
		newX := newPosition.xPosition
		newY := newPosition.yPosition
		s.eat(g, newX, newY)
	} else {
		availablePositions := s.checkAvailablePositions(g, maxX, maxY)
		if len(availablePositions) != 0 {
			newPosition := genRadomPosition(availablePositions)
			newX := newPosition.xPosition
			newY := newPosition.yPosition
			s.setNewPosition(g, newX, newY)
		}
	}
}

// Set new Position
func (s *Shark) setNewPosition(g *Game, newX int, newY int) {
	g.semChannel <- true
	g.grid.locations[newX][newY] = g.grid.locations[s.x][s.y]
	s.x = newX
	s.y = newY
	// If its breeding time don't erase fish
	if g.chronon%g.fishBreed != 0 {
		g.grid.locations[s.x][s.y] = *newSeacreatureEmpty()
	} else {
		s.makeNewShark(g)
	}
	<-g.semChannel
}

// Make new Shark
func (s *Shark) makeNewShark(g *Game) {
	newShark := newShark(g, s.x, s.y)
	g.sharkSlice = append(g.sharkSlice, newShark)
	g.grid.locations[newShark.x][newShark.y] = *newSeacreatureShark(newShark)
}

// Checks whether the shark starved
func (s *Shark) checkStarve(g *Game) {
	if s.energyLeft < 1 {
		g.grid.locations[s.x][s.y] = *newSeacreatureEmpty()
		removeShark(g, s)
	}
}

// Uses energy per chronon
func (s *Shark) useEnergy() {
	s.energyLeft--
}

// Updates the shark
func (s *Shark) updateShark(g *Game, maxX int, maxY int) {
	s.updateSharkPosition(g, maxX, maxY) //Checks Position
	s.useEnergy()                        // Uses energy
	s.checkStarve(g)                     // Checks if no energy left
}
