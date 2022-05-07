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
	// Number of days for the simulation as a command line argument
	var simulation_days int
	flag.IntVar(&simulation_days, "simulation_days", 80, "number of days for the simulation")

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

	// Here just iterating through the list becomes too intensive, as the list grows exponentially
	// The ideal solution is to track how many lanternfish we have on each counter value
	// Then rotate the list to simultaneously decrease the counter of all lanternfishes
	// Then handle the spawning event

	// Stages of the lanternfish evolution, tracking the counters with its index
	// And holding the number of fishes on each stage as its value
	stages := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Maps the initial lanternfish population to the stages list
	for _, lanternfish := range lanternfishes {
		// Since we must first process the counter decrease before spawning new lanternfish
		// We need the index 0 to hold lanternfish that were at counter 0 and got decreased
		// So we map the lanternfish to one index above their counters
		stages[lanternfish+1]++
	}

	// Simulate specific number of days of lanternfish evolution
	for day := 0; day < simulation_days; day++ {
		// Rotate the list of lanternfishes' stages
		stages = append(stages[1:], stages[0])

		// Replicate number of lanternfishes that were on counter 0 to position 7, related to counter 6
		// This way that value incremented position 7 and 9, related to counters 6 and 8
		// Related to old and new lanternfish respectively
		stages[7] += stages[0]
	}

	// Calculate number of lanternfishes
	sum := 0

	for _, stage := range stages {
		sum += stage
	}

	fmt.Println("Number of lanternfishes at the end of the simulation:", sum)
}
