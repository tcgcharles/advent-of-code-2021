package main

import (
	"fmt"
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

	x := 0
	y := 0

	commands := strings.Split(string(f), "\n")
	for i := 0; i < len(commands); i++ {
		command := strings.Split(commands[i], " ")
		c := command[0]
		d, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		if c == "up" {
			d *= -1
		}

		if c == "forward" {
			x += d
		} else {
			y += d
		}
	}

	log.Println(fmt.Sprintf("%d,%d", x, y))
	log.Println(x * y)

	x = 0
	depth := 0
	aim := 0

	commands = strings.Split(string(f), "\n")
	for i := 0; i < len(commands); i++ {
		command := strings.Split(commands[i], " ")
		c := command[0]
		d, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		if c == "up" {
			d *= -1
		}

		if c == "forward" {
			x += d
			depth += d * aim
		} else {
			aim += d
		}
	}

	log.Println(fmt.Sprintf("%d,%d", x, depth))
	log.Println(x * depth)
}
