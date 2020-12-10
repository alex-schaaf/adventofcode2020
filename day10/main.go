package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func parseLinesInt(lines []string) []int {
	var numbers []int
	for _, line := range lines {
		number, _ := strconv.Atoi(strings.TrimSpace(line))
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	lines := readFileLines("./test", "\n")
	joltages := parseLinesInt(lines)
	for i, jolt := range joltages {
		fmt.Println(i, "\t", jolt)
	}
}
