package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	ID      string
	numbers []int
	winning []int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	games := make([]Game, len(lines))
	for i, l := range lines {
		game := gameFromString(l)
		games[i] = game
	}

	sum := 0
	for _, g := range games {
		points := 0
		for _, n := range g.numbers {
			if slices.Contains(g.winning, n) {
				if points == 0 {
					points++
				} else {
					points *= 2
				}
			}
		}

		fmt.Println(g.ID, g.numbers, g.winning, "|", points)

		sum += points
	}

	fmt.Println("RESULT:", sum)
}

func gameFromString(line string) Game {
	split := strings.Split(line, ":")
	head := split[0]
	body := split[1]

	ID := strings.Split(head, " ")[1]

	body = strings.Replace(body, "  ", " ", -1)
	split = strings.Split(body, "|")
	numbersStr := strings.Trim(split[0], " ")
	winningStr := strings.Trim(split[1], " ")

	numbersStrArr := strings.Split(numbersStr, " ")
	winningStrArr := strings.Split(winningStr, " ")

	numbers := make([]int, len(numbersStrArr))
	winning := make([]int, len(winningStrArr))

	for i, numStr := range numbersStrArr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("error while converting number", err)
		}

		numbers[i] = num
	}

	for i, winStr := range winningStrArr {
		num, err := strconv.Atoi(winStr)
		if err != nil {
			log.Fatal("error while converting number", err)
		}

		winning[i] = num
	}

	return Game{
		ID:      ID,
		numbers: numbers,
		winning: winning,
	}
}
