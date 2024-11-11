//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// 1. Fish Priorities going north - found in updateFish function
// ToDO:
// 1. Fish Reproduction
// 2. add Sharks
// 3. Starvation
// 4. shark Reproduction
// 5. Add dimensionality
// 6. Add threads
//--------------------------------------------
// Note: A lot of the comments were placed while I was learning more about the language

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(30)
	g := NewGame(0, 0, 10, 0, 0, nil)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
