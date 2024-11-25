package main

import "time"

var movements = [4][2]int{
	{0, -1}, // North
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
}

var timeout = 1000 * time.Millisecond
