package main

type Seacreature struct {
	species int
	moved   bool
	fish    *Fish
	shark   *Shark
}

func newSeacreatureFish(fish *Fish) *Seacreature {
	return &Seacreature{
		species: 1,
		moved:   false,
		fish:    fish,
		shark:   nil,
	}
}

func newSeacreatureShark(shark *Shark) *Seacreature {
	return &Seacreature{
		species: 2,
		moved:   false,
		fish:    nil,
		shark:   shark,
	}
}

func newSeacreatureEmpty() *Seacreature {
	return &Seacreature{
		species: 0,
		moved:   false,
		fish:    nil,
		shark:   nil,
	}
}
