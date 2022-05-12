package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	// List of output values from the input, each containing the segments turned on
	var segment_outputs [][]string

	// List of all the unique signal patterns from the input
	var unique_signal_patterns [][]string

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Separate the input into the unique signal patterns and the output digits
		input_sections := strings.Split(line, "|")
		unique_signal_pattern, segment_output := input_sections[0], input_sections[1]

		// Save the output segments to the list
		segment_outputs = append(segment_outputs, strings.Split(segment_output, " "))

		// Same for the unique signal patterns
		unique_signal_patterns = append(unique_signal_patterns, strings.Split(unique_signal_pattern, " "))
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Counter for the occurrences of values 1, 4, 7 and 8
	unique_segments_counter := 0

	// Traverse the list of outputs searching for unique number of seguments
	for _, segment_output := range segment_outputs {
		for _, segment_value := range segment_output {
			if len(segment_value) == 2 || len(segment_value) == 4 || len(segment_value) == 3 || len(segment_value) == 7 {
				unique_segments_counter++
			}
		}
	}

	fmt.Printf("Number of occurences of the values 1, 4, 7 and 8 on the output seguments: %d\n", unique_segments_counter)

	// Sum of all of the output display values
	output_display_sum := 0

	// Original mapping, so we can translate the values
	original_mapping := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	// For each line find the output display value and accumulate
	for i := 0; i < len(unique_signal_patterns); i++ {
		// Final value
		output_display_value := 0

		// The first step is to determine the correct mapping between the original values for each segment
		// To the actual values we're getting for each segment
		// For the numbers 1, 4 and 7 we can filter which values can and can't be the ones for each segment
		// So we start with those

		// Current mapping with all possible segments from original to current values
		// Once the mapping is 1:1 we can revert the map to get back the actual numbers
		current_mapping := map[string]string{
			"a": "abcdefg",
			"b": "abcdefg",
			"c": "abcdefg",
			"d": "abcdefg",
			"e": "abcdefg",
			"f": "abcdefg",
			"g": "abcdefg",
		}

		// TODO: Look into numbers 1, 4 and 7 and update the possible mapping

		// Accumulate
		output_display_sum += output_display_value
	}

	fmt.Printf("Sum of all the output display values: %d\n", output_display_sum)
}
