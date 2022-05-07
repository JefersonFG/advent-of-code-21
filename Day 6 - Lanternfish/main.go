package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	// TODO: Simulate the lanternfish evolution
}