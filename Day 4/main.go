package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	zLog "github.com/rs/zerolog/log"
)

func parseInput(input string) [][]string {
	output := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		//For each line in lines
		splitLine := make([]string, 0)
		for _, s := range line {
			splitLine = append(splitLine, string(s))
		}
		output = append(output, splitLine)
	}
	return output
}

func addAdjacentCount(space string) int {
	if space == "@" {
		return 1
	} else {
		return 0
	}
}

func accessibleRolls(warehouseMap [][]string) int {
	sum := 0
	xDimension := len(warehouseMap[0])
	yDimension := len(warehouseMap)
	for x := 0; x < xDimension; x++ {
		for y := 0; y < yDimension; y++ {
			this := warehouseMap[x][y]
			if this != "@" {
				// if the value we're looking at is not @ we just skip
				continue
			}
			adjacentCount := 0
			if x > 0 {
				adjacentCount += addAdjacentCount(warehouseMap[x-1][y])
				if y > 0 {
					adjacentCount += addAdjacentCount(warehouseMap[x-1][y-1])
				}
				if y < yDimension-1 {
					adjacentCount += addAdjacentCount(warehouseMap[x-1][y+1])
				}
			}
			if x < xDimension-1 {
				adjacentCount += addAdjacentCount(warehouseMap[x+1][y])
				if y > 0 {
					adjacentCount += addAdjacentCount(warehouseMap[x+1][y-1])
				}
				if y < yDimension-1 {
					adjacentCount += addAdjacentCount(warehouseMap[x+1][y+1])
				}
			}
			if y > 0 {
				adjacentCount += addAdjacentCount(warehouseMap[x][y-1])
			}
			if y < yDimension-1 {
				adjacentCount += addAdjacentCount(warehouseMap[x][y+1])
			}
			if adjacentCount < 4 {
				sum++
			}
		}
	}
	return sum
}

func readInput() []string {
	// Placeholder for reading input from a file
	path := filepath.Join("input.txt")
	content := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		zLog.Fatal().Err(err).Msg("Failed to open input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		zLog.Fatal().Err(err).Msg("Failed to read input file")
	}
	return content
}

func main() {
	count := accessibleRolls(parseInput(strings.Join(readInput(), "\n")))
	fmt.Printf("%d, accessible rolls", count)
}
