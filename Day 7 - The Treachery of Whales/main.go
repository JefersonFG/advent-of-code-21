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
	// Parameter for calculating with a linear fuel consumption rate or not
	var linear_consumption bool
	flag.BoolVar(&linear_consumption, "linear_consumption", true, "calculate with a linear consumption rate or increasing consumption rate")

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

	fuel_consumption := 0

	if linear_consumption {
		// In the case of linear fuel consumption the mean value gets very skewed by outliers
		// So we use the median value instead
		optimal_position := median(horizontal_positions)

		for _, position := range horizontal_positions {
			fuel_consumption += int(math.Abs(float64(position - optimal_position)))
		}
	} else {
		// In the case of non linear fuel consumption we have two candidates for the optimal position
		// So we check both for the lower fuel consumption
		// Also here the outliers are important as they mean more fuel consumption, so we use the mean
		ceiled_position := meanCeiled(horizontal_positions)
		floored_position := meanFloored(horizontal_positions)

		ceiled_consumption := 0
		floored_consumption := 0

		for _, position := range horizontal_positions {
			distance := int(math.Abs(float64(position - ceiled_position)))
			ceiled_consumption += (distance * (distance + 1)) / 2
			distance = int(math.Abs(float64(position - floored_position)))
			floored_consumption += (distance * (distance + 1)) / 2
		}

		if ceiled_consumption < floored_consumption {
			fuel_consumption = ceiled_consumption
		} else {
			fuel_consumption = floored_consumption
		}
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

// Function for calculating the mean value ceiling the result
func meanCeiled(values []int) int {
	sum := 0
	num_values := float64(len(values))
	for _, value := range values {
		sum += value
	}
	return int(math.Ceil(float64(sum) / num_values))
}

// Function for calculating the mean value flooring the result
func meanFloored(values []int) int {
	sum := 0
	num_values := float64(len(values))
	for _, value := range values {
		sum += value
	}
	return int(math.Floor(float64(sum) / num_values))
}
