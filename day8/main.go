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

// acc <arg> 		increases a single global variable accumulator
// 					after acc, the instruction immediately below is executed next
// jmp <offset>		jump above or below instruction with offset
// nop				No OPeration - it does nothing. next instruction is executed next

// Instruction data structure
type Instruction struct {
	op       string
	arg      int
	executed int
}

func parseInstructions(lines []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		op := split[0]
		arg, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			log.Fatal(err)
		}
		instruct := Instruction{op, arg, 0}

		instructions = append(instructions, instruct)
	}
	return instructions
}

var accumulator int

func main() {
	lines := readFileLines("./input", "\n")
	instructions := parseInstructions(lines)

	executing := true
	index := 0

	for executing {
		instruction := &instructions[index]
		fmt.Println(index, instruction)
		if instruction.executed > 0 {
			executing = false
			fmt.Printf("Exercise 1: The accumulator value is %d.\n", accumulator)

		}
		if instruction.op == "nop" {
			instruction.executed++
			index++
			continue
		} else if instruction.op == "acc" {
			accumulator += instruction.arg
			instruction.executed++
			index++
			continue
		} else if instruction.op == "jmp" {
			instruction.executed++
			index += instruction.arg
			continue
		}
	}

}
