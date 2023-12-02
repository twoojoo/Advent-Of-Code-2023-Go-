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

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	list := string(bytes)

	records := strings.Split(list, "\n")

	var values = []int{}
	for _, r := range records {
		first := '-'
		last := '-'

		for _, char := range r {
			if slices.Contains(digits, char) {
				if first == '-' {
					first = char
					last = char
				} else {
					last = char
				}
			}
		}

		value := string(first) + string(last)

		valNum, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, valNum)
	}

	var sum int
	for _, v := range values {
		sum += v
	}

	fmt.Println(sum)
}
