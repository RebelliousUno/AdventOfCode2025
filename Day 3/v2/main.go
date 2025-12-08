package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	zLog "github.com/rs/zerolog/log"
)

func findLargestNumberInString(s string) (int, int) {
	largest := 0
	indxed := 0
	for index, value := range s {
		currentNumber := int(value - '0')
		if currentNumber > largest {
			largest = currentNumber
			indxed = index
		}
		if currentNumber == 9 {
			return 9, index
		}
	}
	return largest, indxed
}

func findJoltageForBankWithLength(bank string, joltLength int) string {
	if joltLength == 1 {
		largest, _ := findLargestNumberInString(bank)
		return strconv.Itoa(largest)
	}
	largest, indexOfLargest := findLargestNumberInString(bank[:len(bank)-joltLength+1])
	result := strconv.Itoa(largest) + findJoltageForBankWithLength(bank[indexOfLargest+1:], joltLength-1)
	return result
}

func findJoltageForBank(bank string) int {
	// find the largest number in the string from 0 to len(bank)-2
	// once found, then find the largest number from that index to the end of the string
	largest, indexOfLargest := findLargestNumberInString(bank[:len(bank)-1])
	largest2, _ := findLargestNumberInString(bank[indexOfLargest+1:])
	return largest*10 + largest2
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

func sumJoltageFromBanks(banks []string) int {
	total := 0
	for _, bank := range banks {
		total += findJoltageForBank(bank)
	}
	return total
}

func sumJoltageFromBanksWithLength(banks []string, length int) int {
	total := 0
	for _, bank := range banks {
		val, err := strconv.Atoi(findJoltageForBankWithLength(bank, 12))
		if err != nil {
			zLog.Error().Err(err).Msgf("Failed to convert joltage for bank %s", bank)
			continue
		}	
		total += val
	}
	return total
}

func main() {
	// Entry point for Day 3 solution
	input := readInput()
	totalJoltage := sumJoltageFromBanks(input)
	zLog.Info().Msgf("Total Joltage: %d", totalJoltage)

	totalJoltageWithLength := sumJoltageFromBanksWithLength(input, 12)
	zLog.Info().Msgf("Total Joltage with Length 12: %d", totalJoltageWithLength)
}
