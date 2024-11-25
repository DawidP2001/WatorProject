package main

import (
	"math/rand"
	"sync"
	"time"
)

type World struct {
	width      int
	height     int
	grid       [][]*Creature
	creatures  []*Creature
	fishBreed  int
	sharkBreed int
	starve     int
	threads    int
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
func newWorld(numShark, numFish, fishBreed, sharkBreed, starve int, gridSize [2]int, threads int) *World {
	width := gridSize[0]
	height := gridSize[1]
	w := &World{
		width:      width,
		height:     height,
		fishBreed:  fishBreed,
		sharkBreed: sharkBreed,
		starve:     starve,
		threads:    threads,
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
func (w *World) getNeighbours(x, y int) (neighbours [4]*Creature) {
	for i, movement := range movements {
		directionX := movement[0]
		directionY := movement[1]
		newX := (x + directionX + w.width) % w.width
		newY := (y + directionY + w.height) % w.height

		neighbours[i] = w.grid[newX][newY]
	}
	return neighbours
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
func getFoodNeighbours(neighbours [4]*Creature) []*Creature {
	var neighbour []*Creature
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].id == 1 {
			neighbour = append(neighbour, neighbours[i])
		}
	}
	return neighbour
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

func randomiseNeighbour(neighbours []*Creature) (neighbour *Creature) {
	neighbour = nil // If no neighbour present thats not being used it will be set to nil
	start := rand.Intn(len(neighbours))
	// Below selects a free random neighbour
Loop:
	for i := 0; i < len(neighbours); i++ {
		select {
		case neighbours[start].usedChan <- true:
			neighbour = neighbours[start]
			break Loop
		default:
			start = (start + 1) % len(neighbours)
		}
	}
	return neighbour

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
func getEmptyNeighbours(neighbours [4]*Creature) (emptyNeighbours []*Creature) {
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].id == 0 {
			emptyNeighbours = append(emptyNeighbours, neighbours[i])
		}
	}
	return emptyNeighbours
}

/**
* @brief iterates all the creatures present on the creature slice
*
* 4 Direct neighbouring positions(North, West, South and East) are passed into this function.
* The the ones that don't contain a fish or shark are returned
*
* @param neighbours 	An array with 4 Creature pointers in positions that are neighbouring a given creature.
 */
func (w *World) iterateCreatures(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	neighbours := w.getNeighbours(creature.x, creature.y)             // Gets 4 nearby grid neighbours
	if !creature.dead && creature == w.grid[creature.x][creature.y] { // If creature dead or no longer on grid just skip this function
		select {
		case creature.usedChan <- true: // If creature is free
			creature.fertility += 1 // Plus 1 chronon towards breeding timer
			moved := false
			x := creature.x
			y := creature.y
			var neighbour *Creature
			if creature.id == 2 {
				foodNeighbours := getFoodNeighbours(neighbours) // Checks if shark can eat
				if len(foodNeighbours) != 0 {
					neighbour = randomiseNeighbour(foodNeighbours)
					if neighbour != nil {
						// Eat Fish below
						creature.energy = w.starve
						w.grid[neighbour.x][neighbour.y].dead = true
						moved = true
					}
				}
			}
			if !moved {
				emptyNeighbours := getEmptyNeighbours(neighbours)
				if len(emptyNeighbours) != 0 {
					neighbour = randomiseNeighbour(emptyNeighbours)
					if creature.id == 2 {
						creature.energy--
					}
					if neighbour != nil {
						moved = true
					}
				}
			}
			// Checks if sharks going to starve
			if creature.energy < 0 {
				creature.dead = true
				w.grid[x][y] = newCreatureEmpty(x, y)
				if neighbour != nil {
					<-neighbour.usedChan // Shark dies so free neighbour
				}
			} else if moved {
				creature.x = neighbour.x
				creature.y = neighbour.y
				w.grid[neighbour.x][neighbour.y] = creature
				// Check if creature can breed
				if creature.fertility >= creature.breedTime {
					creature.fertility = 0
					w.spawnCreature(creature.id, x, y, mutex)
				} else {
					w.grid[x][y] = newCreatureEmpty(x, y)
				}
				<-neighbour.usedChan
			}
			<-creature.usedChan
		case <-time.After(timeout):
		}
	}
	wg.Done()
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

	semChannel := make(chan bool, w.threads)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	// Since creatures slice grows during execution I needed a way to keep track the max for this itteration
	ncreatures := len(w.creatures)
	for i := 0; i < ncreatures; i++ {
		creature := w.creatures[i]
		if creature.dead {
			continue
		}
		wg.Add(1)
		semChannel <- true
		go w.iterateCreatures(creature, semChannel, &mutex, &wg)
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
