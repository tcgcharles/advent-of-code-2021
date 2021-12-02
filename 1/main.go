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

	measurements := strings.Split(string(f), "\n")

	var n int
	for i := 1; i < len(measurements); i++ {
		b, err := strconv.Atoi(measurements[i-1])
		if err != nil {
			log.Fatal(err)
		}

		a, err := strconv.Atoi(measurements[i])
		if err != nil {
			log.Fatal(err)
		}

		if a > b {
			n++
		}
	}
	log.Println(n)

	m := make([]int, len(measurements))
	for i := 0; i < len(measurements); i++ {
		n, err := strconv.Atoi(measurements[i])
		if err != nil {
			log.Fatal(err)
		}

		m[i] = n
	}
	//m = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	n = 0
	for i := 3; i < len(m); i++ {
		if m[i-2]+m[i-1]+m[i] > m[i-3]+m[i-2]+m[i-1] {
			n++
		}
	}
	log.Println(n)
}
