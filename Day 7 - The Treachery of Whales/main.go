package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	// Create scanner to read the file
	scanner := bufio.NewScanner(input_file)

	// List of horizontal values for the crab submarines
	var horizontal_positions []int

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Split the input values
		input_values := strings.Split(line, ",")

		// Save each position as an integer
		for _, value := range input_values {
			horizontal_position, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
			}
			horizontal_positions = append(horizontal_positions, int(horizontal_position))
		}
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Calculate the optimal horizontal position
	// The mean value gets very skewed by outliers, so we use the median value instead
	optimal_position := median(horizontal_positions)

	// Calculate the fuel consumption to get there
	fuel_consumption := 0
	for _, position := range horizontal_positions {
		fuel_consumption += int(math.Abs(float64(position - optimal_position)))
	}

	fmt.Printf("Fuel consumption: %d\n", fuel_consumption)
}

// Simple median value calculating function
func median(values []int) int {
	sort.Ints(values)
	if len(values)%2 != 0 {
		return values[len(values)/2]
	} else {
		return (values[len(values)/2-1] + values[len(values)/2]) / 2
	}
}
