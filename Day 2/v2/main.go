package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	zLog "github.com/rs/zerolog/log"
)

func splitInput(input string) []string {
	var result []string
	result = strings.Split(input, ",")
	return result
}

func getRanges(input string) (uint64, uint64, error) {
	parts := strings.Split(input, "-")

	start, err := strconv.ParseUint(parts[0], 10, 0)
	if err != nil {
		return 0, 0, err
	}
	end, err := strconv.ParseUint(parts[1], 10, 0)
	if err != nil {
		return 0, 0, err
	}
	return start, end, nil
}

func isValid(number int) bool {
	numberStr := strconv.Itoa(number)
	//split the number string in half
	mid := len(numberStr) / 2
	firstHalf := numberStr[:mid]
	secondHalf := numberStr[mid:]
	return firstHalf != secondHalf
}

func arePartsEqual(parts []string) bool {
	maxIndex := len(parts) - 1
	for index := range parts {
		if index >= maxIndex {
			break
		}
		if parts[index] != parts[index+1] {
			return false
		}
	}
	return true
}

func getParts(numberStr string, partCount int) ([]string, error) {
	if partCount <= 0 {
		return []string{}, fmt.Errorf("partCount must be greater than 0")
	}
	if partCount > len(numberStr) {
		return []string{}, fmt.Errorf("partCount must be less than or equal to the length of numberStr")
	}
	if len(numberStr)%partCount != 0 {
		return []string{}, fmt.Errorf("numberStr length must be divisible by partCount")
	}
	result := make([]string, 0)
	partLength := len(numberStr) / partCount
	for i := 0; i < partCount; i++ {
		split := numberStr[partLength*i : partLength*(i+1)]
		result = append(result, split)
	}
	return result, nil
}

func isValidDay2(number uint64) bool {
	numberStr := strconv.FormatUint(number, 10)
	for parts := 2; parts <= len(numberStr); parts++ {
		if len(numberStr)%parts != 0 {
			//skip if not divisible
			continue
		}
		partList, err := getParts(numberStr, parts)
		if err != nil {
			return true
		}
		if arePartsEqual(partList) {
			return false
		}
	}
	return true
}

func isValidRegex(number uint64) (bool, error) {
	numberStr := strconv.FormatUint(number, 10)
	mid := len(numberStr) / 2
	rest := len(numberStr) - mid
	pattern := fmt.Sprintf(`(\d{%d})(\d{%d})`, mid, rest)
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(numberStr)
	return match[1] != match[2], nil
}

func findDupilicatedNumbersInRangeDay2(rangeStart, rangeEnd uint64) []uint64 {
	result := make([]uint64, 0)
	for i := rangeStart; i <= rangeEnd; i++ {
		if !isValidDay2(i) {
			result = append(result, i)
		}
	}
	return result
}

func findDupilicatedNumbersInRange(rangeStart, rangeEnd int) []int {
	result := make([]int, 0)
	for i := rangeStart; i <= rangeEnd; i++ {
		if !isValid(i) {
			result = append(result, i)
		}
	}
	return result
}

func findDupilicatedNumbersInRangeRegex(rangeStart, rangeEnd uint64) ([]uint64, error) {
	result := make([]uint64, 0)
	for i := rangeStart; i <= rangeEnd; i++ {
		valid, err := isValidRegex(i)
		if err != nil {
			return result, err
		}
		if !valid {
			result = append(result, i)
		}
	}
	return result, nil
}

func sumDuplicatedNumbers(numbers []uint64) uint64 {
	var sum uint64 = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	var sum uint64
	sum = 0
	input := "655-1102,2949-4331,885300-1098691,1867-2844,20-43,4382100-4484893,781681037-781860439,647601-734894,2-16,180-238,195135887-195258082,47-64,4392-6414,6470-10044,345-600,5353503564-5353567532,124142-198665,1151882036-1151931750,6666551471-6666743820,207368-302426,5457772-5654349,72969293-73018196,71-109,46428150-46507525,15955-26536,65620-107801,1255-1813,427058-455196,333968-391876,482446-514820,45504-61820,36235767-36468253,23249929-23312800,5210718-5346163,648632326-648673051,116-173,752508-837824"
	ranges := splitInput(input)
	for _, r := range ranges {
		start, end, err := getRanges(r)
		if err != nil {
			zLog.Fatal().Msgf("Error parsing range %s: %v", r, err)
		}
		duplicated := findDupilicatedNumbersInRangeDay2(start, end)
		fmt.Printf("Duplicated Numbers %v", duplicated)
		sum += sumDuplicatedNumbers(duplicated)
		fmt.Printf("Range %d-%d: Sum of duplicated numbers: %d\n", start, end, sum)
	}
	fmt.Printf("Total Sum of duplicated numbers: %d\n", sum)
}
