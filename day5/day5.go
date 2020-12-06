package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputToArray(filename string) []string {
	// Opening input.txt and storing the lines in array `input`
	var input []string

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func partOne(codes []string) (int, []int) {
	// RETURNS max ID and list of IDs
	var allIds []int

	for _, code := range codes {
		// Initializing columns and rows
		maxrow := 127
		minrow := 0
		maxcolumn := 7
		mincolumn := 0

		for i := 0; i < 10; i++ {
			// ROWS
			if string(code[i]) == "F" {
				// Lower half
				// WORKS BECAUSE GO FLOORS
				maxrow = (maxrow + minrow) / 2
			}
			if string(code[i]) == "B" {
				// Upper half
				minrow = (maxrow + minrow + 1) / 2
			}

			// COLUMNS
			if string(code[i]) == "L" {
				// Lower half
				// WORKS BECAUSE GO FLOORS
				maxcolumn = (maxcolumn + mincolumn) / 2
			}
			if string(code[i]) == "R" {
				// Upper half
				mincolumn = (maxcolumn + mincolumn + 1) / 2
			}
		}

		// Calculate the ID and append it to the list
		// NOTE: At this point maxrow and maxcolumn = the final numbers
		allIds = append(allIds, (maxrow*8 + maxcolumn))
	}

	// Find max item in `allIds`
	max := allIds[0]
	for _, i := range allIds {
		if i > max {
			max = i
		}
	}
	return max, allIds
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findID(ids []int) int {
	for _, i := range ids {
		if contains(ids, i+2) && !(contains(ids, i+1)) {
			return i + 1
		}
		if contains(ids, i-2) && !(contains(ids, i-1)) {
			return i - 1
		}
	}
	return -1
}

func main() {
	array := inputToArray("input.txt")

	maxID, allIds := partOne(array)

	fmt.Println("Part1:", maxID)
	fmt.Println("Part2:", findID(allIds))

}
