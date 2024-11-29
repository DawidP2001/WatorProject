package Wator

// @brief 				A Creature on the planet
type Creature struct {
	id        int
	x         int
	y         int
	energy    int
	breedTime int
	fertility int
	dead      bool
	usedChan  chan bool
}

/**
* @brief					Creates a new struct of type Creature
*
* This function is a constructor for the struct Creature, it takes several int variables and assigns it to this new instance.
*
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
		usedChan:  make(chan bool, 1),
	}
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
func newCreatureEmpty(x, y int) *Creature {
	return &Creature{
		id:       0,
		x:        x,
		y:        y,
		usedChan: make(chan bool, 1),
	}
}
