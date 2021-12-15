package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var shortestPath [][]int
var shortestSum = 0

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	grid := make([][]int, len(input))
	path := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		grid[i] = make([]int, len(input[i]))
		path[i] = make([]int, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			n, err := strconv.Atoi(input[i][j : j+1])
			if err != nil {
				log.Fatal(err)
			}
			grid[i][j] = n
			path[i][j] = 0
		}
	}

	for i := 0; i < len(grid); i++ {
		shortestSum += grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		shortestSum += grid[len(grid)-1][i]
	}

	findPath(0, 0, path, grid)

	for i := 0; i < len(shortestPath); i++ {
		log.Println(shortestPath[i])
	}
}

func findPath(i, j int, path [][]int, grid [][]int) {
	sum := 0
	for y := 0; y < len(path); y++ {
		for x := 0; x < len(path[y]); x++ {
			sum += path[y][x]
		}
	}
	if sum >= shortestSum {
		return
	}

	if i == len(grid)-1 && j == len(grid[i])-1 {
		sum := 0
		for y := 0; y < len(path); y++ {
			for x := 0; x < len(path[y]); x++ {
				sum += path[y][x]
			}
		}
		if sum < shortestSum {
			shortestPath = path
			shortestSum = sum
		}
	}
	copy := make([][]int, len(path))
	for y := 0; y < len(path); y++ {
		copy[y] = make([]int, len(path[y]))
		for x := 0; x < len(path[y]); x++ {
			copy[y][x] = path[y][x]
		}
	}
	copy[i][j] = 1
	for n := 0; n < len(copy); n++ {
		log.Println(copy[n])
	}
	log.Println()

	if i-1 >= 0 && copy[i-1][j] == 0 {
		findPath(i-1, j, copy, grid)
	}
	if i+1 < len(grid) && copy[i+1][j] == 0 {
		findPath(i+1, j, copy, grid)
	}
	if j-1 >= 0 && copy[i][j-1] == 0 {
		findPath(i, j-1, copy, grid)
	}
	if j+1 < len(grid[i]) && copy[i][j+1] == 0 {
		findPath(i, j+1, copy, grid)
	}
}
