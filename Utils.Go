package main

import (
	"math/rand"
	"time"
)

type Position struct {
	xPosition int
	yPosition int
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
	randomNum := rand.Intn(len(array))
	randomPosition = array[randomNum]
	return randomPosition
}
