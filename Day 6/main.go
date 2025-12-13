package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"

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
	//This doesn't quite work because the row might not be square
	for i, v := range ops {
		switch v {
		case "*":
			values[i] = 1
			for inumbers := range numbers[i] {
				values[i] *= numbers[i][inumbers]
			}
		case "+":
			for inumbers := range numbers[i] {
				values[i] += numbers[i][inumbers]
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

func makeDay2Sums(input [][]string) ([][]int, []string, error) {
	numbers := make([][]int, 1)
	lastVal := ""
	for iCol := range input[0] {
		val := ""
		for iRow := range input[:len(input)-1] {
			val += input[iRow][iCol]
		}
		trimmedVal := strings.TrimSpace(val)
		if trimmedVal == "" {
			if trimmedVal != lastVal {
				numbers = append(numbers, make([]int, 0))
			}
			lastVal = trimmedVal
			continue
		}
		lastVal = trimmedVal
		number, err := strconv.Atoi(trimmedVal)
		if err != nil {
			return nil, nil, err
		}
		numbers[len(numbers)-1] = append(numbers[len(numbers)-1], number)
	}
	for _, v := range numbers {
		slices.Sort(v)
	}
	opsRegex, errOps := regexp.Compile(`([\+\*])`)
	if errOps != nil {
		zLog.Error().Err(errOps).Msg("Bad ops regext")
		return nil, nil, errOps
	}
	ops := opsRegex.FindAllString(strings.Join(input[len(input)-1], ""), -1)
	return numbers, ops, nil
}

func parseDay2(input []string) [][]string {
	output := make([][]string, 0)
	for _, v := range input {
		output = append(output, strings.Split(v, ""))
	}
	return output
}

func main() {
	sums := readInput()
	//numbers, operations, err := parseInput(sums)
	// if err != nil {
	// 	zLog.Fatal().Err(err).Msg("It's f***ed")
	// }
	// doneMaths := doTheMaths(numbers, operations)
	// final := sumTheMaths(doneMaths)
	// zLog.Info().Int("Result", final).Msg("Final Result")

	// ///day 2 part 2

	numbersDay2, opsDay2, errDay2 := makeDay2Sums(parseDay2(sums))
	if errDay2 != nil {
		zLog.Fatal().Err(errDay2).Msg("It's f***ed")
	}
	doneMaths := doTheMaths(numbersDay2, opsDay2)
	final := sumTheMaths(doneMaths)
	zLog.Info().Int("Result", final).Msg("Day 2 Final Result")

}
