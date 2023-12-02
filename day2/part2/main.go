package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	sum := 0
	for _, l := range lines {
		parts := strings.Split(l, ":")
		content := parts[1]

		runs := strings.Split(content, ";")

		sets := make([][]int, len(runs))
		for i, r := range runs {
			sets[i] = []int{0, 0, 0}

			pick := strings.Split(r, ", ")

			for _, color := range pick {
				val := strings.Split(strings.Trim(color, " "), " ")

				count, err := strconv.Atoi(val[0])
				if err != nil {
					log.Fatal("failed to extract count")
				}

				if val[1] == "blue" {
					sets[i][0] += count
					continue
				}

				if val[1] == "red" {
					sets[i][1] += count
					continue
				}

				if val[1] == "green" {
					sets[i][2] += count
					continue
				}
			}
		}

		minSet := findMinimumSet(sets)

		var power int
		for i, n := range minSet {
			if i == 0 {
				power = n
			} else if n != 0 {
				power = power * n
			}
		}

		sum += power
	}

	fmt.Println("RESULT:", sum)
}

func findMinimumSet(vals [][]int) []int {
	minSet := make([]int, len(vals[0]))

	for i := range vals[0] {
		var max int
		for j := 0; j < len(vals); j++ {
			if vals[j][i] > max {
				max = vals[j][i]
			}
		}

		minSet[i] = max
	}

	return minSet
}
