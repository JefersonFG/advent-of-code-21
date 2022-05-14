package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	original_mappings = make(map[int][]string)
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

	// Set the original mapping, so we can translate the values
	original_mappings = map[int][]string{
		0: {"a", "b", "c", "e", "f", "g"},
		1: {"c", "f"},
		2: {"a", "c", "d", "e", "g"},
		3: {"a", "c", "d", "f", "g"},
		4: {"b", "c", "d", "f"},
		5: {"a", "b", "d", "f", "g"},
		6: {"a", "b", "d", "e", "f", "g"},
		7: {"a", "c", "f"},
		8: {"a", "b", "c", "d", "e", "f", "g"},
		9: {"a", "b", "c", "d", "f", "g"},
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
		current_mapping := map[string][]string{
			"a": {"a", "b", "c", "d", "e", "f", "g"},
			"b": {"a", "b", "c", "d", "e", "f", "g"},
			"c": {"a", "b", "c", "d", "e", "f", "g"},
			"d": {"a", "b", "c", "d", "e", "f", "g"},
			"e": {"a", "b", "c", "d", "e", "f", "g"},
			"f": {"a", "b", "c", "d", "e", "f", "g"},
			"g": {"a", "b", "c", "d", "e", "f", "g"},
		}

		for _, unique_signal_pattern := range unique_signal_patterns[i] {
			// Look into numbers 1, 4 and 7 and update the possible mapping
			if len(unique_signal_pattern) == 2 {
				update_current_mapping(current_mapping, unique_signal_pattern, []int{1})
			} else if len(unique_signal_pattern) == 4 {
				update_current_mapping(current_mapping, unique_signal_pattern, []int{4})
			} else if len(unique_signal_pattern) == 3 {
				update_current_mapping(current_mapping, unique_signal_pattern, []int{7})
			}

			// TODO: We must find out which of the possible numbers it is
			// The code below doesn't work
			// To find out we must check the mapping at this point and see if the current unique pattern
			// Could be matched with all of these numbers, until we know which it is and update the current_mapping
			if len(unique_signal_pattern) == 6 {
				update_current_mapping(current_mapping, unique_signal_pattern, []int{0, 6, 9})
			} else if len(unique_signal_pattern) == 5 {
				update_current_mapping(current_mapping, unique_signal_pattern, []int{2, 3, 5})
			}
		}

		// Accumulate
		output_display_sum += output_display_value
	}

	fmt.Printf("Sum of all the output display values: %d\n", output_display_sum)
}

// Function for updating the current mapping with the hints the current unique signal pattern gives
// Such as that the original segment c might only be mapped to segments b and e
// Because the unique pattern contains these on the number 1, which covers the original segment c
func update_current_mapping(current_mapping map[string][]string, unique_signal_pattern string, possible_numbers []int) {
	// Repeat for every possible number, determined previously by the length of the unique signal pattern
	for _, possible_number := range possible_numbers {
		segments := strings.Split(unique_signal_pattern, "")
		original_mapping := original_mappings[possible_number]

		// For each segment of the original mapping change the values of the current mapping
		for _, original_segment := range original_mapping {
			var updated_mapping []string
			// Look into each segment of the current unique pattern
			for _, current_segment := range segments {
				// Traverse the current mapping for the original segment
				for _, possible_segment := range current_mapping[original_segment] {
					// If the value on the current mapping is the same as the unique pattern we keep it
					// Otherwise we discard it
					if possible_segment == current_segment {
						updated_mapping = append(updated_mapping, possible_segment)
					}
				}
			}

			current_mapping[original_segment] = updated_mapping
		}
	}
}
