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

	// Intermediary lists
	updated_oxygen_values := make([]string, len(input_lines))
	updated_co2_values := make([]string, len(input_lines))

	copy(updated_oxygen_values, input_lines)
	copy(updated_co2_values, input_lines)

	// Final values
	var oxygen_generator_rating int64
	var co2_scrubber_rating int64
	var oxygen_generator_rating_binary string
	var co2_scrubber_rating_binary string

	// For each bit on the gamma and epsilon rates filter the input values
	for i := 0; i < len(gamma_rate_binary); i++ {
		// Updates the list of values checking the next position
		updated_oxygen_values = valuesWithMostCommonBitOnPosition(updated_oxygen_values, i)
		updated_co2_values = valuesWithLeastCommonBitOnPosition(updated_co2_values, i)

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

// Calculates the most common bit on the given position for the given values on the list
// Then returns a list with only the values that have the most common bit on that position
func valuesWithMostCommonBitOnPosition(value_list []string, position int) []string {
	// Number of occurrences of each bit
	zero_occurrences, one_occurrences := calculateBitFrequency(value_list, position)

	// Checks with was more frequent, tie keeps the ones with bit 1 on the given position
	if zero_occurrences > one_occurrences {
		// Removes values with bit 1 on the position
		return filterValueList(value_list, position, '1')
	} else {
		// Removes values with bit 0 on the position
		return filterValueList(value_list, position, '0')
	}
}

// Calculates the least common bit on the given position for the given values on the list
// Then returns a list with only the values that have the least common bit on that position
func valuesWithLeastCommonBitOnPosition(value_list []string, position int) []string {
	// Number of occurrences of each bit
	zero_occurrences, one_occurrences := calculateBitFrequency(value_list, position)

	// Checks with was less frequent, tie keeps the ones with bit 0 on the given position
	if zero_occurrences <= one_occurrences {
		// Removes values with bit 1 on the position
		return filterValueList(value_list, position, '1')
	} else {
		// Removes values with bit 0 on the position
		return filterValueList(value_list, position, '0')
	}
}

// Returns the frequency of each bit on a given position for the given values
func calculateBitFrequency(value_list []string, position int) (zero_occurrences int, one_occurrences int) {
	for _, bits := range value_list {
		// Counts the bit at the given position
		if bits[position] == '0' {
			zero_occurrences++
		} else {
			one_occurrences++
		}
	}

	return
}

// Filters the input list returning only the elements that don't have the given bit on the given position
func filterValueList(value_list []string, position int, bit byte) []string {
	filtered_list := []string{}

	for _, value := range value_list {
		if value[position] != bit {
			filtered_list = append(filtered_list, value)
		}
	}

	return filtered_list
}
