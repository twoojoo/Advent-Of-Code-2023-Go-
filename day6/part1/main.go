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

	racesTimes, err := strSliceToNumSlice(strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("times:", racesTimes)

	racesMaxDist, err := strSliceToNumSlice(strings.Split(strings.Trim(strings.Split(lines[1], ":")[1], " "), " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("records", racesMaxDist)

	results := []int{}
	for i := 0; i < len(racesTimes); i++ {
		raceTime := racesTimes[i]
		maxDist := racesMaxDist[i]

		sumWaysToWin := 0
		for chargeTime := 1; chargeTime < raceTime; chargeTime++ {
			speed := chargeTime * 1
			travelTime := raceTime - chargeTime
			distance := travelTime * speed

			if distance > maxDist {
				sumWaysToWin++
			}
		}

		fmt.Println("race", i, ", ways to win:", sumWaysToWin)

		if sumWaysToWin > 0 {
			results = append(results, sumWaysToWin)
		}
	}

	mult := 1
	for _, r := range results {
		mult *= r
	}

	fmt.Println("RESULT:", mult)
}

func strSliceToNumSlice(s []string) ([]int, error) {
	numS := []int{}

	for i := range s {
		if s[i] == "" {
			continue
		}

		num, err := strconv.Atoi(s[i])
		if err != nil {
			return numS, err
		}

		numS = append(numS, num)
	}

	return numS, nil
}
