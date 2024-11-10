package main

import (
	"math/rand"
	"time"
)

type Seacreature struct {
	species int
	moved   bool
	fish    Fish
	shark   Shark
}
type Grid struct {
	// Value 0 in position = empty
	// Value 1 in poistion = fish
	// Value 2 in position = shark
	locations [320][240]Seacreature
}
type Position struct {
	xPosition float32
	yPosition float32
}

/**
 * @brief Generates Random Direction
 *
 * This function takes an array of available positions and it
 * randomly chooses one of them and returns one of these values
 *
 * @param array An array of available directions.
 * @return an random position from given array.
 */
func genRadomPosition(array []Position) (randomPosition Position) {
	rand.NewSource(time.Now().UnixNano()) // Generates random seed
	randomPosition = array[rand.Intn(len(array))]
	return randomPosition
}
