package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Octopus struct {
	X       int
	Y       int
	Flashed bool
	Energy  int
}

var flashes int

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	grid := make([][]Octopus, len(input))
	for i := 0; i < len(input); i++ {
		grid[i] = make([]Octopus, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			e, err := strconv.Atoi(input[i][j : j+1])
			if err != nil {
				log.Fatal(err)
			}
			grid[i][j] = Octopus{
				X:       j,
				Y:       i,
				Energy:  e,
				Flashed: false,
			}
		}
	}

	flashes = 0
	for step := 0; step < 1000; step++ {

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j].Energy++
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j].Energy > 9 && !grid[i][j].Flashed {
					log.Println(fmt.Sprintf("%d,%d", j, i))
					flash(j, i, grid)
				}
			}
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j].Flashed {
					grid[i][j].Energy = 0
					grid[i][j].Flashed = false
				}
			}
		}

		sum := 0
		for i := 0; i < len(grid); i++ {
			log.Println(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d", grid[i][0].Energy, grid[i][1].Energy, grid[i][2].Energy, grid[i][3].Energy, grid[i][4].Energy, grid[i][5].Energy, grid[i][6].Energy, grid[i][7].Energy, grid[i][8].Energy, grid[i][9].Energy))
			for j := 0; j < len(grid[i]); j++ {
				sum += grid[i][j].Energy
			}
		}
		if sum == 0 {
			log.Println(step + 1)
			break
		}
	}

	log.Println(flashes)
	flashes = 0
}

func flash(x, y int, grid [][]Octopus) {
	flashes++
	grid[y][x].Flashed = true
	if y-1 >= 0 {
		grid[y-1][x].Energy++
		if grid[y-1][x].Energy > 9 && !grid[y-1][x].Flashed {
			flash(x, y-1, grid)
		}
	}
	if y+1 < len(grid) {
		grid[y+1][x].Energy++
		if grid[y+1][x].Energy > 9 && !grid[y+1][x].Flashed {
			flash(x, y+1, grid)
		}
	}
	if x-1 >= 0 {
		grid[y][x-1].Energy++
		if grid[y][x-1].Energy > 9 && !grid[y][x-1].Flashed {
			flash(x-1, y, grid)
		}
	}
	if x+1 < len(grid[y]) {
		grid[y][x+1].Energy++
		if grid[y][x+1].Energy > 9 && !grid[y][x+1].Flashed {
			flash(x+1, y, grid)
		}
	}
	if y-1 >= 0 && x-1 >= 0 {
		grid[y-1][x-1].Energy++
		if grid[y-1][x-1].Energy > 9 && !grid[y-1][x-1].Flashed {
			flash(x-1, y-1, grid)
		}
	}
	if y-1 >= 0 && x+1 < len(grid[y-1]) {
		grid[y-1][x+1].Energy++
		if grid[y-1][x+1].Energy > 9 && !grid[y-1][x+1].Flashed {
			flash(x+1, y-1, grid)
		}
	}
	if y+1 < len(grid) && x-1 >= 0 {
		grid[y+1][x-1].Energy++
		if grid[y+1][x-1].Energy > 9 && !grid[y+1][x-1].Flashed {
			flash(x-1, y+1, grid)
		}
	}
	if y+1 < len(grid) && x+1 < len(grid[y+1]) {
		grid[y+1][x+1].Energy++
		if grid[y+1][x+1].Energy > 9 && !grid[y+1][x+1].Flashed {
			flash(x+1, y+1, grid)
		}
	}
}
