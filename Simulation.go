//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// None I hope
//--------------------------------------------
// Note: A lot of the comments were placed while I was learning more about the language

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

// Updates Logical side of the game
// (g *Game) -> receiver type works like object method in java so Game.Update()
func (g *Game) Update() error {
	//var fish []Fish

	return nil
}

// Draws the screen
func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, 10, 10, 320, 240, 10, color.White, true)
	ebitenutil.DebugPrint(screen, "Hello, Worsld!")
}

// The game has inner logical screen size which is set to 320x and 240y
// Resizing the window won't change this logic
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

////////////////////////////

type Grid struct {
	// Value 0 in position = empty
	// Value 1 in poistion = fish
	// Value 2 in position = shark
	locations [320][240]float32
}
type Position struct {
	xPosition float32
	yPosition float32
}
type Fish struct {
	position Position
}

type Shark struct {
	position Position
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(30)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
