package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

const tree = 35

func main() {
	lines := readFileLines("./input")
	var xSteps = []int{1, 3, 5, 7, 1}
	var ySteps = []int{1, 1, 1, 1, 2}
	var multiply = 1
	for i, xStep := range xSteps {
		yStep := ySteps[i]
		treeCounter := countTrees(lines, xStep, yStep)
		fmt.Printf("(%d, %d): %d\n", xStep, yStep, treeCounter)
		multiply *= treeCounter
	}
	fmt.Printf("Multiplied together: %d\n", multiply)
}

func countTrees(lines []string, xStep int, yStep int) int {
	repeater := int(math.Ceil(float64(len(lines))/float64(len(lines[0])))) * xStep
	var treeCounter int
	var x int
	for i, line := range lines {
		if yStep == 2 {
			if i%2 != 0 {
				continue
			} else {
				x = i / 2
			}
		} else {
			x = i * xStep
		}
		line := strings.Repeat(line, repeater)

		if line[x] == tree {
			treeCounter++
		}
	}
	return treeCounter
}
