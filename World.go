package main

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
	w.populateWorld(nfish, nsharks)
	return w
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
	//w.placeCreatures(w.nsharks, SHARK)
}

func (w *World) placeCreatures(ncreatures, creatureId int) {
	w.spawnCreature(creatureId, 25, 25)
}

func get_neighbours() {

}
func evolveCreatures() {

}
func evolveWorld() {

}
