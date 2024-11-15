package main

import (
	"math/rand"
	"sync"
)

type World struct {
	width      int
	height     int
	grid       [][]*Creature
	creatures  []*Creature
	fishBreed  int
	sharkBreed int
	starve     int
}

/**
* @brief Creates a new struct of type Creature that will be used as an empty location on the grid
*
* This function is a constructor for the struct Creature, it takes several int variables and assigns it to this new instance.
*
* @param x X coordinate
* @param y Y coordinate
* @return A pointer towards a newly created struct of type Creature that is of type 0
 */
func newWorld(numShark, numFish, fishBreed, sharkBreed, starve int, gridSize [2]int) *World {
	width := gridSize[0]
	height := gridSize[1]
	w := &World{
		width:      width,
		height:     height,
		fishBreed:  fishBreed,
		sharkBreed: sharkBreed,
		starve:     starve,
	}
	w.fillTheGrid()
	w.populateWorld(numFish, numShark)
	return w
}

// Fills the world with creature objects
func (w *World) fillTheGrid() {
	w.grid = make([][]*Creature, w.width)
	for i := range w.grid {
		w.grid[i] = make([]*Creature, w.height)
	}
	for i := 0; i < w.width; i++ {
		for j := 0; j < w.height; j++ {
			w.grid[i][j] = newCreatureEmpty(i, j)
		}
	}
}
func (w *World) spawnCreature(creatureId, x, y int) {
	var breedTime int
	if creatureId == 1 {
		breedTime = w.fishBreed
	} else if creatureId == 2 {
		breedTime = w.sharkBreed
	}
	creature := newCreature(
		creatureId, x, y,
		w.starve,
		breedTime)
	w.creatures = append(w.creatures, creature)
	w.grid[x][y] = creature
}

func (w *World) populateWorld(numFish, numShark int) {
	w.placeCreatures(numFish, FISH)
	w.placeCreatures(numShark, SHARK)
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
		{0, -1}, // North
		{1, 0},  // East
		{0, 1},  // South
		{-1, 0}, // West
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
func (w *World) evolveCreatures(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	var newX int
	var newY int
	neighbours := w.get_neighbours(creature.x, creature.y)
	creature.fertility += 1
	moved := false
	if creature.id == SHARK {
		if checkIfAnyNeighbourIsFood(neighbours) {
			mutex.Lock()
			neighbours := getFoodNeighbours(neighbours)
			neighbour := randomiseNeighbour(neighbours)
			newX = neighbour.x
			newY = neighbour.y
			creature.energy += 2
			w.grid[newX][newY].dead = true
			w.grid[newX][newY] = newCreatureEmpty(newX, newY)
			mutex.Unlock()
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
		mutex.Lock()
		w.grid[creature.x][creature.y].id = 0
		mutex.Unlock()
	} else if moved {
		x := creature.x
		y := creature.y
		creature.x = newX
		creature.y = newY
		mutex.Lock()
		w.grid[newX][newY] = creature
		if creature.fertility >= creature.breedTime {
			creature.fertility = 0
			w.spawnCreature(creature.id, x, y)
		} else {
			w.grid[x][y] = newCreatureEmpty(x, y)
		}
		mutex.Unlock()
	}
	<-semChannel
	wg.Done()
}
func (w *World) evolveWorld() {
	// Shuffles the creature slice
	rand.Shuffle(len(w.creatures),
		func(i, j int) { w.creatures[i], w.creatures[j] = w.creatures[j], w.creatures[i] })

	semChannel := make(chan bool, 16)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	ncreatures := len(w.creatures)
	for i := 0; i < ncreatures; i++ {
		creature := w.creatures[i]
		if creature.dead {
			continue
		}
		wg.Add(1)
		semChannel <- true
		go w.evolveCreatures(creature, semChannel, &mutex, &wg)
	}
	wg.Wait()
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
