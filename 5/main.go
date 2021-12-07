package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	lines := make([][]Point, len(input))

	for i := 0; i < len(input); i++ {
		points := strings.Split(input[i], " -> ")
		lines[i] = make([]Point, len(points))

		for j := 0; j < len(points); j++ {
			point := strings.Split(points[j], ",")

			x, err := strconv.Atoi(point[0])
			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(point[1])
			if err != nil {
				log.Fatal(err)
			}

			lines[i][j] = Point{
				X: x,
				Y: y,
			}
		}
	}

	grid := make([][]int, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1000)
	}

	for i := 0; i < len(lines); i++ {
		p1 := lines[i][0]
		p2 := lines[i][1]
		if p1.X != p2.X && p1.Y != p2.Y {
			continue
		}
		x := p1.X
		y := p1.Y
		grid[x][y]++
		for x != p2.X || y != p2.Y {
			if x < p2.X {
				x++
			} else if x > p2.X {
				x--
			}

			if y < p2.Y {
				y++
			} else if y > p2.Y {
				y--
			}

			grid[x][y]++
		}
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				sum++
			}
		}
	}

	log.Println(sum)

	grid = make([][]int, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1000)
	}

	for i := 0; i < len(lines); i++ {
		p1 := lines[i][0]
		p2 := lines[i][1]
		x := p1.X
		y := p1.Y
		grid[x][y]++
		for x != p2.X || y != p2.Y {
			if x < p2.X {
				x++
			} else if x > p2.X {
				x--
			}

			if y < p2.Y {
				y++
			} else if y > p2.Y {
				y--
			}

			grid[x][y]++
		}
	}

	sum = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				sum++
			}
		}
	}

	log.Println(sum)
}
