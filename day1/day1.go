package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Opening input.txt and storing the lines in array `input`
	var input []int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	// For each number in list, find number that adds to equal 2020, and search list
	for _, v := range input {
		dif := 2020 - v

		// PART 1 SOLUTION - Search for dif in array
		/*for _, d := range input {
			if d == dif {
				fmt.Println(v, "+", dif, "= 2020")
				fmt.Println(v * dif)
				os.Exit(3)
			}
		}*/

		// PART 2 SLUTION - Search for two numbers whose sum is dif
		for _, num1 := range input {
			for _, num2 := range input {
				if num1+num2 == dif {
					fmt.Println(v, "+", num1, "+", num2, "= 2020")
					fmt.Println(v * num1 * num2)
					os.Exit(3)
				}
			}
		}
	}
}
