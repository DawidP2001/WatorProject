package main

type Grid struct {
	// Value 0 in position = empty
	// Value 1 in poistion = fish
	// Value 2 in position = shark
	locations [320][240]Seacreature
}
