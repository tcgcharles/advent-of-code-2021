package main

import (
	"fmt"
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
	pairs := map[string]int64{}
	for k := range poly {
		if k+2 <= len(poly) {
			pairs[poly[k:k+2]] += 1
		}
	}
	log.Println(poly)
	log.Println(pairs)
	input = input[2:len(input)]
	m := map[string]string{}
	for _, v := range input {
		line := strings.Split(v, " -> ")
		m[line[0]] = line[1]
	}

	for i := 0; i < 40; i++ {
		log.Println(i)

		inserts := map[string]int64{}
		for k, v := range pairs {
			insert := m[k]
			p := strings.Split(k, "")
			inserts[p[0]+p[1]] -= v
			inserts[p[0]+insert] += v
			inserts[insert+p[1]] += v
		}
		for k, v := range inserts {
			pairs[k] += v
		}
		/*
			var insert []string
			for j := 0; j < len(poly)-1; j++ {
				s := poly[j : j+2]
				insert = append(insert, m[s])
			}
			for k, v := range insert {
				s0 := poly[:k*2+1]
				l := k*2 + 1
				s1 := ""
				if l <= len(poly) {
					s1 = poly[k*2+1:]
				}
				poly = s0 + v + s1
			}
			log.Println(poly)
		*/
		log.Println(pairs)
		for k, v := range pairs {
			if v > 0 {
				log.Println(fmt.Sprintf("%s:%d %d %v", k, v, count(poly, k), v == count(poly, k)))
			}
		}
	}

	counts := map[string]int64{
		"B": 0,
		"C": 0,
		"H": 0,
		"N": 0,
	}

	for i := 0; i < len(poly); i++ {
		counts[poly[i:i+1]]++
	}

	log.Println(counts)

	c := map[string]int64{
		"B": 0,
		"C": 0,
		"H": 0,
		"N": 0,
	}

	for k, v := range pairs {
		p := strings.Split(k, "")
		c[p[0]] += v
		c[p[1]] += v
	}

	log.Println(c)

	var max int64 = 0
	var min int64 = math.MaxInt64

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

	max = 0
	min = math.MaxInt64

	for _, v := range c {
		c := v
		if c%2 != 0 {
			c++
		}
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	log.Println(max)
	log.Println(min)
	log.Println(max - min)

	log.Println(max / 2)
	log.Println(min / 2)
	log.Println((max - min) / 2)
}

func count(s, substr string) int64 {
	var c int64 = 0
	for k := range s {
		if k < len(s)-1 && s[k:k+2] == substr {
			c++
		}
	}
	return c
}
