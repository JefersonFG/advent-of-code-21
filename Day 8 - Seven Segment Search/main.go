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

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Separate the input into the unique signal patterns and the output digits
		input_sections := strings.Split(line, "|")
		segment_output := input_sections[1]

		// Save the output segments to the list
		segment_outputs = append(segment_outputs, strings.Split(segment_output, " "))
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
}
