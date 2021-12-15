package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	var points [][]int
	var folds []map[string]int
	x := 0
	y := 0
	for _, v := range input {
		if len(v) == 0 {
			continue
		}

		if strings.Contains(v, ",") {
			point := strings.Split(v, ",")
			x, err := strconv.Atoi(point[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(point[1])
			if err != nil {
				log.Fatal(err)
			}
			points = append(points, []int{x, y})
		}

		if strings.Contains(v, "=") {
			line := strings.TrimLeft(v, "fold along ")
			fold := strings.Split(line, "=")
			f, err := strconv.Atoi(fold[1])
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, map[string]int{fold[0]: f})
		}
	}

	for _, v := range points {
		if v[0] > x {
			x = v[0]
		}
		if v[1] > y {
			y = v[1]
		}
	}

	grid := make([][]int, y+1)
	for i := 0; i <= y; i++ {
		grid[i] = make([]int, x+1)
	}

	for _, v := range points {
		grid[v[1]][v[0]] = 1
	}

	for _, v := range folds {
		i, ok := v["x"]
		if ok {
			grid = foldLeft(i, grid)
		} else {
			i, ok = v["y"]
			if ok {
				grid = foldUp(i, grid)
			}
		}

		sum := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] > 0 {
					sum++
				}
			}
		}

		log.Println(sum)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		s := ""
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				s += " "
			} else {
				s += "#"
			}
		}
		log.Println(s)
	}
}

func foldUp(y int, grid [][]int) [][]int {
	fold := grid[y+1:]
	g := grid[:y]

	for i := 0; i < len(fold); i++ {
		for j := 0; j < len(fold[i]); j++ {
			g[i][j] += fold[len(fold)-1-i][j]
		}
	}

	return g
}

func foldLeft(x int, grid [][]int) [][]int {
	g := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		g[i] = make([]int, x)
		for j := 0; j < x; j++ {
			g[i][j] = grid[i][j]
			g[i][j] += grid[i][len(grid[i])-1-j]
		}
	}

	return g
}
