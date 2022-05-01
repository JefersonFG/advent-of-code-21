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

// Structure to hold information related to one table for the bingo
type bingo struct {
	table                        [5][5]int // The table itself, holding only the values for each field
	line_matches, column_matches []int     // Lists for value matches on lines and columns, if size reaches 5 the table won
}

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

	// The first line contains the list of random values to be drawn during the game of bingo
	var random_values []int

	// Following lines define a series of tables
	var bingo_tables []bingo

	// Table line being currently read, used to indicate when a new table must be created
	current_line := 0

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Checks if the list of random values is empty, if so read the first line into it
		if len(random_values) == 0 {
			// Splits the input line into a slice of strings, each string holding a number
			input_values := strings.Split(line, ",")

			// For each number convert it and place on the random_values slice
			for _, value := range input_values {
				random_value, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					panic(err)
				}
				random_values = append(random_values, int(random_value))
			}
			continue
		}

		// If line is not empty it is part of a table

		// If line currently being read is the first then create a new table
		if current_line == 0 {
			bingo_tables = append(bingo_tables, bingo{})
		}

		// Read the current line of the last table

		// Splits the input line into a slice of strings, each string holding a number
		input_values := strings.Fields(line)

		// For each number convert it and place on the table row
		for index, value := range input_values {
			table_value, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}
			bingo_tables[len(bingo_tables)-1].table[current_line][index] = int(table_value)
		}

		// Increments the current line and resets it once it reaches the end
		if current_line++; current_line > 4 {
			current_line = 0
		}
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// TODO: Determine the winning value
	fmt.Println(bingo_tables)
}
