package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseBathroomProblem(line string) (numStalls uint64, numPeople uint64) {
	splitSpace := strings.Split(line, " ")
	if len(splitSpace) < 2 {
		panic("wrong input to parse; should have received two strings per line")
	}

	numStalls, err := strconv.ParseUint(splitSpace[0], 10, 64)
	if err != nil {
		panic("could not convert string to int")
	}

	numPeople, err = strconv.ParseUint(splitSpace[1], 10, 64)
	if err != nil {
		panic("could not convert string to int")
	}

	return numStalls, numPeople
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
			continue // skip first line containing number of cases
		}

		numStalls, numPeople := parseBathroomProblem(scanner.Text())
		min, max := lastPersonDistances(numStalls, numPeople)

		fmt.Fprintf(outputFile, "Case #%v: %v %v\n", count, max, min)
	}
}
