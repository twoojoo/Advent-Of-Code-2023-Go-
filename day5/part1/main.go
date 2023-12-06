package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	numbers []int
	winning []int
}

type Mapping struct {
	source      int
	destination int
	_range      int
}

type Mappings []Mapping

func (m Mappings) getDestination(source int) int {
	for i := range m {
		if source >= m[i].source && source < m[i].source+m[i]._range {
			sourceRange := source - m[i].source
			return m[i].destination + sourceRange
		}
	}

	return source
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	//dynamically get all mapping lines with their name
	mappingLines, mappingNames := splitMapingLinesAndGetNames(lines)

	// dynamically get oreder maps
	mappings := make([]Mappings, len(mappingNames))
	for i, name := range mappingNames {
		mappings[i] = getMappings(name, mappingLines)
	}

	seeds := getSeeds(lines)

	// dynamically iterate over all mappings in order for each seed
	lowest := math.MaxInt
	for _, seed := range seeds {
		lastDest := seed

		for _, mapping := range mappings {
			lastDest = mapping.getDestination(lastDest)
		}

		if lastDest < lowest {
			lowest = lastDest
		}
	}

	fmt.Println(lowest)
}

func getSeeds(input []string) []int {
	seedsLine := input[0]
	seedsLine = strings.Split(seedsLine, ": ")[1]
	seedsNumStr := strings.Split(seedsLine, " ")

	seedsNum := []int{}
	for i := range seedsNumStr {
		num, err := strconv.Atoi(seedsNumStr[i])
		if err != nil {
			log.Fatal(err)
		}

		seedsNum = append(seedsNum, num)
	}

	return seedsNum
}

func splitMapingLinesAndGetNames(input []string) ([][]string, []string) {
	groupedLines := [][]string{}
	names := []string{}
	glIdx := -1
	for _, line := range input[2:] {
		if len(line) == 0 {
			continue
		}

		if _, err := strconv.Atoi(string(line[0])); err != nil {
			groupedLines = append(groupedLines, []string{line})
			names = append(names, line)
			glIdx++
		} else {
			groupedLines[glIdx] = append(groupedLines[glIdx], line)
		}
	}

	return groupedLines, names
}

func getMappings(name string, groupedLines [][]string) Mappings {
	group := []string{}

	for i := range groupedLines {
		if strings.Contains(groupedLines[i][0], name) {
			group = groupedLines[i]
			break
		}
	}

	mappings := []Mapping{}
	for i := range group {
		if i == 0 {
			continue
		}

		stringNums := strings.Split(group[i], " ")

		destNum, err := strconv.Atoi(stringNums[0])
		if err != nil {
			log.Fatal(err)
		}

		sourceNum, err := strconv.Atoi(stringNums[1])
		if err != nil {
			log.Fatal(err)
		}

		rangeNum, err := strconv.Atoi(stringNums[2])
		if err != nil {
			log.Fatal(err)
		}

		mappings = append(mappings, Mapping{
			_range:      rangeNum,
			destination: destNum,
			source:      sourceNum,
		})
	}

	return mappings
}
