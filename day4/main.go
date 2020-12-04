package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func readPassportStrings(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n\n")
}

func parsePassportStrings(passportStrings []string) []map[string]string {
	var passports []map[string]string

	for _, line := range passportStrings {
		passport := make(map[string]string)
		line := strings.Replace(line, "\n", " ", -1)
		entries := strings.Split(line, " ")
		for _, entry := range entries {
			passport[entry[:3]] = entry[4:]
		}
		passports = append(passports, passport)

	}
	return passports
}

func exercise1(passports []map[string]string) int {
	var validPasswords = 0
	for _, passport := range passports {
		var isValid = 1
		for _, field := range requiredFields {
			if _, ok := passport[field]; !ok {
				isValid = 0
			}
		}
		validPasswords += isValid
	}
	return validPasswords
}

func main() {
	passportStrings := readPassportStrings("./test")
	passports := parsePassportStrings(passportStrings)

	for _, passport := range passports {
		fmt.Println(passport)
	}

	sol1 := exercise1(passports)
	fmt.Printf("Exercise 1: %d valid passports detected.\n", sol1)
}
