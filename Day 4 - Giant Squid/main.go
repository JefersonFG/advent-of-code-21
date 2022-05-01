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
	line_matches, column_matches [5]int    // Lists for value matches on lines and columns, if the value reaches 5 the table won
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
	var bingo_tables []*bingo

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
			bingo_tables = append(bingo_tables, &bingo{})
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

	// Hash for storing all number occurrences; if a number was drawn it is on the hash
	// Comes at a faster lookup than a simple list
	drawn_hash := make(map[int]bool)

	// Boolean indicating if the game was won to finish the loop
	game_won := false

	// Final score
	score := 0

	// For the game loop we start drawing the number
	for _, random_value := range random_values {
		// Add the value to the hash
		drawn_hash[random_value] = true

		// Search for the number on each table
		for _, bingo_table := range bingo_tables {
			// Tables have fixed size
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					// If the value has been found update both arrays
					if bingo_table.table[i][j] == random_value {
						// If one of the arrays reach value 5 the table has won
						if bingo_table.line_matches[i]++; bingo_table.line_matches[i] >= 5 {
							score = calculateScore(bingo_table, random_value, drawn_hash)
							game_won = true
						}
						if bingo_table.column_matches[j]++; bingo_table.column_matches[j] >= 5 {
							score = calculateScore(bingo_table, random_value, drawn_hash)
							game_won = true
						}
					}

					if game_won {
						break
					}
				}

				if game_won {
					break
				}
			}

			if game_won {
				break
			}
		}

		if game_won {
			break
		}
	}

	// Shows the final score
	fmt.Printf("Final score: %d\n", score)
}

func calculateScore(bingo_table *bingo, random_value int, drawn_hash map[int]bool) (score int) {
	// Calculate the sum of all unmarked numbers
	sum := 0

	// For every value on the table check if it exists on the hash
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// If it doesn't exist accumulate the value
			_, ok := drawn_hash[bingo_table.table[i][j]]
			if !ok {
				sum += bingo_table.table[i][j]
			}
		}
	}

	// Calculates and returns the score
	score = sum * random_value
	return
}
