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

// Bag data structure
type Bag struct {
	name      string
	children  []string
	childrenN []int
}

func parseBags(lines []string) map[string]Bag {
	bags := make(map[string]Bag)
	for _, line := range lines {
		split := strings.Split(line, " bags contain ")
		bagName := split[0]
		bag := Bag{name: bagName, children: []string{}, childrenN: []int{}}

		fmt.Println(split)

		if split[1] == "no other bags." {
			bag.children = append(bag.children, "none")
			bag.childrenN = append(bag.childrenN, 0)
			continue
		}

		for _, child := range strings.Split(split[1], ", ") {
			split = strings.Split(child, " ")
			childName := split[1] + " " + split[2]
			bag.children = append(bag.children, childName)
			childCount, _ := strconv.Atoi(split[0])
			bag.childrenN = append(bag.childrenN, childCount)
		}
		bags[bagName] = bag
	}
	return bags
}

func main() {
	lines := readFileLines("./test", "\n")
	allBags := parseBags(lines)
	fmt.Println(allBags)
	// for k, v := range allBags {
	// 	fmt.Println(k)
	// 	fmt.Println(v.children)
	// 	fmt.Println("---")
	// }
}
