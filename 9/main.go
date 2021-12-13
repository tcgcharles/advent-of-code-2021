package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X     int
	Y     int
	H     int
	IsMin bool
	Basin map[string]int
}

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	grid := make([][]Point, len(input))
	for i := 0; i < len(input); i++ {
		grid[i] = make([]Point, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			n, err := strconv.Atoi(input[i][j : j+1])
			if err != nil {
				log.Fatal(err)
			}
			grid[i][j] = Point{H: n, X: j, Y: i}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			n := i - 1
			s := i + 1
			w := j - 1
			e := j + 1
			low := true
			if n >= 0 && grid[n][j].H <= grid[i][j].H {
				low = false
			}
			if low && s < len(grid) && grid[s][j].H <= grid[i][j].H {
				low = false
			}
			if low && w >= 0 && grid[i][w].H <= grid[i][j].H {
				low = false
			}
			if low && e < len(grid[i]) && grid[i][e].H <= grid[i][j].H {
				low = false
			}
			grid[i][j].IsMin = low
		}
	}

	risk := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].IsMin {
				risk += 1 + grid[i][j].H
			}
		}
	}

	log.Println(risk)

	var basins []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].IsMin {
				grid[i][j].Basin = map[string]int{}

				checkBasin(j, i, grid[i][j], grid)
				basins = append(basins, len(grid[i][j].Basin))
			}
		}
	}

	sort.Ints(basins)
	log.Println(basins)
	log.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}

func checkBasin(x, y int, p Point, grid [][]Point) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) || grid[y][x].H == 9 {
		return
	} else {
		k := fmt.Sprintf("%d,%d", y, x)
		if _, ok := p.Basin[k]; ok {
			return
		} else {
			p.Basin[k] = grid[y][x].H
			checkBasin(x-1, y, p, grid)
			checkBasin(x+1, y, p, grid)
			checkBasin(x, y-1, p, grid)
			checkBasin(x, y+1, p, grid)
		}
	}
}
