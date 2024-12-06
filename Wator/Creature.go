package Wator

// A Creature on the planet
type Creature struct {
	Id        int       // Holds the value that identifies the species of the creature
	X         int       // Holds the X Coordinate
	Y         int       // Holds the Y coordinate
	Energy    int       // If its a shark stores the amount of chronons it can move before it dies
	BreedTime int       // Stores the time it takes until itll breed
	Fertility int       // Stores the counter until it can breed
	Dead      bool      // Stores whether it just died
	UsedChan  chan bool // Stores a channel used to stop other threads from accessing it
}

// @brief					Creates a new struct of type Creature
//
// This function is a constructor for the struct Creature, it takes several int variables and assigns it to this new instance.
//
// @param 	id 				Stores what type of creature it is (0-none, 1-Fish, 2-Shark)
//
// @param 	x 				X coordinate
// @param 	y 				Y coordinate
// @param 	initialEnergy 	How much energy should the creature start with (only really affects sharks)
// @param 	breedTime 		When this fertility level is reached the creature reproduces
// @return 					A pointer towards a newly created struct of type Creature
func NewCreature(id, x, y, initialEnergy, breedTime int) *Creature {
	return &Creature{
		Id:        id,
		X:         x,
		Y:         y,
		Energy:    initialEnergy,
		BreedTime: breedTime,
		Fertility: 0,
		Dead:      false,
		UsedChan:  make(chan bool, 1),
	}
}

//@brief Creates a new struct of type Creature that will be used as an empty location on the grid
//
//This function is a constructor for the struct Creature, it takes several int variables and assigns it to this new instance.
//
//@param x X coordinate
//@param y Y coordinate
//@return A pointer towards a newly created struct of type Creature that is of type 0
func NewCreatureEmpty(x, y int) *Creature {
	return &Creature{
		Id:       0,
		X:        x,
		Y:        y,
		UsedChan: make(chan bool, 1),
	}
}
