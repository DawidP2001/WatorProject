package main

type World struct {
	width     int
	height    int
	ncells    int
	grid      [100][100]*Creature
	creatures []*Creature
	nfish     int
	nsharks   int
}

func newWorld(width, height int) *World {
	return &World{
		width:  width,
		height: height,
		ncells: width * height,
		//	grid:   [2][2]int,
		// Declare Creatures ???
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
}

func (w *World) placeCreatures(ncreatures, creatureId int) {
	for i := 0; i < ncreatures; i++ {
		for true {
			x :=
			y := w.height
			break
		}
	}
}
