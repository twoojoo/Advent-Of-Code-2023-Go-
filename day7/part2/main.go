package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	AllDifferent = 0
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Game struct {
	hand []int
	bid  int
}

type Games []Game

func (g Games) Len() int {
	return len(g)
}

// sort games of the samve val
func (g Games) Less(i, j int) bool {
	for k := range g[i].hand {
		if g[i].hand[k] > g[j].hand[k] {
			return false
		}

		if g[i].hand[k] < g[j].hand[k] {
			return true
		}
	}

	panic("unexpected equality")
}

func (g Games) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	games, err := linesToGames(lines)
	if err != nil {
		log.Fatal(err)
	}

	gamesByHandVal := groupGamesByHandValue(games)
	for v := range gamesByHandVal {
		sort.Sort(gamesByHandVal[v])
	}

	fmt.Println(countPoints(gamesByHandVal))

}

func linesToGames(lines []string) (games Games, err error) {
	games = make([]Game, len(lines))

	for i := range games {
		split := strings.Split(lines[i], " ")

		games[i].bid, err = strconv.Atoi(split[1])
		if err != nil {
			return games, err
		}

		for _, char := range strings.Split(split[0], "") {
			var val int

			switch char {
			case "A":
				val = 14
			case "K":
				val = 13
			case "Q":
				val = 12
			case "J":
				val = 11
			case "T":
				val = 10
			default:
				val, err = strconv.Atoi(char)
				if err != nil {
					return games, err
				}
			}

			games[i].hand = append(games[i].hand, val)
		}
	}

	return games, nil
}

func getHandValue(hand []int) int {
	seeds := countSeeds(hand)

	if len(seeds) == 1 {
		return FiveOfAKind
	}

	if len(seeds) == 5 {
		if jokers, ok := seeds[11]; ok {
			if jokers == 1 {
				return OnePair
			}
			log.Fatal("impossible", seeds)
		}
		return AllDifferent
	}

	if len(seeds) == 4 {
		if jokers, ok := seeds[11]; ok {
			if jokers == 2 || jokers == 1 {
				return ThreeOfAKind
			}

			log.Fatal("impossible", seeds)
		}

		return OnePair
	}

	if len(seeds) == 3 {
		for k := range seeds {
			if seeds[k] == 3 { //normally is 2 of a kind
				if jokers, ok := seeds[11]; ok {
					if jokers == 3 || jokers == 1 {
						return FourOfAKind
					}

					if jokers == 2 {
						return FiveOfAKind
					}
				}

				return ThreeOfAKind
			}
		}

		//should be one pair
		if jokers, ok := seeds[11]; ok {
			if jokers == 2 || jokers == 1 {
				return ThreeOfAKind
			}

			log.Fatal("impossible", seeds)
		}

		return OnePair
	}

	if len(seeds) == 2 {
		for k := range seeds {
			if seeds[k] == 4 { //should be 4 of a kind
				if _, ok := seeds[11]; ok {
					// if jokers == 4 || jokers == 3 || jokers == 2 {
					return FiveOfAKind
					// }
				}
			}
		}

		//should be full house
		if _, ok := seeds[11]; ok {
			// if jokers == 3 {
			return FiveOfAKind
			// }
		}

		return FullHouse
	}

	log.Fatal("unknown case", hand)
	return 0

}

func countSeeds(slice []int) map[int]int {
	set := map[int]int{}

	for _, v := range slice {
		if _, ok := set[v]; ok {
			set[v]++
		} else {
			set[v] = 1
		}
	}

	return set
}

func groupGamesByHandValue(games Games) map[int]Games {
	grouped := map[int]Games{}

	for i := range games {
		val := getHandValue(games[i].hand)
		if _, ok := grouped[val]; ok {
			grouped[val] = append(grouped[val], games[i])
		} else {
			grouped[val] = []Game{games[i]}
		}
	}

	return grouped
}

func countPoints(orederdGames map[int]Games) int {
	total := 0
	counter := 1
	for i := 0; i < 7; i++ {
		for j := range orederdGames[i] {
			total += (counter * orederdGames[i][j].bid)
			counter++
		}
	}

	return total
}
