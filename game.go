package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid       Grid
	chronon    int
	numShark   int     // Starting population of sharks
	numFish    int     // Starting population of fish
	fishBreed  int     // Number of time units that pass before a fish can reproduce
	sharkBreed int     // Number of time units that must pass before a shark can reproduce;
	starve     int     // Period of time a shark can go without food before dying;
	gridSize   [][]int // Dimensions of world

	fishSlice  []*Fish
	sharkSlice []*Shark
}

func NewGame(numShark int, numFish int, fishBreed int, sharkBreed int, starve int, gridSize [][]int) *Game {
	g := &Game{
		grid:       Grid{},
		chronon:    0,
		numShark:   numShark,
		numFish:    numFish,
		fishBreed:  fishBreed,
		sharkBreed: sharkBreed,
		starve:     starve,
		gridSize:   gridSize,
	}
	fish := newFish(g, 200, 100)
	g.grid.locations[200][100] = *newSeacreatureFish(fish)
	return g
}
func (g *Game) updateChronon() {
	g.chronon++
}

// Updates Logical side of the game
func (g *Game) Update() error {
	g.updateChronon()
	g.updateFish(320, 240)
	return nil
}
func (g *Game) updateFish(maxX int, maxY int) {
	if len(g.fishSlice) > 0 {
		for i := len(g.fishSlice) - 1; i >= 0; i-- {
			currentFish := g.fishSlice[i]
			currentFish.setNewPosition(g, maxX, maxY)
		}
	}
}

// Draws the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGrid(screen)
	g.drawFish(screen)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
}
func (g *Game) drawFish(screen *ebiten.Image) {
	for i := 0; i < len(g.fishSlice); i++ {
		fish := g.fishSlice[i]
		screen.Set(fish.x, fish.y, color.RGBA{0, 255, 0, 255})
	}
}

// The game has inner logical screen size which is set to 320x and 240y
// Resizing the window won't change this logic
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
