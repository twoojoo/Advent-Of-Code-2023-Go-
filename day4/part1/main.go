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

		sum += points
	}

	fmt.Println("RESULT:", sum)
}

func gameFromString(line string) Game {
	split := strings.Split(line, ":")
	body := split[1]

	body = strings.Replace(body, "  ", " ", -1)
	split = strings.Split(body, "|")
	numbersStr := strings.Trim(split[0], " ")
	winningStr := strings.Trim(split[1], " ")

	numbersStrArr := strings.Split(numbersStr, " ")
	winningStrArr := strings.Split(winningStr, " ")

	numbers := stringSliceToNumbers(numbersStrArr)
	winning := stringSliceToNumbers(winningStrArr)

	return Game{
		numbers: numbers,
		winning: winning,
	}
}

func stringSliceToNumbers(strSl []string) []int {
	nums := make([]int, len(strSl))

	for i, strNum := range strSl {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal("error while converting number", err)
		}

		nums[i] = num
	}

	return nums
}
