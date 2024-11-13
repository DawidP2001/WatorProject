package main

import (
	"math/rand"
)

type World struct {
	width     int
	height    int
	ncells    int
	grid      [250][250]*Creature
	creatures []*Creature
	nfish     int
	nsharks   int
}

func newWorld(width, height, nfish, nsharks int) *World {
	w := &World{
		width:  width,
		height: height,
		ncells: width * height,
	}
	w.fillTheGrid()
	w.populateWorld(nfish, nsharks)
	return w
}

// Fills the world with creature objects
func (w *World) fillTheGrid() {
	for i := 0; i < w.width; i++ {
		for j := 0; j < w.height; j++ {
			w.grid[i][j] = newCreatureEmpty(i, j)
		}
	}
}
func (w *World) spawnCreature(creatureId, x, y int) {
	creature := newCreature(
		creatureId, x, y,
		initialEnergies[creatureId-1],
		fertilityThresholds[creatureId-1])
	w.creatures = append(w.creatures, creature)
	w.grid[x][y] = creature
}

func (w *World) populateWorld(nfish, nsharks int) {
	w.nfish, w.nsharks = nfish, nsharks

	w.placeCreatures(w.nfish, FISH)
	w.placeCreatures(w.nsharks, SHARK)
}

func (w *World) placeCreatures(ncreatures, creatureId int) {
	for i := 0; i < ncreatures; i++ {
		finish := false
		for !finish {
			randomX := rand.Intn(w.width)
			randomY := rand.Intn(w.height)
			if w.grid[randomX][randomY].id == 0 {
				w.spawnCreature(creatureId, randomX, randomY)
				finish = true
			}
		}
	}
}
func (w *World) get_neighbours(x, y int) [4]*Creature {
	movements := [][2]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
	var neighbours [4]*Creature
	for i, movement := range movements {
		directionX := movement[0]
		directionY := movement[1]

		newX := (x + directionX + w.width) % w.width
		newY := (y + directionY + w.height) % w.height
		neighbours[i] = w.grid[newX][newY]
	}
	return neighbours
}

// Moves Creatures
func (w *World) evolveCreatures(creature *Creature) {
	var newX int
	var newY int
	neighbours := w.get_neighbours(creature.x, creature.y)
	creature.fertility += 1
	moved := false
	if creature.id == SHARK {
		if checkIfAnyNeighbourIsFood(neighbours) {
			neighbours := getFoodNeighbours(neighbours)
			neighbour := randomiseNeighbour(neighbours)
			newX = neighbour.x
			newY = neighbour.y
			creature.energy += 2
			w.grid[newX][newY].dead = true
			w.grid[newX][newY] = newCreatureEmpty(newX, newY)
			moved = true
		}
	}
	if !moved {
		emptyNeighbours := getEmptyNeighbours(neighbours)
		if len(emptyNeighbours) != 0 {
			randomNeighbor := randomiseNeighbour(emptyNeighbours)
			newX = randomNeighbor.x
			newY = randomNeighbor.y
			if creature.id == SHARK {
				creature.energy--
			}
			moved = true
		}
	}
	if creature.energy < 0 {
		creature.dead = true
		w.grid[creature.x][creature.y].id = 0
	} else if moved {
		x := creature.x
		y := creature.y
		creature.x = newX
		creature.y = newY
		w.grid[newX][newY] = creature
		if creature.fertility >= creature.fishBreed {
			creature.fertility = 0
			w.spawnCreature(creature.id, x, y)
		} else {
			w.grid[x][y] = newCreatureEmpty(x, y)
		}
	}
}
func (w *World) evolveWorld() {
	// Shuffles the creature slice
	rand.Shuffle(len(w.creatures),
		func(i, j int) { w.creatures[i], w.creatures[j] = w.creatures[j], w.creatures[i] })

	ncreatures := len(w.creatures)
	for i := 0; i < ncreatures; i++ {
		creature := w.creatures[i]
		if creature.dead {
			continue
		}
		w.evolveCreatures(creature)
	}
	var newCreatures []*Creature
	for i := 0; i < len(w.creatures); i++ {
		if !w.creatures[i].dead {
			newCreatures = append(newCreatures, w.creatures[i])
		}
	}
	w.creatures = newCreatures
}

func getEmptyNeighbours(neighbours [4]*Creature) []*Creature {
	var emptyNeighbours []*Creature
	for i := 0; i < 4; i++ {
		if neighbours[i].id == 0 {
			emptyNeighbours = append(emptyNeighbours, neighbours[i])
		}
	}
	return emptyNeighbours
}

func randomiseNeighbour(neighbours []*Creature) *Creature {
	return neighbours[rand.Intn(len(neighbours))]
}

func checkIfAnyNeighbourIsFood(neighbours [4]*Creature) bool {
	foodPresent := false
	for i := 0; i < 4; i++ {
		if neighbours[i].id == 1 {
			foodPresent = true
		}
	}
	return foodPresent
}

func getFoodNeighbours(neighbours [4]*Creature) []*Creature {
	var neighbour []*Creature
	for i := 0; i < 4; i++ {
		if neighbours[i].id == 1 {
			neighbour = append(neighbour, neighbours[i])
		}
	}
	return neighbour
}
