package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFileLines(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	lines := readFileLines("input")

	// parse
	var min []int
	var max []int
	var character []string
	var passw []string

	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		nRange := strings.Split(splitLine[0], "-")

		nMin, _ := strconv.Atoi(nRange[0])
		nMax, _ := strconv.Atoi(nRange[1])
		min = append(min, nMin)
		max = append(max, nMax)
		character = append(character, splitLine[1])
		passw = append(passw, splitLine[2])
	}

	var count1 int
	var count2 int
	// validation
	for i := range lines {
		isValid := validatePassword(min[i], max[i], character[i], passw[i])
		if isValid {
			count1++
		}
		isValid2 := validatePassword2(min[i], max[i], character[i], passw[i])
		if isValid2 {
			count2++
		}
	}
	fmt.Printf("Exercise 1: Counted %d valid passwords.\n", count1)
	fmt.Printf("Exercise 2: Counted %d valid passwords.\n", count2)
}

func validatePassword(min int, max int, character string, passw string) bool {
	var count int
	for _, char := range passw {
		if char == rune(character[0]) {
			count++
		}
	}
	if count >= min && count <= max {
		return true
	}
	return false
}

func validatePassword2(min int, max int, character string, passw string) bool {
	char := rune(character[0])
	char1 := rune(passw[min-1])
	char2 := rune(passw[max-1])

	if char1 == char && char2 != char {
		return true
	} else if char1 != char && char2 == char {
		return true
	}
	return false
}
