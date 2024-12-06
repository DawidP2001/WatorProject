package Wator

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// This struct is used to run the world/game
type Game struct {
	World   *World // Stores pointer to the world the game takes place on
	Chronon int    // Stores each itteration counter
	//startTime time.Time used to test execution time
}

// @brief 				Creates a new struct of type Game
//
// This function is a constructor for the struct Game, it takes several ints variables as well as a int array of size 2 and assigns it to this new instance.
//
// @param numShark 		Number of Sharks at the start of the game
// @param numFish		Number of Fish at the start of the game
// @param fishBread		Number of chronons that pass before a fish breeds
// @param sharkBread		Number of chronons that pass before a shark breeds
// @param starve			Number of chronons that pass before a shark starves
// @param gridSize		An int array of size 2, position 0 holds the width of the grid, position 1 holds the height
// @param threads		Number of threads that will be used by the program
// @return 				A pointer towards a newly created struct of type Game
func NewGame(numShark, numFish, fishBreed, sharkBreed, starve int, gridSize [2]int, threads int) *Game {
	w := NewWorld(numShark, numFish, fishBreed, sharkBreed, starve, gridSize, threads)
	return &Game{
		World:   w,
		Chronon: 0,
		//startTime: time.Now(), used to test execution time
	}
}

// @brief This function updates the game at each frame
//
// This function updates the game at each frame and returns a nil error.
//
// @return 				Returns a nil error
func (g *Game) Update() error {
	g.Chronon++
	/* This commented out section was used to test execution time
	if g.chronon == 1000 {
		elapsed := time.Since(g.startTime)
		print(elapsed.String())
		os.Exit(0)
	}
	*/
	g.World.IterateProgram()
	return nil
}

// @brief 				Draws the games screen at each frame
//
// # This function draws the games image at each itteration through the code
//
// @param screen			A pointer to an ebiten.Image object which will be used to represent the screen the game takes place in
func (g *Game) Draw(screen *ebiten.Image) {
	ebiten.SetWindowSize(g.World.Width*2, g.World.Height*2)
	g.DrawGrid(screen)
}

// @brief 				Draws the grid on the screen
//
// This function draws the games grid on the main screen. Fills the screen with blue pixels. Sets fish pixels as green and shark pixels as red.
//
// @param screen			A pointer to an ebiten.Image object containing the map
func (g *Game) DrawGrid(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	for i := 0; i < g.World.Width; i++ {
		for j := 0; j < g.World.Height; j++ {
			if g.World.Grid[i][j] != nil {
				creature := g.World.Grid[i][j]
				if creature.Id == 1 {
					screen.Set(i, j, color.RGBA{0, 255, 0, 255})
				} else if creature.Id == 2 {
					screen.Set(i, j, color.RGBA{255, 0, 0, 255})
				}
			}
		}
	}
}

// @brief 				Creates a new struct of type Game
//
// This function is a constructor for the struct Game, it takes several ints variables as well as a int array of size 2 and assigns it to this new instance.
//
// @param outsideWidth 		The width of the displayed window the game is in
// @param outsideHeight		The height of the displayed window the game is in
// @return screenWidth		The actual width of the screen, in the games logic
// @return screenHeight		The actual height of the screen, in the games logic
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.World.Width, g.World.Height
}
