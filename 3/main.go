package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Part 1
/*
func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	data := string(f)
	var l int
	for i := 0; i < len(data); i++ {
		if data[i] == byte(10) {
			l = i
			break
		}
	}

	g := make([]int, l)
	e := make([]int, l)
	half := (len(data) + 1) / (l + 1) / 2

	// 48 and 49 are ASCII 0 and 1, use some math to convert to 0 and 1 without string converstion
	for i := 0; i < len(data); i += l + 1 {
		for j := 0; j < l; j++ {
			g[j] += int(data[i : i+l][j] - 48)
			e[j] += int(data[i : i+l][j]-49) / 255
		}
	}

	log.Println(g)
	log.Println(e)

	// Want to not have to branch here for when the entire column is the same value
	for i := 0; i < len(g); i++ {
		if g[i] == len(data)+1 {
			g[i] = 1

		} else {
			g[i] /= half
		}
	}

	for i := 0; i < len(e); i++ {
		if e[i] == len(data)+1 {
			e[i] = 1

		} else {
			e[i] /= half
		}
	}

	log.Println(g)
	log.Println(e)

	gamma := 0
	epsilon := 0

	// Use a bit shift on the bit slice to get the represented number
	for i := 0; i < len(g); i++ {
		gamma += g[i] << (len(g) - i - 1)
	}

	for i := 0; i < len(e); i++ {
		epsilon += e[i] << (len(e) - i - 1)
	}

	log.Println(gamma)
	log.Println(epsilon)

	log.Println(gamma * epsilon)
}
*/

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	b := strings.Split(string(f), "\n")
	length := len(b[0])
	bits := make([][]int, len(b))
	for i := 0; i < len(b); i++ {
		bits[i] = make([]int, length)
		for j := 0; j < length; j++ {
			bit, err := strconv.Atoi(b[i][j : j+1])
			if err != nil {
				log.Fatal(err)
			}
			bits[i][j] = bit
		}
	}

	o2 := BitsToInt(Search(bits, 0, true))
	co2 := BitsToInt(Search(bits, 0, false))

	log.Println(o2)
	log.Println(co2)

	log.Println(o2 * co2)
}

func Search(b [][]int, p int, o2 bool) []int {
	if len(b) == 1 {
		log.Println(b[0])
		return b[0]
	} else {
		var bits [][]int
		sum := 0
		common := 0
		uncommon := 1
		for i := 0; i < len(b); i++ {
			sum += b[i][p]
		}
		if sum > len(b)/2 || len(b)%2 == 0 && sum >= len(b)/2 {
			common = 1
			uncommon = 0
		}

		for i := 0; i < len(b); i++ {
			v := common
			if !o2 {
				v = uncommon
			}
			if b[i][p] == v {
				bits = append(bits, b[i])
			}
		}
		return Search(bits, p+1, o2)
	}
}

func Sum(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func BitsToInt(b []int) int {
	i := 0
	for j := 0; j < len(b); j++ {
		i += b[j] << (len(b) - j - 1)
	}
	return i
}
