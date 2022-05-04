package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	// Keeps track of the highest value that appeared to create a table of that size
	field_size := 0

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Parse the input values
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d, %d", &x1, &y1, &x2, &y2)

		// Saves the maximum size
		if x1 > field_size {
			field_size = x1
		}
		if y1 > field_size {
			field_size = y1
		}
		if x2 > field_size {
			field_size = x2
		}
		if y2 > field_size {
			field_size = y2
		}

		// TODO: Determine the lines
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// TODO: Generate the table with the lines and count the numbers higher than one
	fmt.Printf("Field size: %d\n", field_size)
}
