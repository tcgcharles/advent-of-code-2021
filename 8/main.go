package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
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
*/

type Line struct {
	Line  []string
	Line0 []string
	Line1 []string
	Num   int
}

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	lines := make([]Line, len(input))
	for i := 0; i < len(input); i++ {
		l := strings.Split(input[i], " | ")
		line0 := strings.Split(l[0], " ")
		line1 := strings.Split(l[1], " ")
		for i := 0; i < len(line0); i++ {
			line0[i] = alphabetize(line0[i])
		}
		for i := 0; i < len(line1); i++ {
			line1[i] = alphabetize(line1[i])
		}
		lines[i] = Line{
			Line:  append(line0, line1...),
			Line0: line0,
			Line1: line1,
		}
	}

	sum := 0
	for _, v := range lines {
		one := oneSegments(v.Line)
		four := fourSegments(v.Line)
		seven := sevenSegments(v.Line)
		eight := eightSegments(v.Line)
		three := threeSegments(v.Line)
		nine := nineSegments(v.Line)
		zero := zeroSegments(v.Line)
		six := sixSegments(v.Line)
		five := fiveSegments(v.Line)
		two := twoSegments(v.Line)
		for i := 0; i < len(v.Line1); i++ {
			if v.Line1[i] == one {
				v.Line1[i] = "1"
			} else if v.Line1[i] == four {
				v.Line1[i] = "4"
			} else if v.Line1[i] == seven {
				v.Line1[i] = "7"
			} else if v.Line1[i] == eight {
				v.Line1[i] = "8"
			} else if v.Line1[i] == three {
				v.Line1[i] = "3"
			} else if v.Line1[i] == nine {
				v.Line1[i] = "9"
			} else if v.Line1[i] == zero {
				v.Line1[i] = "0"
			} else if v.Line1[i] == six {
				v.Line1[i] = "6"
			} else if v.Line1[i] == five {
				v.Line1[i] = "5"
			} else if v.Line1[i] == two {
				v.Line1[i] = "2"
			}
		}
		v.Num = n(v.Line1)
		sum += v.Num
	}
	log.Println(sum)
}

func n(line []string) int {
	sum := 0
	pow := 0
	for i := len(line) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(line[i])
		if err != nil {
			log.Fatal(err)
		}
		sum += int(math.Pow10(pow)) * num
		pow++
	}
	return sum
}

func isZero(s string, line []string) bool {
	return len(s) == 6 && !isNine(s, line) && contains(s, sevenSegments(line))
}

func zeroSegments(line []string) string {
	for _, v := range line {
		if isZero(v, line) {
			return v
		}
	}
	return ""
}

func isOne(s string) bool {
	return len(s) == 2
}

func oneSegments(line []string) string {
	for _, v := range line {
		if isOne(v) {
			return v
		}
	}
	return ""
}

func isTwo(s string, line []string) bool {
	return len(s) == 5 && !isThree(s, line) && !isFive(s, line)
}

func twoSegments(line []string) string {
	for _, v := range line {
		if isTwo(v, line) {
			return v
		}
	}
	return ""
}

func isThree(s string, line []string) bool {
	return len(s) == 5 && contains(s, sevenSegments(line))
}

func threeSegments(line []string) string {
	for _, v := range line {
		if isThree(v, line) {
			return v
		}
	}
	return ""
}

func isFour(s string) bool {
	return len(s) == 4
}

func fourSegments(line []string) string {
	for _, v := range line {
		if isFour(v) {
			return v
		}
	}
	return ""
}

func isFive(s string, line []string) bool {
	four := fourSegments(line)
	return len(s) == 5 && !isThree(s, line) && (contains(s, four[0:3]) || contains(s, four[1:4]) || contains(s, four[0:2]+four[3:4]) || contains(s, four[0:1]+four[2:4]))
}

func fiveSegments(line []string) string {
	for _, v := range line {
		if isFive(v, line) {
			return v
		}
	}
	return ""
}

func isSix(s string, line []string) bool {
	return len(s) == 6 && !isNine(s, line) && !isZero(s, line)
}

func sixSegments(line []string) string {
	for _, v := range line {
		if isSix(v, line) {
			return v
		}
	}
	return ""
}

func isSeven(s string) bool {
	return len(s) == 3
}

func sevenSegments(line []string) string {
	for _, v := range line {
		if isSeven(v) {
			return v
		}
	}
	return ""
}

func isEight(s string) bool {
	return len(s) == 7
}

func eightSegments(line []string) string {
	for _, v := range line {
		if isEight(v) {
			return v
		}
	}
	return ""
}

func isNine(s string, line []string) bool {
	return len(s) == 6 && contains(s, fourSegments(line))
}

func nineSegments(line []string) string {
	for _, v := range line {
		if isNine(v, line) {
			return v
		}
	}
	return ""
}

func alphabetize(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}

func contains(s0, s1 string) bool {
	c := true
	for i := 0; i < len(s1); i++ {
		c = c && strings.Contains(s0, s1[i:i+1])
		if !c {
			break
		}
	}
	return c
}
