package main

// @brief 				A Creature on the planet
// @param 	id 			Stores what type of creature it is (0-none, 1-Fish, 2-Shark)
// @param 	x 			X coordinate
// @param	y 			Y coordinate
// @param 	energy 		How much energy the creature has left
// @param 	breedTime 	How many chronons need to pass before the creature will breed
// @param 	fertility 	Counts the amount of chronons have passed, used for reproduction
// @param 	dead 		Stores if the creature died
type Creature struct {
	id        int
	x         int
	y         int
	energy    int
	breedTime int
	fertility int
	dead      bool
}

/**
* @brief					 Creates a new struct of type Creature
* @param 	id 				Stores what type of creature it is (0-none, 1-Fish, 2-Shark)
* @param 	x 				X coordinate
* @param 	y 				Y coordinate
* @param 	initialEnergy 	How much energy should the creature start with (only really affects sharks)
* @param 	breedTime 		When this fertility level is reached the creature reproduces
* @return 					A pointer towards a newly created struct of type Creature
 */
func newCreature(id, x, y, initialEnergy, breedTime int) *Creature {
	return &Creature{
		id:        id,
		x:         x,
		y:         y,
		energy:    initialEnergy,
		breedTime: breedTime,
		fertility: 0,
		dead:      false,
	}
}

/**
* @brief Creates a new struct of type Creature that will be used as an empty location on the grid
* @param x X coordinate
* @param y Y coordinate
* @return A pointer towards a newly created struct of type Creature that is of type 0
 */
func newCreatureEmpty(x, y int) *Creature {
	return &Creature{
		id: 0,
		x:  x,
		y:  y,
	}
}
