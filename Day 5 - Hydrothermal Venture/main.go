package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type line_segment struct {
	x1, y1, x2, y2 int
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

	// Keeps track of the highest value that appeared to create a table of that size
	field_size := 0

	// Saves a list of horizontal and vertical lines
	var horizontal_and_vertical_lines []line_segment

	// For part 2 we want the diagonal lines as well
	var diagonal_lines []line_segment

	for scanner.Scan() {
		// Reads a line and check that it isn't empty
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Parse the input values
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

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

		// Saves only horizontal and vertical lines
		if x1 == x2 || y1 == y2 {
			// Saves the lines such that following them on the field can be done from x1 to x2 and from y1 to y2
			if (y1 == y2 && x2 < x1) || (x1 == x2 && y2 < y1) {
				x1, x2, y1, y2 = x2, x1, y2, y1
			}
			horizontal_and_vertical_lines = append(horizontal_and_vertical_lines, line_segment{x1, y1, x2, y2})
		} else {
			// Saves the diagonal lines making sure that x1 <= x2
			if x2 < x1 {
				x1, x2, y1, y2 = x2, x1, y2, y1
			}
			diagonal_lines = append(diagonal_lines, line_segment{x1, y1, x2, y2})
		}
	}

	// Input values start at zero and go to the maximum passed, so we increment the size to cover all on the table indexes
	field_size++

	// Checks for scanner errors, panics if one is found
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create the table and initialize it with zeroes
	field := make([][]int, field_size)
	for i := 0; i < field_size; i++ {
		field[i] = make([]int, field_size)
		for j := 0; j < field_size; j++ {
			field[i][j] = 0
		}
	}

	// Go through each line marking them on the table
	for _, line := range horizontal_and_vertical_lines {
		// Horizontal line
		if line.x1 == line.x2 {
			for j := line.y1; j <= line.y2; j++ {
				field[j][line.x1]++
			}
		}

		// Vertical line
		if line.y1 == line.y2 {
			for i := line.x1; i <= line.x2; i++ {
				field[line.y1][i]++
			}
		}
	}

	// Count how many occurrences of crossing lines there are on the table
	count := 0

	for i := 0; i < field_size; i++ {
		for j := 0; j < field_size; j++ {
			if field[i][j] >= 2 {
				count++
			}
		}
	}

	fmt.Println("Points with horizontal and vertical lines:", count)

	// Add the diagonal lines and calculate again

	// Go through each line marking them on the table
	for _, line := range diagonal_lines {
		// Diagonal line
		for i := 0; i <= line.x2-line.x1; i++ {
			// x and y grow at the same pace, so just use one variable
			y := line.x1 + i
			var x int

			// We have no guarantee that y goes from y1 to y2, so we verify
			if line.y2 > line.y1 {
				x = line.y1 + i
			} else {
				x = line.y1 - i
			}

			field[x][y]++
		}
	}

	// Count how many occurrences of crossing lines there are on the table
	count = 0

	for i := 0; i < field_size; i++ {
		for j := 0; j < field_size; j++ {
			if field[i][j] >= 2 {
				count++
			}
		}
	}

	fmt.Println("Points with diagonal lines:", count)
}
