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
	// Open input file
	var input_file_path string
	flag.StringVar(&input_file_path, "input_path", "input.txt", "path to the input file, with one command on each line")
	flag.Parse()
	input_file, err := os.Open(input_file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer input_file.Close()

	// Create scanner to read file line by line
	scanner := bufio.NewScanner(input_file)

	// Read each input line into a slice of commands and one of offsets
	var commands []string
	var offsets []int64

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Splits the line between the command and the value
		input_values := strings.Split(line, " ")

		command, offset := input_values[0], input_values[1]

		converted_offset, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Saves the parsed input
		commands = append(commands, command)
		offsets = append(offsets, converted_offset)
	}

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// ### Part 1 ###
	// Submarine position
	horizontal_position := 0
	depth := 0

	// Traverse the list of instructions operating on the position values
	for i := 0; i < len(commands); i++ {
		// Checks for current command
		// Assumes only valid commands
		value := int(offsets[i])
		switch commands[i] {
		case "forward":
			horizontal_position += value
		case "down":
			depth += value
		default:
			depth -= value
		}
	}

	// Prints the result
	fmt.Printf("Final position value: %d\n", horizontal_position*depth)

	// ### Part 2 ###
	// Reset previous values, add aim to the mix
	horizontal_position = 0
	depth = 0
	aim := 0

	// Traverse the list of instructions operating on the position values
	for i := 0; i < len(commands); i++ {
		// Checks for current command
		// Assumes only valid commands
		value := int(offsets[i])
		switch commands[i] {
		case "forward":
			horizontal_position += value
			depth += aim * value
		case "down":
			aim += value
		default:
			aim -= value
		}
	}

	// Prints the result
	fmt.Printf("Final position value with aim covered: %d\n", horizontal_position*depth)
}
