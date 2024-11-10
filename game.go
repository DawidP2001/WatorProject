package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid Grid
}

func NewGame() *Game {
	g := &Game{
		grid: Grid{},
	}
	s := Seacreature{
		species: 1,
		fish:    *newFish(),
	}
	g.grid.locations[100][200] = s
	return g
}

// Updates Logical side of the game
func (g *Game) Update() error {
	g.grid.resetMovedPositions()
	g.updateFish()
	time.Sleep(1 * time.Second)
	return nil
}
func (g *Game) updateFish() {
	for i := 0; i < len(g.grid.locations); i++ {
		for j := 0; j < len(g.grid.locations[i]); j++ {
			if g.grid.locations[i][j].species == 1 {
				g.grid.locations[i][j].moved = true
				xPosition, yPoistion := g.grid.locations[i][j].fish.fishNextPosition(i, j)
				g.grid.locations[xPosition][yPoistion] = g.grid.locations[i][j]
				g.grid.locations[i][j].species = 0

			}
		}
	}
}
func (grid *Grid) resetMovedPositions() {
	for i := 0; i < len(grid.locations); i++ {
		for j := 0; j < len(grid.locations[i]); j++ {
			if grid.locations[i][j].species == 1 {
				grid.locations[i][j].moved = false
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
