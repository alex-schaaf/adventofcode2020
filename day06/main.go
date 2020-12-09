package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
	groups := readFileLines("./input", "\n\n")

	var groupCountSum int
	var groupCountSum2 int
	for _, group := range groups {
		grp := strings.ReplaceAll(group, "\n", "")
		charMap := make(map[rune]int)
		for _, char := range grp {
			charMap[char]++
		}
		groupCountSum += len(charMap)

		// exercise 2
		groupLen := 1
		for _, char := range group {
			if char == rune('\n') {
				groupLen++
			}
		}

		for _, count := range charMap {
			if count == groupLen {
				groupCountSum2++
			}
		}

	}
	fmt.Printf("Exercise 1: %d questions answered with 'yes'.\n", groupCountSum)
	fmt.Printf("Exercise 2: %d questions answered collectively with 'yes' within groups.\n", groupCountSum2)
}
