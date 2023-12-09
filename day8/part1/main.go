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

	nextStep := "AAA"
	count := 0

	destinationReached := false

	for !destinationReached {
		for _, direction := range strings.Split(lines[0], "") {
			if direction == "R" {
				nextStep = steps[nextStep].right
			} else if direction == "L" {
				nextStep = steps[nextStep].left
			} else {
				log.Fatal("unknown direction", direction)
			}

			count++

			if nextStep == "ZZZ" {
				destinationReached = true
				break
			}
		}
	}

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
