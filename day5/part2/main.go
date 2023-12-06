package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type SeedRange struct {
	start int
	end   int
}

type Mapping struct {
	source      int
	destination int
	_range      int
	sourceMax   int
}

type Mappings struct {
	mappings  []Mapping
	sourceMax int
	sourceMin int
}

func (m Mappings) getDestination(source int) int {
	for i := range m.mappings {
		max := m.mappings[i].source + m.mappings[i]._range
		if source >= m.mappings[i].source && source < max {
			return m.mappings[i].destination + (source - m.mappings[i].source)
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

	//sees are a range of seeds now
	seeds := getSeeds(lines)

	wg := sync.WaitGroup{}

	// dynamically iterate over all mappings in order for each seed
	for i, seed := range seeds {
		s := seed
		l := i
		//iterate over seed range (REALLY SLOW, MILIONS OF ROWS)
		wg.Add(1)
		go func() {
			lowest := math.MaxInt
			_range := s.end - s.start
			fmt.Println("seed", l, "range", _range)
			for i := 0; i < _range; i++ {

				lastDest := s.start + i
				for j := range mappings {
					if lastDest > mappings[j].sourceMax || lastDest < mappings[j].sourceMin {
						continue
					}

					lastDest = mappings[j].getDestination(lastDest)
				}

				if lastDest < lowest {
					lowest = lastDest
				}
			}

			wg.Done()
			fmt.Println("seed", l, "done:", lowest)
		}()

	}
	wg.Wait()
}

func getSeeds(input []string) []SeedRange {
	seedsLine := input[0]
	seedsLine = strings.Split(seedsLine, ": ")[1]
	seedsNumStr := strings.Split(seedsLine, " ")

	seedsNum := make([]SeedRange, len(seedsNumStr)/2)
	for i := 0; i < len(seedsNumStr); i += 2 {
		end, err := strconv.Atoi(seedsNumStr[i])
		if err != nil {
			log.Fatal(err)
		}

		start, err := strconv.Atoi(seedsNumStr[i+1])
		if err != nil {
			log.Fatal(err)
		}

		if end < start {
			temp := start
			start = end
			end = temp
		}

		seedsNum[i/2] = SeedRange{
			start: start,
			end:   end,
		}
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

	mappings := Mappings{
		mappings:  []Mapping{},
		sourceMax: math.MinInt,
		sourceMin: math.MaxInt,
	}

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

		if sourceNum+rangeNum > mappings.sourceMax {
			mappings.sourceMax = sourceNum + rangeNum
		}

		if sourceNum < mappings.sourceMin {
			mappings.sourceMin = sourceNum
		}

		mappings.mappings = append(mappings.mappings, Mapping{
			_range:      rangeNum,
			destination: destNum,
			source:      sourceNum,
		})
	}

	return mappings
}
