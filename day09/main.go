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

func inArray(n int64, arr []int64) bool {
	for _, num := range arr {
		if n == num {
			return true
		}
	}
	return false
}

func sumArr(arr []int64) int64 {
	var sum int64
	for _, num := range arr {
		sum += num
	}
	return sum
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

	var sumRange int = 25
	for i, number := range numbers {
		if i < sumRange {
			continue
		}
		sums := sumCombinations(numbers[i-sumRange : i])
		isValid := inArray(number, sums)

		if !isValid {
			fmt.Printf("Exercise 1: %d is not the sum of any two previous %d numbers.\n", number, sumRange)

			for j := 0; j < i; j++ {
				for size := 2; j+size < len(numbers); size++ {
					numRange := numbers[j : j+size]
					sum := sumArr(numRange)
					if sum == number {
						fmt.Println(number, numRange)
						min := numRange[0]
						max := numRange[0]
						for _, num := range numRange {
							if num < min {
								min = num
							} else if num > max {
								max = num
							}
						}
						fmt.Printf("Exercise 2: %d + %d = %d.\n", min, max, min+max)
						return
					}
				}
			}
		}
	}
}
