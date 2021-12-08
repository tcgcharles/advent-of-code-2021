package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	sums := make([]int, 10)
	lines := strings.Split(string(f), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], "|")
		output := strings.Split(line[1], " ")
		for j := 0; j < len(output); j++ {
			log.Println(output[j])
			sums[len(output[j])]++
		}
	}

	log.Println(sums)
	log.Println(sums[2] + sums[3] + sums[4] + sums[7])
}
