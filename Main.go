// --------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// 1. Game sometimes crashes/Freezes
// 2. Fix concurrency issue
// ToDO:
// 6. Add threads
// --------------------------------------------
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

/**
* @brief This is the main function of the Wa-Tor program
 */
func main() {
	ebiten.SetWindowTitle("Wa-Tor")
	ebiten.SetTPS(100)

	numShark := 100              //Starting population of sharks;
	numFish := 100               //Starting population of fish;
	fishBreed := 2               //Number of time units that pass before a fish can reproduce;
	sharkBreed := 20             //Number of time units that must pass before a shark can reproduce;
	starve := 15                 //Period of time a shark can go without food before dying;
	gridSize := [2]int{250, 250} //Dimensions of world;
	threads := 0                 //Number of threads to use.
	g := NewGame(numShark, numFish, fishBreed, sharkBreed, starve, gridSize, threads)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
