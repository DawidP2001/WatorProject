package main

type Fish struct {
}

func newFish() *Fish {
	return &Fish{}
}
func (f *Fish) checkAvailablePositions(grid *Grid, xPosition int, yPosition int) (availablePositions []Position) {

	// Checks North

	// Checks East

	// Checks South

	// Cheks West

	return availablePositions
}

func (f *Fish) fishNextPosition(xPosition int, yPoistion int) (int, int) {
	//availablePositions := checkAvailablePositions()
	//newPosition := genRadomPosition(availablePositions)
	if xPosition < 319 {
		if yPoistion > 0 {
			return xPosition, yPoistion - 1
		}
	}

	return xPosition, yPoistion
}
