package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputToArray() []string {
	// Opening input.txt and storing the lines in array `input`
	var input []string

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func fixArray(input []string) []string {
	// NOTE: Blank line is considered == ""

	space1 := 0
	space2 := 0
	var newArray []string

	// Creating NewArray
	// i = index
	// c = contents
	for i, c := range input {
		if c == "" {
			space2 = i

			if space1 != 0 {
				newEntry := ""
				for v := space1 + 1; v < space2; v++ {
					newEntry = newEntry + input[v]
				}
				newArray = append(newArray, newEntry)
			} else {
				newEntry := ""
				for v := space1; v < space2; v++ {
					newEntry = newEntry + input[v]
				}
				newArray = append(newArray, newEntry)
			}
			space1 = space2
		}
	}

	return newArray
}

func totalValid(part int, array []string) int {
	if part == 1 {
		totalValid := 0
		for _, i := range array {
			if strings.Contains(i, "byr") && strings.Contains(i, "iyr") && strings.Contains(i, "eyr") && strings.Contains(i, "hgt") && strings.Contains(i, "hcl") && strings.Contains(i, "ecl") && strings.Contains(i, "pid") {
				totalValid++
			}
		}
		return totalValid
	}

	// ELSE
	totalValid := 0

	for _, i := range array {
		iValid := 0
		// Check birth year (byr)
		if strings.Contains(i, "byr") {
			birthyear, _ := strconv.Atoi(i[strings.Index(i, "byr")+4 : strings.Index(i, "byr")+8])

			fmt.Println("Birthyear:", birthyear)

			if birthyear >= 1920 && birthyear <= 2002 {
				iValid++
			}
		}

		// Check Issue year (iyr)
		if strings.Contains(i, "iyr") {
			issueyear, _ := strconv.Atoi(i[strings.Index(i, "iyr")+4 : strings.Index(i, "iyr")+8])

			fmt.Println("Issue year:", issueyear)

			if issueyear >= 2010 && issueyear <= 2020 {
				iValid++
			}
		}

		// Check Expiration year (eyr)
		if strings.Contains(i, "eyr") {
			expyear, _ := strconv.Atoi(i[strings.Index(i, "eyr")+4 : strings.Index(i, "eyr")+8])

			fmt.Println("expiration year:", expyear)

			if expyear >= 2020 && expyear <= 2030 {
				iValid++
			}
		}

		// Check height (hgt)
		if strings.Contains(i, "hgt") {
			// Check for cm or in
			if strings.Contains(i, "in") {
				if strings.Index(i, "hgt")+8 < len(i) {
					// Grab anything between the `:` and `in`
					height, err := strconv.Atoi(i[strings.Index(i, "hgt")+4 : strings.Index(i, "in")])
					if err == nil {
						fmt.Println("height:", height)

						if height >= 59 && height <= 76 {
							iValid++
						}
					}
				}
			}
			if strings.Contains(i, "cm") {
				if strings.Index(i, "hgt")+9 < len(i) {
					// Grab anything between the `:` and `cm`
					height, err := strconv.Atoi(i[strings.Index(i, "hgt")+4 : strings.Index(i, "cm")])
					if err == nil {
						fmt.Println("height:", height)

						if height >= 150 && height <= 193 {
							iValid++
						}
					}
				}
			}
		}

		// Check hair color (hcl)
		if strings.Contains(i, "hcl") {
			// Grab 7 characters after `hcl:` and
			// CHECK THIS IF IT FAILS-------------------------------------------------------------------------------------
			if strings.Index(i, "hcl")+11 < len(i) {
				hex := i[strings.Index(i, "hcl")+4 : strings.Index(i, "hcl")+11]

				fmt.Println("haircolor:", hex)

				if string(hex[0]) == "#" {
					hexValid := 0

					for i := 1; i < 7; i++ {
						if (string(hex[i]) <= "9" && string(hex[i]) >= "0") || (string(hex[i]) <= "A" && string(hex[i]) >= "F") {
							hexValid++
						}
					}
					if hexValid == 6 {
						iValid++
					}
				}
			}
		}

		// Check eye color (ecl)
		if strings.Contains(i, "ecl") {
			ec := i[strings.Index(i, "ecl")+4 : strings.Index(i, "ecl")+7]

			fmt.Println("eyecolor:", ec)

			if ec == "amb" || ec == "blu" || ec == "brn" || ec == "gry" || ec == "grn" || ec == "hzl" || ec == "oth" {
				iValid++
			}
		}

		// Check passport id (pid)
		if strings.Contains(i, "pid") {
			pid := i[strings.Index(i, "pid")+4 : strings.Index(i, "pid")+13]

			fmt.Println("passport id:", pid)

			if _, err := strconv.Atoi(pid); err == nil {
				iValid++
			}
		}
		if iValid == 7 {
			totalValid++
		}
	}
	return totalValid
}

func main() {
	pass := inputToArray()
	//partOneAns := totalValid(1, fixArray(pass))
	//fmt.Println(partOneAns)
	partTwoAns := totalValid(2, fixArray(pass))
	fmt.Println(partTwoAns)
}
