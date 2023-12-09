package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	AllDifferent = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const (
	JokerVal = 1
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
				val = JokerVal
			case "T":
				val = 10
			default:
				val, err = strconv.Atoi(char)
				if err != nil {
					return games, err
				}

				if val == JokerVal {
					log.Fatal("invalid joker value")
				}
			}

			games[i].hand = append(games[i].hand, val)
		}
	}

	return games, nil
}

func getHandValue(hand []int) int {
	if !slices.Contains(hand, JokerVal) {
		return getHandValueNoJokers(hand)
	}

	return getHandValueWithJokers(hand)
}

func getHandValueNoJokers(hand []int) int {
	seeds := countSeeds(hand)

	if len(seeds) == 1 {
		return FiveOfAKind
	}

	if len(seeds) == 5 {
		return AllDifferent
	}

	if len(seeds) == 4 {
		return OnePair
	}

	if len(seeds) == 2 {
		if _, ok := mapContains(seeds, 4); ok {
			return FourOfAKind
		}
		return FullHouse
	}

	if len(seeds) == 3 {
		if _, ok := mapContains(seeds, 3); ok {
			return ThreeOfAKind
		}
		return TwoPairs
	}

	log.Fatal("unknown case", seeds)
	return 0
}

func getHandValueWithJokers(hand []int) int {
	orderdSeeds, jokersCount := orderSeedsByFrequenceAndGetJokers(hand)

	if len(orderdSeeds) == 0 {
		return FiveOfAKind
	}

	mostFrequent := orderdSeeds[0]
	for i := 0; i < jokersCount; i++ {
		orderdSeeds = slices.Insert(orderdSeeds, 0, mostFrequent)
	}

	if len(orderdSeeds) != 5 {
		log.Fatal("something wrong:", orderdSeeds)
	}

	return getHandValueNoJokers(orderdSeeds)
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

func orderSeedsByFrequenceAndGetJokers(slice []int) ([]int, int) {
	seeds := countSeeds(slice)

	var jokers int
	if j, ok := seeds[JokerVal]; ok {
		jokers = j
	}

	ordered := []int{}
	for i := 5; i >= 1; i-- {
		ok := true
		excluding := []int{JokerVal}

		for ok {
			var key int
			if key, ok = mapContains(seeds, i, excluding...); ok {
				excluding = append(excluding, key)
				for j := 0; j < i; j++ {
					ordered = append(ordered, key)
				}
			}
		}
	}

	return ordered, jokers
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

func mapContains[K, V comparable](m map[K]V, v V, excluding ...K) (K, bool) {
	for k := range m {
		if slices.Contains(excluding, k) {
			continue
		}

		if val, ok := m[k]; ok {
			if val == v {
				return k, true
			}
		}
	}

	var zero K
	return zero, false
}
