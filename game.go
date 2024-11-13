package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *World
}

func NewGame(numShark, numFish, fishBreed, sharkBreed, starve int, gridSize [2]int, threads int) *Game {
	w := newWorld(numShark, numFish, fishBreed, sharkBreed, starve, gridSize)
	return &Game{
		world: w,
	}
}

// Updates Logical side of the game
func (g *Game) Update() error {
	g.world.evolveWorld()
	return nil
}

// Draws the screen
func (g *Game) Draw(screen *ebiten.Image) {
	ebiten.SetWindowSize(g.world.width*2, g.world.height*2)
	g.drawGrid(screen)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	for i := 0; i < g.world.width; i++ {
		for j := 0; j < g.world.height; j++ {
			if g.world.grid[i][j] != nil {
				creature := g.world.grid[i][j]
				if creature.id == FISH {
					screen.Set(i, j, color.RGBA{0, 255, 0, 255})
				} else if creature.id == SHARK {
					screen.Set(i, j, color.RGBA{255, 0, 0, 255})
				}
			}
		}
	}
}

// The game has inner logical screen size which is set to 320x and 240y
// Resizing the window won't change this logic
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 250, 250
}
