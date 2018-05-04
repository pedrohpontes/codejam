package main

import (
	"bufio"
	"fmt"
	"os"
)

// parseNumber parses the line into a number (an array of its digits)
func parseNumber(line string) number {
	n := make([]uint8, len(line))

	for i, c := range line {
		n[i] = uint8(c - '0')
	}

	return n
}

func main() {
	if len(os.Args) < 3 {
		panic("not enough arguments")
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	count := -1
	for scanner.Scan() {
		count++
		if count == 0 {
			continue
		}

		n := parseNumber(scanner.Text())
		tidy := lastTidy(n)

		fmt.Fprintf(outputFile, "Case #%v: %v\n", count, tidy)
	}
}
