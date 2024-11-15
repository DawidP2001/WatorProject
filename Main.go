//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// ToDO:
// 5. Add dimensionality
// 6. Add threads
// Mkae grid resizable
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
	ebiten.SetWindowTitle("Wa-Tor")
	ebiten.SetTPS(30)

	numShark := 100              //Starting population of sharks;
	numFish := 100               //Starting population of fish;
	fishBreed := 4               //Number of time units that pass before a fish can reproduce;
	sharkBreed := 20             //Number of time units that must pass before a shark can reproduce;
	starve := 15                 //Period of time a shark can go without food before dying;
	gridSize := [2]int{250, 250} //Dimensions of world;
	threads := 0                 //Number of threads to use.
	g := NewGame(numShark, numFish, fishBreed, sharkBreed, starve, gridSize, threads)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
