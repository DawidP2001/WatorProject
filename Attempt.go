package main

import (
	"math/rand"
	"sync"
	"time"
)

func (w *World) getNeighbours(x, y int) (neighbours [4]*Creature) {
	for i, movement := range movements {
		directionX := movement[0]
		directionY := movement[1]
		newX := (x + directionX + w.width) % w.width
		newY := (y + directionY + w.height) % w.height

		neighbours[i] = w.grid[newX][newY]
	}
	return neighbours
}

func getFoodNeighbours2(neighbours [4]*Creature) []*Creature {
	var neighbour []*Creature
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].id == 1 {
			neighbour = append(neighbour, neighbours[i])
		}
	}
	return neighbour
}

func randomiseNeighbour2(neighbours []*Creature) (neighbour *Creature) {
	neighbour = nil
	start := rand.Intn(len(neighbours))
Loop:
	for i := 0; i < len(neighbours); i++ {
		select {
		case neighbours[start].usedChan <- true:
			neighbour = neighbours[start]
			break Loop
		default:
			start = (start + 1) % len(neighbours)
		}
	}
	return neighbour

}
func getEmptyNeighbours2(neighbours [4]*Creature) (emptyNeighbours []*Creature) {
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i].id == 0 {
			emptyNeighbours = append(emptyNeighbours, neighbours[i])
		}
	}
	return emptyNeighbours
}

func (w *World) iterateCreatures(creature *Creature, semChannel chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	neighbours := w.getNeighbours(creature.x, creature.y)
	if !creature.dead && creature == w.grid[creature.x][creature.y] {
		select {
		case creature.usedChan <- true:
			creature.fertility += 1
			moved := false
			x := creature.x
			y := creature.y
			var neighbour *Creature
			if creature.id == 2 {
				foodNeighbours := getFoodNeighbours2(neighbours)
				if len(foodNeighbours) != 0 {
					neighbour = randomiseNeighbour2(foodNeighbours)
					if neighbour != nil {
						creature.energy = w.starve
						w.grid[neighbour.x][neighbour.y].dead = true
						moved = true
					}
				}
			}
			if !moved {
				emptyNeighbours := getEmptyNeighbours2(neighbours)
				if len(emptyNeighbours) != 0 {
					neighbour = randomiseNeighbour2(emptyNeighbours)
					if creature.id == 2 {
						creature.energy--
					}
					if neighbour != nil {
						moved = true
					}
				}
			}
			if creature.energy < 0 {
				creature.dead = true
				w.grid[x][y] = newCreatureEmpty(x, y)
				if neighbour != nil {
					<-neighbour.usedChan
				}
			} else if moved {
				creature.x = neighbour.x
				creature.y = neighbour.y
				w.grid[neighbour.x][neighbour.y] = creature
				if creature.fertility >= creature.breedTime {
					creature.fertility = 0
					w.spawnCreature(creature.id, x, y, mutex)
				} else {
					w.grid[x][y] = newCreatureEmpty(x, y)
				}
				<-neighbour.usedChan
			}
			<-creature.usedChan
		case <-time.After(timeout):
		}
	}
	wg.Done()
	<-semChannel
}
