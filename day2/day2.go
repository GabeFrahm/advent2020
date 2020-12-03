package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Returns first num, second num, search character, and code
func parseItem(s string) (int, int, string, string) {
	// Variables to return
	var num1, num2 int
	var char, code string

	// Variables for use in function
	var dashPos, spacePos int
	foundSpace := false

	// Find ints and strings
	for pos, c := range s {
		sc := string(c)
		if sc == "-" {
			dashPos = pos
			num1, _ = strconv.Atoi(s[0:dashPos]) // num1 will always be before the dash
		}
		if sc == " " {
			if foundSpace == false {
				spacePos = pos
				num2, _ = strconv.Atoi(s[dashPos+1 : spacePos]) //num2 will always be before the space and after the dash
				char = string(s[spacePos+1])                    // char will always be after the space
				foundSpace = true
			}
		}
		if sc == ":" {
			code = s[pos+2 : len(s)] // code will always be 2 spaces after the colon because of the space
		}
	}

	// Return the values
	return num1, num2, char, code
}

func partOne(input []string) int {
	numValid := 0
	for _, item := range input {
		num1, num2, char, code := parseItem(item)
		count := strings.Count(code, char)
		if count >= num1 && count <= num2 {
			numValid++
		}
	}
	return numValid
}

func partTwo(input []string) int {
	numValid := 0
	for _, item := range input {
		num1, num2, char, code := parseItem(item)
		if code[num1-1] != code[num2-1] {
			if string(code[num1-1]) == char || string(code[num2-1]) == char {
				numValid++
			}
		}
	}
	return numValid
}

func main() {
	// Opening input.txt and storing the lines in array `input`
	var input []string

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
}
