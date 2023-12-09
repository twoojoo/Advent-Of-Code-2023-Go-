package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Step struct {
	left  string
	right string
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	steps := parseSteps(lines[2:])

	count := traverseAndGetStepsCount(steps, lines[0])

	fmt.Println(count)
}

func parseSteps(lines []string) map[string]Step {
	steps := map[string]Step{}

	for i := range lines {
		split1 := strings.Split(lines[i], " = (")
		name := split1[0]

		split2 := strings.Split(split1[1], ", ")
		left := split2[0]

		split3 := strings.Split(split2[1], ")")
		right := split3[0]

		if _, ok := steps[name]; ok {
			log.Fatal("duplicated step")
		}

		steps[name] = Step{
			right: right,
			left:  left,
		}
	}

	return steps
}

func getStartingSteps(steps map[string]Step) []string {
	starging := []string{}

	for k := range steps {
		if strings.HasSuffix(k, "A") {
			starging = append(starging, k)
		}
	}

	return starging
}

func traverseAndGetStepsCount(steps map[string]Step, directions string) int {
	startingSteps := getStartingSteps(steps)
	destinationReached := false
	count := 0
	nextSteps := startingSteps
	fmt.Println(nextSteps)

	for !destinationReached {
		for _, direction := range strings.Split(directions, "") {
			for i, ns := range nextSteps {
				if direction == "R" {
					nextSteps[i] = steps[ns].right
				} else if direction == "L" {
					nextSteps[i] = steps[ns].left
				} else {
					log.Fatal("unknown direction", direction)
				}

			}

			count++

			if isArrived(nextSteps) {
				destinationReached = true
				break
			}
		}
	}

	return count
}

func isArrived(steps []string) bool {
	for i := range steps {
		if !strings.HasSuffix(steps[i], "Z") {
			return false
		}
	}

	return true
}
