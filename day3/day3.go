package main

import (
	"bufio"
	"fmt"
	"os"
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

func findCollisions(input []string, across int, down int) int {
	// NOTE: input[0][X] RETURNS ASCII VALUE
	// 46 == Open ; 35 == Tree
	length := len(input[0])
	totalHits := 0
	counter := 0
	pos := 0

	for row, _ := range input {
		if row != 0 && row%down == 0 {
			// Finding cyclical position
			counter = counter + across
			pos = counter % length

			// If current position is Tree, add to hits
			if (input[row])[pos] == 35 {
				totalHits++
			}
		}
	}

	return totalHits
}

func main() {
	input := inputToArray()
	//fmt.Println(findCollisions(input, 3, 1))
	fmt.Println(findCollisions(input, 1, 1) * findCollisions(input, 3, 1) * findCollisions(input, 5, 1) * findCollisions(input, 7, 1) * findCollisions(input, 1, 2))
}
