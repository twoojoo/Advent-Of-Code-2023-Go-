package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type Number struct {
	value    int
	startIdx int
	endIdx   int
}

func (n Number) isPartiallyInRange(indexes []int) bool {
	return n.endIdx >= indexes[0] && n.startIdx <= indexes[len(indexes)-1]
}

type NumSymbolsLine struct {
	numbers   []Number
	starsIdxs []int
}

type Candidate struct {
	num     int
	starIdx int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nsLines := []NumSymbolsLine{}
	for _, l := range strings.Split(string(bytes), "\n") {
		nums, starsIdxs := findNumbersAndStars(l)

		nsLines = append(nsLines, NumSymbolsLine{
			numbers:   nums,
			starsIdxs: starsIdxs,
		})
	}

	sum := 0
	// graph := ds.NewGraph[int, int]()

	for i := 0; i < len(nsLines); i++ {
		prevLineNumbers := []Number{}
		if i > 0 {
			prevLineNumbers = nsLines[i-1].numbers
		}

		nextLineNumbers := []Number{}
		if i < len(nsLines)-1 {
			nextLineNumbers = nsLines[i+1].numbers
		}

		for _, starIdx := range nsLines[i].starsIdxs {
			currCandIdxs, prevCandIdxs, nextCandIdxs := getNeighbourPositions(starIdx)
			if len(currCandIdxs) > 2 {
				log.Fatal("length too high:", currCandIdxs)
			}

			gears := []int{}

			for _, n := range prevLineNumbers {
				if n.isPartiallyInRange(prevCandIdxs) {
					gears = append(gears, n.value)
				}
			}

			for _, n := range nextLineNumbers {
				if n.isPartiallyInRange(nextCandIdxs) {
					gears = append(gears, n.value)
				}
			}

			bef := currCandIdxs[0]
			for _, n := range nsLines[i].numbers {
				if n.endIdx == bef {
					gears = append(gears, n.value)
				}
			}

			aft := currCandIdxs[1]
			for _, n := range nsLines[i].numbers {
				if n.startIdx == aft {
					gears = append(gears, n.value)
				}
			}

			if len(gears) >= 3 {
				log.Println("line", i+1, ":", gears)
				log.Fatal("there are more than 2 neighbours")
			}

			if len(gears) == 2 {
				gearRatio := gears[0] * gears[1]
				log.Println("line", i+1, ":", gears[0], "*", gears[1], "=", gearRatio)
				sum += gearRatio
			}
		}

	}

	fmt.Println("RESULT:", sum)
}

func findNumbersAndStars(line string) ([]Number, []int) {
	numbers := []Number{}
	starsIdxs := []int{}

	lastNum := ""
	lastNumStartIdx := 0
	for idx, char := range line {
		isNum := slices.Contains(digits, char)
		isTrash := !isNum && char != '*'
		isStar := char == '*'

		if lastNum == "" && isNum {
			lastNum += string(char)
			lastNumStartIdx = idx
			continue
		}

		if lastNum != "" && isNum {
			lastNum += string(char)
			continue
		}

		if isTrash && lastNum == "" {
			continue
		}

		if isStar && lastNum == "" {
			starsIdxs = append(starsIdxs, idx)
			continue
		}

		if !isNum && lastNum != "" {
			if isStar {
				starsIdxs = append(starsIdxs, idx)
			}

			numVal, err := strconv.Atoi(lastNum)
			if err != nil {
				log.Fatal("failed converting", lastNum)
			}

			numbers = append(numbers, Number{
				value:    numVal,
				startIdx: lastNumStartIdx,
				endIdx:   idx - 1,
			})

			lastNum = ""
			lastNumStartIdx = 0

			continue
		}

		log.Fatalf("missed some case %v, %v", lastNum, string(char))
	}

	if lastNum != "" {
		numVal, err := strconv.Atoi(lastNum)
		if err != nil {
			log.Fatal("failed converting", lastNum)
		}

		numbers = append(numbers, Number{
			value:    numVal,
			startIdx: lastNumStartIdx,
			endIdx:   len(line) - 1,
		})
	}

	return numbers, starsIdxs
}

func matchWithCandidates(candidates []Candidate, idx int) (Candidate, bool) {
	for _, c := range candidates {
		if c.starIdx == idx {
			return c, true
		}
	}

	var zero Candidate
	return zero, false
}

func getNeighbourPositions(idx int) (currLineCandIdxs []int, prevLineCandIdxs []int, nextLineCandIdxs []int) {
	currLineCandIdxs = []int{}
	prevLineCandIdxs = []int{}
	nextLineCandIdxs = []int{}

	for i := idx - 1; i <= idx+1; i++ {
		if i == -1 {
			continue
		}

		if i == idx-1 || i == idx+1 {
			currLineCandIdxs = append(currLineCandIdxs, i)
		}

		prevLineCandIdxs = append(prevLineCandIdxs, i)
		nextLineCandIdxs = append(nextLineCandIdxs, i)
	}

	return currLineCandIdxs, prevLineCandIdxs, nextLineCandIdxs
}
