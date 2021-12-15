package main

import (
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	poly := input[0]
	input = input[2:len(input)]
	m := map[string]string{}
	for _, v := range input {
		line := strings.Split(v, " -> ")
		m[line[0]] = line[1]
	}

	for i := 0; i < 40; i++ {
		log.Println(i)
		var inserts []string
		for j := 0; j < len(poly)-1; j++ {
			s := poly[j : j+2]
			inserts = append(inserts, m[s])
		}
		for k, v := range inserts {
			s0 := poly[:k*2+1]
			l := k*2 + 1
			s1 := ""
			if l <= len(poly) {
				s1 = poly[k*2+1:]
			}
			poly = s0 + v + s1
		}
	}

	counts := map[string]int{
		"B": 0,
		"C": 0,
		"H": 0,
		"N": 0,
	}

	for i := 0; i < len(poly); i++ {
		counts[poly[i:i+1]]++
	}

	log.Println(counts)

	max := 0
	min := math.MaxInt32

	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	log.Println(max)
	log.Println(min)
	log.Println(max - min)

}
