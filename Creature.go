package main

type Creature struct {
	id        int
	x         int
	y         int
	energy    int
	fishBreed int
	fertility int
	dead      bool
}

func newCreature(id, x, y, initialEnergy, fishBreed int) *Creature {
	return &Creature{
		id:        id,
		x:         x,
		y:         y,
		energy:    initialEnergy,
		fishBreed: fishBreed,
		fertility: 0,
		dead:      false,
	}
}
