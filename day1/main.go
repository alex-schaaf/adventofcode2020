package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFileLines(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	lines := readFileLines("input")
	for i, line1 := range lines {
		num1, _ := strconv.Atoi(line1)
		for j := i + 1; j < len(lines); j++ {
			num2, _ := strconv.Atoi(lines[j])
			sum := num1 + num2
			if sum == 2020 {
				fmt.Println("Success!")
				fmt.Printf("%d + %d = %d\n", num1, num2, sum)
				fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
			}
			for k := j + 1; k < len(lines); k++ {
				num3, _ := strconv.Atoi(lines[k])
				sum = num1 + num2 + num3
				if sum == 2020 {
					fmt.Println("Success!")
					fmt.Printf("%d + %d + %d = %d\n", num1, num2, num3, sum)
					fmt.Printf("%d * %d * %d = %d\n", num1, num2, num3, num1*num2*num3)
				}
			}
		}
	}
}
