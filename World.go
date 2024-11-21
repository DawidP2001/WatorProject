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
* @brief Creates a new struct of type World
*
* This function is a constructor for the struct world, it takes several int variables and assigns it to this new instance.
*
* @param numShark 		Number of Sharks at the start of the game
* @param numFish		Number of Fish at the start of the game
* @param fishBread		Number of chronons that pass before a fish breeds
* @param sharkBread		Number of chronons that pass before a shark breeds
* @param starve			Number of chronons that pass before a shark starves
* @param gridSize		An int array of size 2, position 0 holds the width of the grid, position 1 holds the height
* @return A pointer towards a newly created struct of type World
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

/**
* @brief Fills the grid with empty Creature object pointers
*
* This function is used to initialise the grid for the game with creature object pointers that are used to represent empty water squares
 */
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
func (w *World) initSpawnCreature(creatureId, x, y int) {
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

/**
* @brief Spawns either a fish or a shark on the map
*
* This function is used to add a creature into a creature array as well adding it into the world grid
*
* @param creatureId 	Holds the value that represents whether its water, fish or shark
* @param x				X coordinate for the creature
* @param y				Y coordinate for the creature
 */
func (w *World) spawnCreature(creatureId, x, y int, mutex *sync.Mutex) {
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
	mutex.Lock()
	w.creatures = append(w.creatures, creature)
	w.grid[x][y] = creature
	mutex.Unlock()
}

/**
* @brief Populates the world with fish and sharks
*
* This function is called initially to call other functions which place the sharks and fish at random postion
*
* @param numFish 		Number of Fish at the start of the game
* @param numShark		Number of Shark at the start of the game
 */
func (w *World) populateWorld(numFish, numShark int) {
	w.placeCreatures(numFish, 1)
	w.placeCreatures(numShark, 2)
}

/**
* @brief Places creatures at a random Position on the grid
*
* Places a certain amount of fish or sharks on the grid. Their position is randomly chosen and assigned if it isn't already preoccupied
*
* @param ncreatures 	Number of creatures to place on the grid
* @param creatureId		The Id of the type of creature that is being placed, just used to pass a parameter to another function
 */
func (w *World) placeCreatures(ncreatures, creatureId int) {
	for i := 0; i < ncreatures; i++ {
		finish := false
		for !finish {
			randomX := rand.Intn(w.width)
			randomY := rand.Intn(w.height)
			if w.grid[randomX][randomY].id == 0 {
				w.initSpawnCreature(creatureId, randomX, randomY)
				finish = true
			}
		}
	}
}

//////////////////////////////////// DO THE ONE BELOW
/**
* @brief Gets neighbours of a given position
*
* 4 Direct neighbouring positions(North, West, South and East) are placed inside an array and returned.
* If a position reaches max or 0 it loops around to the other end of grid.
*
* @param x 		X coordinate of the creature
* @param y		Y coordinate of the creature
* @return 		Returns a pointer array of 4 of the direct neighbours of a given creature
 */
//////////////////////////////////
func (w *World) getAndBlockNeighbours(x, y int) (neighbours [4]*Creature) {
	for i, movement := range movements {
		directionX := movement[0]
		directionY := movement[1]
		newX := (x + directionX + w.width) % w.width
		newY := (y + directionY + w.height) % w.height
		neighbours[i] = w.grid[newX][newY]
		neighbours[i].usedChan <- true
	}
	return neighbours
}
func (w *World) moveCreatures(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	creature.usedChan <- true
	neighbours := w.getAndBlockNeighbours(creature.x, creature.y)
	mutex.Unlock()
	defer wg.Done()
	var newX int
	var newY int
	creature.fertility += 1
	moved := false

	if creature.id == 2 {
		if checkIfAnyNeighbourIsFood(neighbours, mutex) {
			neighbours := getFoodNeighbours(neighbours, mutex)
			neighbour := randomiseNeighbour(neighbours)
			creature.energy = w.starve
			mutex.Lock()
			newX = neighbour.x
			newY = neighbour.y
			w.grid[newX][newY].dead = true
			w.grid[newX][newY] = newCreatureEmpty(newX, newY)
			mutex.Unlock()
			moved = true
		}
	}
	if !moved {
		emptyNeighbours := getEmptyNeighbours(neighbours, mutex)
		if len(emptyNeighbours) != 0 {
			randomNeighbor := randomiseNeighbour(emptyNeighbours)
			mutex.Lock()
			newX = randomNeighbor.x
			newY = randomNeighbor.y
			mutex.Unlock()
			if creature.id == 2 {
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
		mutex.Unlock()
		if creature.fertility >= creature.breedTime {
			creature.fertility = 0
			w.spawnCreature(creature.id, x, y, mutex)
		} else {
			mutex.Lock()
			w.grid[x][y] = newCreatureEmpty(x, y)
			mutex.Unlock()
		}
	}
	<-semChannel
}

func (w *World) moveFish(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	creature.beingUsed = true
	var newX int
	var newY int
	neighbours := w.get_neighbours(creature.x, creature.y, mutex)
	emptyNeighbours := getEmptyNeighbours(neighbours, mutex)
	creature.fertility += 1

	if len(emptyNeighbours) != 0 {
		randomNeighbor := randomiseNeighbour(emptyNeighbours)
		mutex.Lock()
		newX = randomNeighbor.x
		newY = randomNeighbor.y
		mutex.Unlock()
		x := creature.x
		y := creature.y
		creature.x = newX
		creature.y = newY
		mutex.Lock()
		/*
			if w.grid[newX][newY].id == 0 && !w.grid[newX][newY].beingUsed {
				w.grid[newX][newY] = creature
				mutex.Unlock()
				if creature.fertility >= creature.breedTime {
					creature.fertility = 0
					w.spawnCreature(creature.id, x, y, mutex)
				} else {
					mutex.Lock()
					if w.grid[newX][newY].id == 0 && !w.grid[newX][newY].beingUsed {
						w.grid[x][y] = newCreatureEmpty(x, y)
					}
					mutex.Unlock()
				}
			} else {
				mutex.Unlock()
			}
		*/

		w.grid[newX][newY] = creature
		mutex.Unlock()
		if creature.fertility >= creature.breedTime {
			creature.fertility = 0
			w.spawnCreature(creature.id, x, y, mutex)
		} else {
			mutex.Lock()
			w.grid[x][y] = newCreatureEmpty(x, y)
			mutex.Unlock()
		}

	}
	<-semChannel
}
func (w *World) moveShark(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	creature.beingUsed = true
	moved := false
	var newX int
	var newY int
	neighbours := w.get_neighbours(creature.x, creature.y, mutex)
	creature.fertility += 1
	if checkIfAnyNeighbourIsFood(neighbours, mutex) {
		neighbours := getFoodNeighbours(neighbours, mutex)
		neighbour := randomiseNeighbour(neighbours)
		mutex.Lock()
		newX = neighbour.x
		newY = neighbour.y
		mutex.Unlock()
		creature.energy = w.starve
		mutex.Lock()
		//	if w.grid[newX][newY].id == 1 && !w.grid[newX][newY].beingUsed {
		w.grid[newX][newY].dead = true
		w.grid[newX][newY] = newCreatureEmpty(newX, newY)
		//	}
		mutex.Unlock()
		moved = true
	}
	if !moved {
		emptyNeighbours := getEmptyNeighbours(neighbours, mutex)
		if len(emptyNeighbours) != 0 {
			randomNeighbor := randomiseNeighbour(emptyNeighbours)
			newX = randomNeighbor.x
			newY = randomNeighbor.y
			if creature.id == 2 {
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
		mutex.Unlock()
		if creature.fertility >= creature.breedTime {
			creature.fertility = 0
			w.spawnCreature(creature.id, x, y, mutex)
		} else {
			w.grid[x][y] = newCreatureEmpty(x, y)
		}
	}
	<-semChannel
}

/**
* @brief Called to itterate the game by 1 chronon
*
* This is used to itterate the game state. First it shuffles the creature slice so different creatures would be moved first.
* Then it creates concurrency tools such as mutexes and channels. It then launches a goroutine for each of the creatures so it can move on the map.
* It then creates a new Creature slice and appends it with all the Creatures that are still alive. It then sets the creature slice with this new created slice.
*
 */
func (w *World) iterateProgram() {
	// Shuffles the creature slice
	rand.Shuffle(len(w.creatures),
		func(i, j int) { w.creatures[i], w.creatures[j] = w.creatures[j], w.creatures[i] })

	semChannel := make(chan bool, 16)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	ncreatures := len(w.creatures)
	for i := 0; i < ncreatures; i++ {
		mutex.Lock()
		creature := w.creatures[i]
		mutex.Unlock()
		if creature.dead {
			continue
		}
		wg.Add(1)
		semChannel <- true
		go w.moveCreatures(creature, semChannel, &mutex, &wg)
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

/**
* @brief Gets neighbours of a given position
*
* 4 Direct neighbouring positions(North, West, South and East) are placed inside an array and returned.
* If a position reaches max or 0 it loops around to the other end of grid.
*
* @param x 		X coordinate of the creature
* @param y		Y coordinate of the creature
* @return 		Returns a pointer array of 4 of the direct neighbours of a given creature
 */
func (w *World) get_neighbours(x, y int, mutex *sync.Mutex) [4]*Creature {
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
		mutex.Lock()
		neighbours[i] = w.grid[newX][newY]
		mutex.Unlock()
	}
	return neighbours
}

/**
* @brief Gets pointers to empty neighbours
*
* 4 Direct neighbouring positions(North, West, South and East) are passed into this function.
* The the ones that don't contain a fish or shark are returned
*
* @param neighbours 	An array with 4 Creature pointers in positions that are neighbouring a given creature.
* @return 				Returns a slice containing pointers to empty creatures from the neighbours array.
 */
func getEmptyNeighbours(neighbours [4]*Creature, mutex *sync.Mutex) []*Creature {
	var emptyNeighbours []*Creature
	for i := 0; i < 4; i++ {
		mutex.Lock()
		if neighbours[i].id == 0 {
			emptyNeighbours = append(emptyNeighbours, neighbours[i])
		}
		mutex.Unlock()
	}
	return emptyNeighbours
}

/**
* @brief Chooses a random position
*
* A slice containing between 1 and 4 Creature pointers is passed in.
* From that array 1 random neighbour is selected and returned.
*
* @param neighbours 	A slice containing Creature pointers containing available positions
* @return 				Returns a pointer to a Creature struct
 */
func randomiseNeighbour(neighbours []*Creature) *Creature {
	return neighbours[rand.Intn(len(neighbours))]
}

/**
* @brief Checks whether any of the neighbours have a fish inside.
*
* An array containing Creature pointers to the 4 neighbours of a given position are checked to see whether they have any fish present.
* If a fish is present true is returned. This function is only used by sharks.
*
* @param neighbours 	An Creature pointer array of all the 4 neighbours
* @return 				Returns a boolean
 */
func checkIfAnyNeighbourIsFood(neighbours [4]*Creature, mutex *sync.Mutex) bool {
	foodPresent := false
	for i := 0; i < 4; i++ {
		mutex.Lock()
		if neighbours[i].id == 1 {
			foodPresent = true
		}
		mutex.Unlock()
	}
	return foodPresent
}

/**
* @brief Gets neighbouring positions that contain fish
*
* An array containing Creature pointers to the 4 neighbours of a given position are checked to see whether they have any fish present.
* If a fish is present they are added to a slice. Then that fish slice is returned.
*
* @param neighbours 	An Creature pointer array of all the 4 neighbours
* @return 				Returns a Creature pointers slice of neighbouring fish.
 */
func getFoodNeighbours(neighbours [4]*Creature, mutex *sync.Mutex) []*Creature {
	var neighbour []*Creature
	for i := 0; i < 4; i++ {
		mutex.Lock()
		if neighbours[i].id == 1 {
			neighbour = append(neighbour, neighbours[i])
		}
		mutex.Unlock()
	}
	return neighbour
}
