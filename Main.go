//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// 1. Not sure if the gaps in the water is a glitch
// 2. Change the grid.locations to a pointer 2d array
// 3. Fix starvation mechanic
// 4. Sharks stop swimming after a while
// 5. Fix the issue with slices
// ToDO:
// 5. Add dimensionality
// 6. Add threads
//--------------------------------------------
// Note: A lot of the comments were placed while I was learning more about the language

// Idea to fix the issue in this branch
// 1. Create fish array and itterate through that array to update positions instead of the nested for loop

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Wa-Tor")
	ebiten.SetTPS(30)

	numShark := 100              //Starting population of sharks;
	numFish := 100               //Starting population of fish;
	fishBreed := 5               //Number of time units that pass before a fish can reproduce;
	sharkBreed := 25             //Number of time units that must pass before a shark can reproduce;
	starve := 20                 //Period of time a shark can go without food before dying;
	gridSize := [2]int{250, 250} //Dimensions of world;
	threads := 0                 //Number of threads to use.

	g := NewGame(numShark, numFish, fishBreed, sharkBreed, starve, gridSize, threads)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
