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

// const empty = 76
// const floor = 46
// const occupied = 35

func countNeighbors(seats [][]string, y int, x int) int {
	var occupiedNeighbors int
	ny := len(seats)
	nx := len(seats[0])

	var coords = [][]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
		{1, -1}, {-1, -1}, {-1, 1}, {1, 1}}

	for _, coord := range coords {
		dy, dx := coord[0], coord[1]
		if y+dy < 0 {
			continue
		} else if y+dy >= ny {
			continue
		}
		if x+dx < 0 {
			continue
		} else if x+dx >= nx {
			continue
		}
		seat := seats[y+dy][x+dx]
		if seat == "#" {
			occupiedNeighbors++
		}
	}
	return occupiedNeighbors
}

func main() {
	lines := readFileLines("./test", "\n")
	var seats [][]string
	for y, row := range lines {
		seats = append(seats, []string{})
		for _, seat := range row {
			seats[y] = append(seats[y], string(seat))
		}
	}
	ny := len(seats)
	nx := len(seats[0])

	var nChanges = 1
	var i int

	for nChanges > 0 {
		var newSeats = make([][]string, nx, ny)
		copy(newSeats, seats)
		// fmt.Printf("\nIteration %d\n", i)
		i++
		nChanges = 0
		for y := 0; y < ny; y++ {
			// fmt.Printf("%v\t", seats[y])
			for x := 0; x < nx; x++ {

				nNeighbors := countNeighbors(seats, y, x)
				// fmt.Printf("%d", nNeighbors)
				seat := seats[y][x]
				if seat == "." {
					continue
				}
				if seat == "L" {
					if nNeighbors == 0 {
						// -> occupied
						newSeats[y][x] = "#"
						nChanges++
						continue
					}
				} else if seat == "#" {
					if nNeighbors >= 4 {
						// -> empty
						newSeats[y][x] = "L"
						nChanges++
						continue
					}
				}

			}
			// fmt.Printf("\n")
		}
		copy(seats, newSeats)

		// fmt.Println(nChanges)
	}
	var occupiedSeats int
	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			if seats[y][x] == "#" {
				occupiedSeats++
			}
		}
	}
	fmt.Println(occupiedSeats)

}
