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
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

// Updates Logical side of the game
// (g *Game) -> receiver type works like object method in java so Game.Update()
func (g *Game) Update() error {
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

type Fish struct {
	xPosition float32
	yPosition float32
}

/**
 * @brief Generates Random Direction
 *
 * This function takes an array of available directions so 0-N, 1-E, 2-S, 3-W
 * randomly chooses one of them and returns one of these values as an int
 *
 * @param an array of available directions.
 * @return an random int from given array.
 */
func genRadomPosition(array []int) int {
	rand.NewSource(time.Now().UnixNano()) // Generates random seed
	randomAnswer := array[rand.Intn(len(array))]
	return randomAnswer
}

func checkAvailablePositions(postions float32) (checkAvailablePositions []float32) {

}

func fish() {

}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(30)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
