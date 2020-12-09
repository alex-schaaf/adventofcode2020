package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func sumCombinations(arr []int64) []int64 {
	var sums []int64
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sums = append(sums, arr[i]+arr[j])
		}
	}
	return sums
}

func main() {
	lines := readFileLines("./input", "\n")
	var numbers []int64
	for _, line := range lines {
		number, err := strconv.ParseInt(strings.TrimSpace(line), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// exercise 1
	for i, number := range numbers {
		isValid := false
		if i < 25 {
			continue
		}
		sums := sumCombinations(numbers[i-25 : i])

		for _, sum := range sums {
			if number == sum {
				isValid = true
			}
		}
		if !isValid {
			fmt.Printf("Exercise 1: %d is not the sum of any two previous 25 numbers.", number)
			break
		}
	}
}
