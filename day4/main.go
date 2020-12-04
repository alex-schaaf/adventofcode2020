package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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

func main() {
	passportStrings := readPassportStrings("./test")
	passports := parsePassportStrings(passportStrings)

	for _, passport := range passports {
		fmt.Println(passport)
	}

}
