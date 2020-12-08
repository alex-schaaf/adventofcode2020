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

func switchInstruction(instruciton *Instruction) bool {
	if instruciton.op == "nop" {
		instruciton.op = "jmp"
		return true
	} else if instruciton.op == "jmp" {
		instruciton.op = "nop"
		return true
	}
	return false
}

func mutateInstructionSet(instructions []Instruction, previous int) []Instruction {
	var newInstructions []Instruction
	switched := false
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		if i > previous && !switched {
			if switchInstruction(&instruction) {
				switched = true
				previous = i
			}
		}
		newInstructions = append(newInstructions, instruction)
	}
	return newInstructions
}

func main() {
	lines := readFileLines("./input", "\n")
	instructions := parseInstructions(lines)
	previous := -1

	for i := 0; i < len(instructions); i++ {
		var accumulator int
		if instructions[i].op == "acc" {
			continue
		}

		var mutatedInstructions []Instruction
		switched := false
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			if i > previous && !switched {
				if switchInstruction(&instruction) {
					switched = true
					previous = i
				}
			}
			mutatedInstructions = append(mutatedInstructions, instruction)
		}

		executing := true
		index := 0
		for executing {
			if index > len(lines)-1 {
				fmt.Printf("The program has terminated successfully. Accumulator value is %d.", accumulator)
				break
			}
			instruction := &mutatedInstructions[index]
			// fmt.Println(index, instruction)
			if instruction.executed > 0 {
				executing = false
				// fmt.Println("The program has terminated due to second execution attempt.")
				// fmt.Printf("Exercise 1: The accumulator value is %d.\n", accumulator)
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

}
