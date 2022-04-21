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

	// List of lines from the input, necessary for part 2
	var input_lines []string

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Saves the line on the list of lines
		input_lines = append(input_lines, line)

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

	// ### Part 2 ###
	// To calculate the oxygen generator and CO2 scrubber ratings
	// we need the most common and least common bits on each position
	// We then use them as prefixes for the actual values until we find only one value left
	// The bits are already available in the gamma and epsilon rates, so we use those

	// TODO: This approach uses the global most and least common bits
	// It should recalculate for each sublist
	// Rework the logic

	// Final values
	var oxygen_generator_rating int64
	var co2_scrubber_rating int64
	var oxygen_generator_rating_binary string
	var co2_scrubber_rating_binary string

	// For each bit on the gamma and epsilon rates filter the input values
	for i := 0; i <= len(gamma_rate_binary); i++ {
		// Intermediary lists
		var updated_oxygen_values []string
		var updated_co2_values []string

		// Current prefixes
		gamma_prefix := gamma_rate_binary[:i]
		epsilon_prefix := epsilon_rate_binary[:i]

		// Check each input value and save the ones that have the valid prefix
		for j := 0; j < len(input_lines); j++ {
			if strings.HasPrefix(input_lines[j], gamma_prefix) {
				updated_oxygen_values = append(updated_oxygen_values, input_lines[j])
			}
			if strings.HasPrefix(input_lines[j], epsilon_prefix) {
				updated_co2_values = append(updated_co2_values, input_lines[j])
			}
		}

		// Checks the list of values, if unitary save the final value
		if len(updated_oxygen_values) == 1 {
			oxygen_generator_rating_binary = updated_oxygen_values[0]
		}
		if len(updated_co2_values) == 1 {
			co2_scrubber_rating_binary = updated_co2_values[0]
		}
	}

	oxygen_generator_rating, err = strconv.ParseInt(oxygen_generator_rating_binary, 2, 64)
	if err != nil {
		panic(err)
	}

	co2_scrubber_rating, err = strconv.ParseInt(co2_scrubber_rating_binary, 2, 64)
	if err != nil {
		panic(err)
	}

	// Print the resulting values
	fmt.Printf("Resulting oxygen generator rating: %d\n", oxygen_generator_rating)
	fmt.Printf("Resulting co2 scrubber rating: %d\n", co2_scrubber_rating)
	fmt.Printf("Submarine power consumption: %d\n", oxygen_generator_rating*co2_scrubber_rating)
}
