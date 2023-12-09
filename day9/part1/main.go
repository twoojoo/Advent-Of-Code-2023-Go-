package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	sum := 0
	for i := range lines {
		values := stringSliceToNumbers(strings.Split(lines[i], " "))
		sum += processValues(values)
	}

	fmt.Println(sum)
}

func processValues(values []int) int {
	if isBaseCase(values) {
		return 0
	}

	diff := make([]int, len(values)-1)

	for i := 1; i < len(values); i++ {
		diff[i-1] = values[i] - values[i-1]
	}

	lastDiff := processValues(diff)

	return values[len(values)-1] + lastDiff
}

func isBaseCase(values []int) bool {
	for i := range values {
		if i != 0 {
			return false
		}
	}

	return true
}

func stringSliceToNumbers(strSl []string) []int {
	nums := make([]int, len(strSl))

	for i, strNum := range strSl {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal("error while converting number", err)
		}

		nums[i] = num
	}

	return nums
}
