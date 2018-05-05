package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parsePancakeProblem parses the input line (as described in the problem) and
// returns the list of pancakes and the size of the flipper (K)
func parsePancakeProblem(line string) ([]pancake, int) {
	splitSpace := strings.Split(line, " ")
	if len(splitSpace) < 2 {
		panic("wrong input to parse; should have received two strings per line")
	}
	pancakes := splitSpace[0]
	flipperStr := splitSpace[1]

	ps := make([]pancake, len(pancakes))
	for i, p := range pancakes {
		ps[i] = pancake(p)
	}

	flipperSize, err := strconv.Atoi(flipperStr)
	if err != nil {
		panic(err)
	}

	return ps, flipperSize
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

		pancakes, flipperSize := parsePancakeProblem(scanner.Text())
		numberOfFlips := minPancakeFlips(pancakes, flipperSize)

		if numberOfFlips >= 0 {
			fmt.Fprintf(outputFile, "Case #%v: %v\n", count, numberOfFlips)
		} else {
			fmt.Fprintf(outputFile, "Case #%v: IMPOSSIBLE\n", count)
		}
	}
}
