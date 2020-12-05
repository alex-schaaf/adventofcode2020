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

func getSeatID(row int, col int) int {
	return row*8 + col
}

func main() {
	boardingPasses := readFileLines("./input")
	var seatIDs []int
	var maxID int

	for _, pass := range boardingPasses {
		rowCode := pass[:len(pass)-3]
		colCode := pass[len(pass)-3:]

		row := binaryPartition(rowCode, 0, 127, lowerRow, upperRow)
		col := binaryPartition(colCode, 0, 7, lowerCol, upperCol)

		seatID := getSeatID(row, col)
		if seatID > maxID {
			maxID = seatID
		}

		seatIDs = append(seatIDs, seatID)

	}

	var seats [128][8]bool

	for r := 0; r <= 127; r++ {
		for c := 0; c <= 7; c++ {
			for _, seatID := range seatIDs {
				if seatID == getSeatID(r, c) {
					seats[r][c] = true
				}
			}
		}
	}
	var mySeatID int
	for r := 8; r < 117; r++ {
		for c := 0; c <= 7; c++ {
			if !seats[r][c] {
				mySeatID = getSeatID(r, c)
			}
		}
	}

	fmt.Printf("Exercise 1: The highest seat ID is %d.\n", maxID)
	fmt.Printf("Exercise 2: My seat ID is %d.\n", mySeatID)
}
