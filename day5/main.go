package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const lowerRow = 70
const upperRow = 66

const lowerCol = 76
const upperCol = 82

func readFileLines(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func binaryPartition(pass string, low int, high int, lower byte, higher byte) int {

	char := pass[0]
	// fmt.Println(char, low, high)
	if char == lower {
		high = (high-low)/2 + low
	} else {
		low = (high-low)/2 + low + 1
	}

	if len(pass) <= 1 {
		if char == lower {
			return low
		} else {
			return high
		}
	}

	return binaryPartition(pass[1:], low, high, lower, higher)
}

func main() {
	boardingPasses := readFileLines("./input")
	// fmt.Println("L:", []rune("L")[0], "R:", []rune("R")[0])
	var seatIDs []int
	var maxID int

	for _, pass := range boardingPasses {
		rowCode := pass[:len(pass)-3]
		colCode := pass[len(pass)-3:]
		// fmt.Println(rowCode, colCode)

		row := binaryPartition(rowCode, 0, 127, lowerRow, upperRow)
		col := binaryPartition(colCode, 0, 7, lowerCol, upperCol)

		// fmt.Println("row:", row)
		// fmt.Println("col:", col)

		seatID := row*8 + col
		if seatID > maxID {
			maxID = seatID
		}

		// fmt.Println("seadID:", seatID)
		seatIDs = append(seatIDs, seatID)

	}
	fmt.Printf("Exercise 1: The highest seat ID is %d.\n", maxID)
}
