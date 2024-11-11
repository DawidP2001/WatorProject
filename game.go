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

	fishSlice  []Fish
	sharkSlice []Shark
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
	s := Seacreature{
		species: 1,
		fish:    *newFish(),
	}
	g.grid.locations[200][100] = s
	g.grid.locations[100][50] = s
	return g
}
func (g *Game) updateChronon() {
	g.chronon++
}

// Updates Logical side of the game
func (g *Game) Update() error {
	g.updateChronon()
	g.grid.resetMovedPositions()
	g.updateFishTest()
	//time.Sleep(200 * time.Millisecond)
	return nil
}
func (g *Game) updateFish() {
	for i := 0; i < len(g.grid.locations); i++ {
		for j := 0; j < len(g.grid.locations[i]); j++ {
			if g.grid.locations[i][j].species == 1 {
				if !g.grid.locations[i][j].moved {
					fish := g.grid.locations[i][j].fish
					fish.setNewPosition(g, i, j, 320, 240)
				}
			}
		}
	}
}
func (g *Game) updateFishTest() {
	for i := 319; i >= 0; i-- {
		for j := 239; j >= 0; j-- {
			if g.grid.locations[i][j].species == 1 {
				if !g.grid.locations[i][j].moved {
					fish := g.grid.locations[i][j].fish
					fish.setNewPosition(g, i, j, 320, 240)
				}
			}
		}
	}
}

// Draws the screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGrid(screen)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for i := 0; i < len(g.grid.locations); i++ {
		for j := 0; j < len(g.grid.locations[i]); j++ {
			if g.grid.locations[i][j].species == 1 {
				screen.Set(i, j, color.RGBA{R: 0, G: 255, B: 0, A: 255})
			} else if g.grid.locations[i][j].species == 2 {
				screen.Set(i, j, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			} else {
				screen.Set(i, j, color.RGBA{R: 0, G: 0, B: 255, A: 255})
			}

		}
	}
}

// The game has inner logical screen size which is set to 320x and 240y
// Resizing the window won't change this logic
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
