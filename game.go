package main

import (
	"image/color"

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
	}
	g.grid.locations[100][100] = s
	return g
}

// Updates Logical side of the game
func (g *Game) Update() error {

	return nil
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
