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
	for _, group := range groups {
		grp := strings.ReplaceAll(group, "\n", "")
		fmt.Println(grp)
		charMap := make(map[rune]int)
		for _, char := range grp {
			charMap[char]++
		}
		groupCountSum += len(charMap)
		fmt.Println("---")
	}
	fmt.Printf("Exercise 1: %d questions answered with 'yes'.\n", groupCountSum)
}
