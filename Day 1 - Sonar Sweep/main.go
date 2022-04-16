package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open input file
	input_file, err := os.Open("sample_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input_file.Close()

	// Create scanner to read file line by line
	scanner := bufio.NewScanner(input_file)

	// Read each input line into a slice
	var values []int64

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Converts the line read into a numeric value, panics if it can't be
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Saves the resulting value in the slice
		values = append(values, value)
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// TODO: Process the input
	fmt.Println("Sample input:")
	for _, l := range values {
		fmt.Println(l)
	}
}
