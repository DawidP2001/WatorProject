package Wator

import (
	"math/rand"
	"sync"
)

// This struct contains information about he grid(planet) the game takes place in
type World struct {
	Width      int           // Stores the width of the world
	Height     int           // Stores the height of the world
	Grid       [][]*Creature // Stores the grid the world takes place on, in each is a pointer to a creature object (includes empty tiles)
	Creatures  []*Creature   // Stores creatures that will be moved each itteration
	FishBreed  int           // Stores how often fish can breed
	SharkBreed int           // Stores how often sharks can breed
	Starve     int           // Stores after how long sharks starve
	Threads    int           // Stores number of threads used in this program
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
func NewWorld(numShark, numFish, fishBreed, sharkBreed, starve int, gridSize [2]int, threads int) *World {
	width := gridSize[0]
	height := gridSize[1]
	w := &World{
		Width:      width,
		Height:     height,
		FishBreed:  fishBreed,
		SharkBreed: sharkBreed,
		Starve:     starve,
		Threads:    threads,
	}
	w.FillTheGrid()
	w.PopulateWorld(numFish, numShark)
	return w
}

/**
* @brief Fills the grid with empty Creature object pointers
*
* This function is used to initialise the grid for the game with creature object pointers that are used to represent empty water squares
 */
func (w *World) FillTheGrid() {
	w.Grid = make([][]*Creature, w.Width)
	for i := range w.Grid {
		w.Grid[i] = make([]*Creature, w.Height)
	}
	for i := 0; i < w.Width; i++ {
		for j := 0; j < w.Height; j++ {
			w.Grid[i][j] = NewCreatureEmpty(i, j)
		}
	}
}
func (w *World) InitSpawnCreature(creatureId, x, y int) {
	var breedTime int
	if creatureId == 1 {
		breedTime = w.FishBreed
	} else if creatureId == 2 {
		breedTime = w.SharkBreed
	}
	creature := NewCreature(
		creatureId, x, y,
		w.Starve,
		breedTime)
	w.Creatures = append(w.Creatures, creature)
	w.Grid[x][y] = creature
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
func (w *World) SpawnCreature(creatureId, x, y int, mutex *sync.Mutex) {
	var breedTime int
	if creatureId == 1 {
		breedTime = w.FishBreed
	} else if creatureId == 2 {
		breedTime = w.SharkBreed
	}
	creature := NewCreature(
		creatureId, x, y,
		w.Starve,
		breedTime)
	mutex.Lock()
	w.Creatures = append(w.Creatures, creature)
	w.Grid[x][y] = creature
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
func (w *World) PopulateWorld(numFish, numShark int) {
	w.PlaceCreatures(numFish, 1)
	w.PlaceCreatures(numShark, 2)
}

/**
* @brief Places creatures at a random Position on the grid
*
* Places a certain amount of fish or sharks on the grid. Their position is randomly chosen and assigned if it isn't already preoccupied
*
* @param ncreatures 	Number of creatures to place on the grid
* @param creatureId		The Id of the type of creature that is being placed, just used to pass a parameter to another function
 */
func (w *World) PlaceCreatures(ncreatures, creatureId int) {
	for i := 0; i < ncreatures; i++ {
		finish := false
		for !finish {
			randomX := rand.Intn(w.Width)
			randomY := rand.Intn(w.Height)
			if w.Grid[randomX][randomY].Id == 0 {
				w.InitSpawnCreature(creatureId, randomX, randomY)
				finish = true
			}
		}
	}
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
func (w *World) GetNeighbours(x, y int) (neighbours [4]*Creature) {
	for i, movement := range Movements {
		directionX := movement[0]
		directionY := movement[1]
		newX := (x + directionX + w.Width) % w.Width
		newY := (y + directionY + w.Height) % w.Height

		neighbours[i] = w.Grid[newX][newY]
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
func GetFoodNeighbours(neighbours [4]*Creature) []*Creature {
	var neighbour []*Creature
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].Id == 1 {
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

func RandomiseNeighbour(neighbours []*Creature) (neighbour *Creature) {
	neighbour = nil // If no neighbour present thats not being used it will be set to nil
	start := rand.Intn(len(neighbours))
	// Below selects a free random neighbour
Loop:
	for i := 0; i < len(neighbours); i++ {
		select {
		case neighbours[start].UsedChan <- true:
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
func GetEmptyNeighbours(neighbours [4]*Creature) (emptyNeighbours []*Creature) {
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].Id == 0 {
			emptyNeighbours = append(emptyNeighbours, neighbours[i])
		}
	}
	return emptyNeighbours
}

/**
* @brief iterates all the creatures present on the creature slice
*
* This method is used to iterate creatures concurrently through the grid. It is both used by sharks and fish.
* Concurrency in this function is implemented using mainly buffered channels, mutex lock and a wait group.
*
* @param 	creature 	A pointer to a creature that needs to be iterated
* @param 	semChannel 	A buffered channel used as a way to implement threads
* @param 	mutex 		A mutex Lock, used only to be passed into spawning creatures function
* @param 	wg 			A weight group synchronisation tool
 */
func (w *World) IterateCreatures(creature *Creature, mutex *sync.Mutex, wg *sync.WaitGroup) {
	neighbours := w.GetNeighbours(creature.X, creature.Y)             // Gets 4 nearby grid neighbours
	if !creature.Dead && creature == w.Grid[creature.X][creature.Y] { // If creature dead or no longer on grid just skip this function
		select {
		case creature.UsedChan <- true: // If creature is free
			creature.Fertility += 1 // Plus 1 chronon towards breeding timer
			moved := false
			x := creature.X
			y := creature.Y
			var neighbour *Creature
			if creature.Id == 2 {
				foodNeighbours := GetFoodNeighbours(neighbours) // Checks if shark can eat
				if len(foodNeighbours) != 0 {
					neighbour = RandomiseNeighbour(foodNeighbours)
					if neighbour != nil {
						// Eat Fish below
						creature.Energy = w.Starve
						w.Grid[neighbour.X][neighbour.Y].Dead = true
						moved = true
					}
				}
			}
			if !moved {
				emptyNeighbours := GetEmptyNeighbours(neighbours)
				if len(emptyNeighbours) != 0 {
					neighbour = RandomiseNeighbour(emptyNeighbours)
					if creature.Id == 2 {
						creature.Energy--
					}
					if neighbour != nil {
						moved = true
					}
				}
			}
			// Checks if sharks going to starve
			if creature.Energy < 0 {
				creature.Dead = true
				w.Grid[x][y] = NewCreatureEmpty(x, y)
				if neighbour != nil {
					<-neighbour.UsedChan // Shark dies so free neighbour
				}
			} else if moved {
				creature.X = neighbour.X
				creature.Y = neighbour.Y
				w.Grid[neighbour.X][neighbour.Y] = creature
				// Check if creature can breed
				if creature.Fertility >= creature.BreedTime {
					creature.Fertility = 0
					w.SpawnCreature(creature.Id, x, y, mutex)
				} else {
					w.Grid[x][y] = NewCreatureEmpty(x, y)
				}
				<-neighbour.UsedChan
			}
			<-creature.UsedChan
		default:
		}
	}
	wg.Done()
	//<-semChannel
}

/**
* @brief Called to itterate the game by 1 chronon
*
* This is used to itterate the game state. First it shuffles the creature slice so different creatures would be moved first.
* Then it creates concurrency tools such as mutexes and channels. It then launches a goroutine for each of the creatures so it can move on the map.
* It then creates a new Creature slice and appends it with all the Creatures that are still alive. It then sets the creature slice with this new created slice.
*
 */
func (w *World) IterateProgram() {
	// Shuffles the creature slice
	rand.Shuffle(len(w.Creatures),
		func(i, j int) { w.Creatures[i], w.Creatures[j] = w.Creatures[j], w.Creatures[i] })

	var mutex sync.Mutex
	var wg sync.WaitGroup
	// Since creatures slice grows during execution I needed a way to keep track the max for this itteration
	ncreatures := len(w.Creatures)
	for i := 0; i < ncreatures; i++ {
		creature := w.Creatures[i]
		if creature.Dead {
			continue
		}
		wg.Add(1)
		go w.IterateCreatures(creature, &mutex, &wg)
	}
	wg.Wait()
	var newCreatures []*Creature
	for i := 0; i < len(w.Creatures); i++ {
		if !w.Creatures[i].Dead {
			newCreatures = append(newCreatures, w.Creatures[i])
		}
	}
	w.Creatures = newCreatures
}
