package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type digit struct {
	value int
	rune  rune
	word  string
}

var digits = []digit{{
	value: 1,
	rune:  '1',
	word:  "one",
}, {
	value: 2,
	rune:  '2',
	word:  "two",
}, {
	value: 3,
	rune:  '3',
	word:  "three",
}, {
	value: 4,
	rune:  '4',
	word:  "four",
}, {
	value: 5,
	rune:  '5',
	word:  "five",
}, {
	value: 6,
	rune:  '6',
	word:  "six",
}, {
	value: 7,
	rune:  '7',
	word:  "seven",
}, {
	value: 8,
	rune:  '8',
	word:  "eight",
}, {
	value: 9,
	rune:  '9',
	word:  "nine",
}}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	records := strings.Split(string(bytes), "\n")

	var values = []int{}
	for _, r := range records {
		first := 100
		last := 100

		contained := [][]int{}
		for _, d := range digits {
			i := strings.Index(r, d.word)
			if i != -1 {
				contained = append(contained, []int{i, d.value})
			}

			i = strings.Index(r, string(d.rune))
			if i != -1 {
				contained = append(contained, []int{i, d.value})
			}

			i = strings.LastIndex(r, d.word)
			if i != -1 {
				contained = append(contained, []int{i, d.value})
			}

			i = strings.LastIndex(r, string(d.rune))
			if i != -1 {
				contained = append(contained, []int{i, d.value})
			}
		}

		max := []int{-1, 0}
		min := []int{100, 0}
		for _, c := range contained {
			if c[0] > max[0] {
				max = c
			}

			if c[0] < min[0] {
				min = c
			}
		}

		last = max[1]
		first = min[1]
		if last == 100 || first == 100 {
			log.Fatal(last, first)
		}

		value := strconv.Itoa(first) + strconv.Itoa(last)
		log.Println(r, "==>", value)

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

	log.Println("TOTAL:", sum)
}
