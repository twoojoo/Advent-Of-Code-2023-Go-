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

type NumSymbolsLine struct {
	numbers     []Number
	symbolsIdxs []int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nsLines := []NumSymbolsLine{}
	for _, l := range strings.Split(string(bytes), "\n") {
		nums, symIdxs := findNumbersAndSymbols(l)

		nsLines = append(nsLines, NumSymbolsLine{
			numbers:     nums,
			symbolsIdxs: symIdxs,
		})
	}

	sum := 0
	for i := 0; i < len(nsLines); i++ {
		prevLineSymIdxs := []int{}
		if i > 0 {
			prevLineSymIdxs = nsLines[i-1].symbolsIdxs
		}

		nextLineSymIdxs := []int{}
		if i < len(nsLines)-1 {
			nextLineSymIdxs = nsLines[i+1].symbolsIdxs
		}

		validNums := getValidNumbers(
			nsLines[i].numbers,
			nsLines[i].symbolsIdxs,
			prevLineSymIdxs,
			nextLineSymIdxs,
		)

		fmt.Println("line", i+1, ":", validNums)

		for _, n := range validNums {
			sum += n
		}
	}

	fmt.Println("RESULT:", sum)
}

func findNumbersAndSymbols(line string) ([]Number, []int) {
	numbers := []Number{}
	symbolsIdxs := []int{}

	lastNum := ""
	lastNumStartIdx := 0
	for idx, char := range line {
		isNum := slices.Contains(digits, char)
		isDot := char == '.'

		if lastNum == "" && isNum {
			lastNum += string(char)
			lastNumStartIdx = idx
			continue
		}

		if lastNum != "" && isNum {
			lastNum += string(char)
			continue
		}

		if isDot && lastNum == "" {
			continue
		}

		if !isDot && !isNum && lastNum == "" {
			symbolsIdxs = append(symbolsIdxs, idx)
			continue
		}

		if !isNum && lastNum != "" {
			if !isDot {
				symbolsIdxs = append(symbolsIdxs, idx)
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

		log.Fatalf("you missed some case %v, %v", lastNum, string(char))
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

	return numbers, symbolsIdxs
}

func getValidNumbers(lineNumbers []Number, lineSymbolsIdxs []int, prevLineSymIdxs []int, nextLineSymIdxs []int) []int {
	okNums := []int{}

	for _, n := range lineNumbers {
		currLineCandIdxs, prevLineCandIdxs, nextLineCandIdxs := getCandidateSymbolsIdxs(n)

		for _, clcIdx := range currLineCandIdxs {
			if slices.Contains(lineSymbolsIdxs, clcIdx) {
				okNums = append(okNums, n.value)
				break
			}
		}

		for _, plcIdx := range prevLineCandIdxs {
			if slices.Contains(prevLineSymIdxs, plcIdx) {
				okNums = append(okNums, n.value)
				break
			}
		}

		for _, nlcIdx := range nextLineCandIdxs {
			if slices.Contains(nextLineSymIdxs, nlcIdx) {
				okNums = append(okNums, n.value)
				break
			}
		}
	}

	return okNums
}

// find all indexes that are neighbours of a number
func getCandidateSymbolsIdxs(n Number) (currLineCandIdxs []int, prevLineCandIdxs []int, nextLineCandIdxs []int) {
	currLineCandIdxs = []int{}
	prevLineCandIdxs = []int{}
	nextLineCandIdxs = []int{}

	for i := n.startIdx - 1; i <= n.endIdx+1; i++ {
		if i == -1 {
			continue
		}

		if i == n.startIdx-1 || i == n.endIdx+1 {
			currLineCandIdxs = append(currLineCandIdxs, i)
		}

		prevLineCandIdxs = append(prevLineCandIdxs, i)
		nextLineCandIdxs = append(nextLineCandIdxs, i)
	}

	return currLineCandIdxs, prevLineCandIdxs, nextLineCandIdxs
}
