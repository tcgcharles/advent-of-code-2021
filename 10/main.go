package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	sum := 0
	var sums []int
	for _, line := range input {
		stack := []string{}
		corrupted := false
		for i := 0; i < len(line); i++ {
			if line[i:i+1] == "(" || line[i:i+1] == "[" || line[i:i+1] == "{" || line[i:i+1] == "<" {
				stack = append(stack, line[i:i+1])
			} else if line[i:i+1] == ")" || line[i:i+1] == "]" || line[i:i+1] == "}" || line[i:i+1] == ">" {
				top := len(stack) - 1
				k := stack[top]
				stack = stack[:top]
				if line[i:i+1] == ")" && k != "(" || line[i:i+1] == "]" && k != "[" || line[i:i+1] == "}" && k != "{" || line[i:i+1] == ">" && k != "<" {
					log.Println(fmt.Sprintf("Expected %s, but found %s instead.", line[i:i+1], k))
					corrupted = true
					if line[i:i+1] == ")" {
						sum += 3
					} else if line[i:i+1] == "]" {
						sum += 57
					} else if line[i:i+1] == "}" {
						sum += 1197
					} else if line[i:i+1] == ">" {
						sum += 25137
					}
					break
				}
			}
		}
		if len(stack) > 0 && !corrupted {
			log.Println(stack)
			s := 0
			for len(stack) > 0 {
				top := len(stack) - 1
				k := stack[top]
				stack = stack[:top]
				s *= 5
				if k == "(" {
					s += 1
				} else if k == "[" {
					s += 2
				} else if k == "{" {
					s += 3
				} else if k == "<" {
					s += 4
				}
			}
			sums = append(sums, s)
		}
	}
	sort.Ints(sums)
	log.Println(sum)
	log.Println(sums)
	log.Println(sums[len(sums)/2 : len(sums)/2+1])
}
