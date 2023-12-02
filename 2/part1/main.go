package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxBlue int = 14
var maxRed int = 12
var maxGreen int = 13

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	sum := 0

game:
	for _, l := range lines {
		parts := strings.Split(l, ":")
		head := parts[0]
		content := parts[1]

		gameID, err := strconv.Atoi(strings.Split(head, " ")[1])
		if err != nil {
			log.Fatal("failed gameID")
		}

		runs := strings.Split(content, ";")

		results := []bool{}
		for _, r := range runs {
			pick := strings.Split(r, ", ")

			var blue int
			var red int
			var green int
			for _, color := range pick {
				val := strings.Split(strings.Trim(color, " "), " ")

				count, err := strconv.Atoi(val[0])
				if err != nil {
					log.Fatal("failed to extract count")
				}

				if val[1] == "blue" {
					blue += count
					continue
				}

				if val[1] == "red" {
					red += count
					continue
				}

				if val[1] == "green" {
					green += count
					continue
				}
			}

			ok := isComboPossible(
				[]int{maxBlue, maxGreen, maxRed},
				[]int{blue, green, red},
			)

			results = append(results, ok)

			if !ok {
				break
			}
		}

		for _, r := range results {
			if !r {
				continue game
			}
		}

		sum += gameID
	}

	fmt.Println(sum)
}

func isComboPossible(maxVals []int, vals []int) bool {
	for i := range vals {
		if maxVals[i] < vals[i] {
			return false
		}
	}

	var sumMaxVals int
	for _, v := range maxVals {
		sumMaxVals += v
	}

	var sumVals int
	for _, v := range vals {
		sumVals += v
	}

	if sumVals > sumMaxVals {
		return false
	}

	for i := 0; i < len(vals); i++ {
		nextI := i + 1
		if nextI == len(vals) {
			nextI = 0
		}

		if vals[i]+vals[nextI] > maxVals[i]+maxVals[nextI] {
			return false
		}
	}

	return true
}
