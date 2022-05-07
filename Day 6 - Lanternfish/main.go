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

// Constants
const (
	simulation_days         = 80
	lanternfish_reset_value = 6
	new_lanternfish_value   = 8
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

	// Create scanner to read the file
	scanner := bufio.NewScanner(input_file)

	// List of lanternfish counters
	var lanternfishes []int

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Split the input values
		input_values := strings.Split(line, ",")

		// Save each lanternfish as an integer
		for _, value := range input_values {
			lanterfish, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}
			lanternfishes = append(lanternfishes, int(lanterfish))
		}
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Simulate specific number of days of lanternfish evolution
	for day := 0; day < simulation_days; day++ {
		// Saves the current number of lanternfishes so as to not calculate the evolution of new ones
		school_size := len(lanternfishes)

		// Calculate the evolution for each lanternfish that started the day
		for i := 0; i < school_size; i++ {
			// If the counter is zero spawn a new lanternfish and reset the counter
			if lanternfishes[i] == 0 {
				// Reset the counter
				lanternfishes[i] = lanternfish_reset_value

				// Add new fish
				lanternfishes = append(lanternfishes, new_lanternfish_value)
			} else {
				// Otherwise just decrease the counter
				lanternfishes[i]--
			}
		}
	}

	// Print number of lanternfishes
	fmt.Println("Number of lanternfishes at the end of the simulation:", len(lanternfishes))
}
