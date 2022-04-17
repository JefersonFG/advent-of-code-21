package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open input file
	var input_file_path string
	flag.StringVar(&input_file_path, "input_path", "sample_input.txt", "path to the input file, with one command on each line")
	flag.Parse()
	input_file, err := os.Open(input_file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer input_file.Close()

	// Create scanner to read file line by line
	scanner := bufio.NewScanner(input_file)

	// Read each input line as a list of strings
	// Each position of the binary is mapped to a hash table
	// The hash table counts the occurrences of 0s and 1s on each position
	var positions []map[string]int

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Splits the line into a list of strings
		input_values := strings.Split(line, "")

		// Initialize the list of maps
		if len(positions) == 0 {
			for range input_values {
				positions = append(positions, make(map[string]int))
			}
		}

		// For each input value update the map
		for index, value := range input_values {
			positions[index][value] += 1
		}
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// After parsing all of the input values calculate the gamma and epsilon rates
	gamma_rate_binary := ""
	epsilon_rate_binary := ""

	for _, frequency := range positions {
		if frequency["0"] > frequency["1"] {
			gamma_rate_binary += "0"
			epsilon_rate_binary += "1"
		} else {
			gamma_rate_binary += "1"
			epsilon_rate_binary += "0"
		}
	}

	gamma_rate, err := strconv.ParseInt(gamma_rate_binary, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon_rate, err := strconv.ParseInt(epsilon_rate_binary, 2, 64)
	if err != nil {
		panic(err)
	}

	// Print the resulting values
	fmt.Printf("Resulting gamma rate: %d\n", gamma_rate)
	fmt.Printf("Resulting epsilon rate: %d\n", epsilon_rate)
	fmt.Printf("Submarine power consumption: %d\n", gamma_rate*epsilon_rate)
}
