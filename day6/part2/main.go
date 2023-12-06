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

	raceTime, err := strSliceToTotal(strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("time:", raceTime)

	raceRecord, err := strSliceToTotal(strings.Split(strings.Trim(strings.Split(lines[1], ":")[1], " "), " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("record", raceRecord)

	sumWaysToWin := 0
	for chargeTime := 1; chargeTime < raceTime; chargeTime++ {
		speed := chargeTime * 1
		travelTime := raceTime - chargeTime
		distance := travelTime * speed

		if distance > raceRecord {
			sumWaysToWin++
		}
	}

	fmt.Println("RESULT:", sumWaysToWin)
}

func strSliceToTotal(s []string) (int, error) {
	numS := ""

	for i := range s {
		if s[i] == "" {
			continue
		}

		numS += s[i]
	}

	num, err := strconv.Atoi(numS)
	if err != nil {
		return 0, err
	}

	return num, nil
}
