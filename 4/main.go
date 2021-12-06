package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`\s+`)

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	d := strings.Split(input[0], ",")
	draws := make([]int, len(d))
	for i := 0; i < len(d); i++ {
		parsed, err := strconv.Atoi(d[i])
		if err != nil {
			log.Fatal(err)
		}
		draws[i] = parsed
	}

	log.Println(draws)
	input = input[1:]
	n := len(input) / 6
	boards := make([][][]int, n)
	for b := 0; b < n; b++ {
		board := input[(b*6)+1 : (b*6)+6]
		boards[b] = make([][]int, 5)
		for i := 0; i < 5; i++ {
			row := strings.Split(strings.Trim(re.ReplaceAllString(board[i], " "), " "), " ")
			boards[b][i] = make([]int, 5)
			for j := 0; j < 5; j++ {
				parsed, err := strconv.Atoi(row[j])
				if err != nil {
					log.Fatal(err)
				}
				boards[b][i][j] = parsed
			}
		}
	}

	b, turn := Play1(boards, draws)

	board := boards[b]
	draw := draws[turn-1]
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !Contains(draws[:turn], board[i][j]) {
				sum += board[i][j]
			}
		}
	}

	log.Println(b)
	log.Println(draw)
	log.Println(sum)
	log.Println(sum * draw)

	b, turn = Play2(boards, draws)
	board = boards[b]
	draw = draws[turn-1]
	sum = 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !Contains(draws[:turn], board[i][j]) {
				sum += board[i][j]
			}
		}
	}

	log.Println(b)
	log.Println(draw)
	log.Println(sum)
	log.Println(sum * draw)
}

func Play1(boards [][][]int, draws []int) (int, int) {
	for i := 0; i < len(draws); i++ {
		for j := 0; j < len(boards); j++ {
			win := CheckWin(boards[j], draws, i)
			if win {
				return j, i
			}
		}
	}
	return -1, -1
}

func Play2(boards [][][]int, draws []int) (int, int) {
	var wins []int
	for i := 0; i < len(draws); i++ {
		for j := 0; j < len(boards); j++ {
			if Contains(wins, j) {
				continue
			}
			win := CheckWin(boards[j], draws, i)
			if win {
				wins = append(wins, j)
				if len(wins) == len(boards) {
					return j, i
				}
			}
		}
	}
	return -1, -1
}

func CheckWin(board [][]int, draws []int, turn int) bool {
	for i := 0; i < 5; i++ {
		sumRow := 0
		sumCol := 0
		for j := 0; j < 5; j++ {
			if Contains(draws[:turn], board[i][j]) {
				sumRow++
			}
			if Contains(draws[:turn], board[j][i]) {
				sumCol++
			}
		}
		if sumRow == 5 || sumCol == 5 {
			return true
		}
	}
	return false
}

func Contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
