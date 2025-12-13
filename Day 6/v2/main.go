package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	zLog "github.com/rs/zerolog/log"
)

func parseInput(input []string) ([][]int, []string, error) {
	//Convert input to list of ints
	numbersRows := make([][]int, 0)
	for _, v := range input[:len(input)-1] {
		reg, err := regexp.Compile(`(\d+)`)
		if err != nil {
			zLog.Error().Err(err).Msg("Bad Regex")
			return nil, nil, err
		}
		row := make([]int, 0)
		matches := reg.FindAllString(v, -1)
		for _, vi := range matches {
			intVal, errConv := strconv.Atoi(vi)
			if errConv != nil {
				zLog.Error().Err(errConv).Msg("Bad convert")
				return nil, nil, errConv
			}
			row = append(row, intVal)
		}
		numbersRows = append(numbersRows, row)
	}
	opsRegex, errOps := regexp.Compile(`([\+\*])`)
	if errOps != nil {
		zLog.Error().Err(errOps).Msg("Bad ops regext")
		return nil, nil, errOps
	}
	ops := opsRegex.FindAllString(input[len(input)-1], -1)

	return numbersRows, ops, nil
}

func doTheMaths(numbers [][]int, ops []string) []int {
	values := make([]int, len(ops))

	for i, v := range ops {
		switch v {
		case "*":
			values[i] = 1
			for inumbers := range numbers {
				values[i] *= numbers[inumbers][i]
			}
		case "+":
			for inumbers := range numbers {
				values[i] += numbers[inumbers][i]
			}
		}
	}
	return values
}

func sumTheMaths(maths []int) int {
	sum := 0
	for _, v := range maths {
		sum += v
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
	sums := readInput()
	numbers, operations, err := parseInput(sums)
	if err != nil {
		zLog.Fatal().Err(err).Msg("It's f***ed")	
	}
	doneMaths := doTheMaths(numbers, operations)
	final := sumTheMaths(doneMaths)
	zLog.Info().Int("Result", final).Msg("Final Result")
}
