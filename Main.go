// --------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A Wator solution using concurrency
// Issues:
// 1. Game can get bit laggy
// ToDO:
// Took some logic for this program from https://scipython.com/blog/wa-tor-world/
// Especially for the iteration side of the program
// --------------------------------------------
package main

import (
	"WatorProject/Wator"
	"log"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

/**
* @brief This is the main function of the Wa-Tor program
*
* This the main function of the program. It sets up the game title and FPS.
* Makes new instance of the game struct and passed several parameters to its constructor.
 */

// Main Function of the program
func main() {
	ebiten.SetWindowTitle("Wa-Tor")

	numShark := 500              //Starting population of sharks;
	numFish := 1000              //Starting population of fish;
	fishBreed := 10              //Number of time units that pass before a fish can reproduce;
	sharkBreed := 20             //Number of time units that must pass before a shark can reproduce;
	starve := 15                 //Period of time a shark can go without food before dying;
	gridSize := [2]int{250, 250} //Dimensions of world;
	threads := 12                //Number of threads to use.
	runtime.GOMAXPROCS(threads)
	g := Wator.NewGame(numShark, numFish, fishBreed, sharkBreed, starve, gridSize, threads)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
