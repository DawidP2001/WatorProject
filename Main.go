//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// 1. Not sure if the gaps in the water is a glitch
// 2. Change the grid.locations to a pointer 2d array
// 4. Sharks stop swimming after a while
// 5. Fix the issue with slices -> concurrency problem
// Add More semaphores
// Make checkAvailablePositions into a util method for both sharks and fish
//  -> no need to check south west north and so on if its 1 function
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
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(30)
	numSharks := 0   // Set this to the amount of starting sharks
	numFish := 0     // Set this to the amount of starting fish
	fishBreed := 20  // Set this to how many chronons should pass before your fish breed
	sharkBreed := 40 // Set this to how many chronons should pass before your sharks breed
	starve := 100    // Set this to how much time can pass before your shark starves
	//	gridSize := []int{320, 240} // Set this to the size of the map
	g := NewGame(numSharks, numFish, fishBreed, sharkBreed, starve, nil)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
