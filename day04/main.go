package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var ecls = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

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

func inRange(val string, min int, max int) int {
	if len(val) == 0 {
		return 0
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	if min <= valInt && valInt <= max {
		return 1
	}
	return 0
}

func exercise2(passports []map[string]string) int {
	var validPassports = 0
	for _, passport := range passports {
		var validFields = 0
		// byr iyr eyr
		validFields += inRange(passport["byr"], 1920, 2002)
		validFields += inRange(passport["iyr"], 2010, 2020)
		validFields += inRange(passport["eyr"], 2020, 2030)
		// hgt
		hgt := passport["hgt"]
		hgtMatched, err := regexp.MatchString(`[0-9]{3}|[0-9]{2}[a-z]{2}`, hgt)
		if err != nil {
			log.Fatal(err)
		}
		if hgtMatched {
			unit := hgt[len(hgt)-2:]
			val := hgt[:len(hgt)-2]
			if unit == "cm" {
				validFields += inRange(val, 150, 193)
			} else if unit == "in" {
				validFields += inRange(val, 59, 76)
			}
		}
		// hcl
		hclMatched, err := regexp.MatchString(`#[0-9a-z]{6}`, passport["hcl"])
		if err != nil {
			log.Fatal(err)
		}
		if hclMatched {
			validFields++
		}
		// ecl
		ecl := passport["ecl"]
		for _, validECL := range ecls {
			if ecl == validECL {
				validFields++
			}
		}
		// pid
		pidMatched, err := regexp.MatchString(`[0-9]{9}`, passport["pid"])
		if err != nil {
			log.Fatal(err)
		}
		if pidMatched {
			validFields++
		}

		if validFields == 7 {
			validPassports++
		}
	}

	return validPassports - 1
}

func main() {
	passportStrings := readPassportStrings("./input")
	passports := parsePassportStrings(passportStrings)

	for _, passport := range passports {
		fmt.Println(passport)
	}

	sol1 := exercise1(passports)
	sol2 := exercise2(passports)
	fmt.Printf("Exercise 1: %d valid passports detected.\n", sol1)
	fmt.Printf("Exercise 2: %d valid passports detected.\n", sol2)
}
