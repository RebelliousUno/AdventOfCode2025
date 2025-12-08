package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitInput(t *testing.T) {
	input := "10-20,30-40"
	expected := []string{"10-20", "30-40"}
	result := splitInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetRanges(t *testing.T) {
	input := "10-20"
	expectedStart := uint64(10)
	expectedEnd := uint64(20)
	start, end, err := getRanges(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if start != expectedStart || end != expectedEnd {
		t.Errorf("Expected (%d, %d), got (%d, %d)", expectedStart, expectedEnd, start, end)
	}
}

func TestInvalid_WithSplit(t *testing.T) {
	if !isValid(123456) {
		t.Errorf("Expected 123456 to be valid")
	}
	if isValid(123123) {
		t.Errorf("Expected 123123 to be invalid")
	}
}

func TestValid_Regex(t *testing.T) {
	valid, err := isValidRegex(110)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !valid {
		t.Errorf("Expected 110 to be valid")
	}
}

func TestInvalid_Regex(t *testing.T) {
	valid, err := isValidRegex(123456)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !valid {
		t.Errorf("Expected 123456 to be valid")
	}
	valid, err = isValidRegex(123123)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if valid {
		t.Errorf("Expected 123123 to be invalid")
	}
}

func Test_valid_odd_length(t *testing.T) {
	valid, err := isValidRegex(12345)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !valid {
		t.Errorf("Expected 12345 to be valid")
	}
	valid = isValid(12345)
	if !valid {
		t.Errorf("Expected 12345 to be valid")
	}
	valid, err = isValidRegex(12312)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !valid {
		t.Errorf("Expected 12312 to be valid")
	}
	valid = isValid(12312)
	if !valid {
		t.Errorf("Expected 12312 to be valid")
	}
}

func Test_DuplicatedValues_Day2(t *testing.T) {
	/*
			11-22 has two invalid IDs, 11 and 22.
		95-115 has one invalid ID, 99.
		998-1012 has one invalid ID, 1010.
		1188511880-1188511890 has one invalid ID, 1188511885.
		222220-222224 has one invalid ID, 222222.
		1698522-1698528 contains no invalid IDs.
		446443-446449 has one invalid ID, 446446.
		38593856-38593862 has one invalid ID, 38593859.
	*/
	var tests = []struct {
		a, b uint64
		want []uint64
	}{
		{11, 22, []uint64{11, 22}},
		{95, 115, []uint64{99, 111}},
		{998, 1012, []uint64{999, 1010}},
		{1188511880, 1188511890, []uint64{1188511885}},
		{222220, 222224, []uint64{222222}},
		{1698522, 1698528, []uint64{}},
		{446443, 446449, []uint64{446446}},
		{38593856, 38593862, []uint64{38593859}},
		{565653, 565659, []uint64{565656}},
		{824824821, 824824827, []uint64{824824824}},
		{2121212118, 2121212124, []uint64{2121212121}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d-%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			// the non-regex function works with ints, so convert inputs and outputs
			duplicated := findDupilicatedNumbersInRangeDay2(tt.a, tt.b)
			// convert duplicated ([]int) to []uint64 for comparison
			dupUint := make([]uint64, 0, len(duplicated))
			for _, v := range duplicated {
				dupUint = append(dupUint, uint64(v))
			}
			if !reflect.DeepEqual(dupUint, tt.want) {
				t.Errorf("findDupilicatedNumbersInRangeDay2(%d, %d): expected %v, got %v", tt.a, tt.b, tt.want, dupUint)
			}
		})
	}
}

func Test_DuplicatedValuesRegex(t *testing.T) {
	/*
			11-22 has two invalid IDs, 11 and 22.
		95-115 has one invalid ID, 99.
		998-1012 has one invalid ID, 1010.
		1188511880-1188511890 has one invalid ID, 1188511885.
		222220-222224 has one invalid ID, 222222.
		1698522-1698528 contains no invalid IDs.
		446443-446449 has one invalid ID, 446446.
		38593856-38593862 has one invalid ID, 38593859.
	*/
	var tests = []struct {
		a, b uint64
		want []uint64
	}{
		{11, 22, []uint64{11, 22}},
		{95, 115, []uint64{99}},
		{998, 1012, []uint64{1010}},
		{1188511880, 1188511890, []uint64{1188511885}},
		{222220, 222224, []uint64{222222}},
		{1698522, 1698528, []uint64{}},
		{446443, 446449, []uint64{446446}},
		{38593856, 38593862, []uint64{38593859}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d-%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			duplicated, err := findDupilicatedNumbersInRangeRegex(tt.a, tt.b)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(duplicated, tt.want) {
				t.Errorf("findDupilicatedNumbersInRangeRegex(%d, %d): expected %v, got %v", tt.a, tt.b, tt.want, duplicated)
			}
		})
	}
}

func Test_SumDuplicatedValues(t *testing.T) {
	numbers := []uint64{11, 22, 99, 1010}
	var expectedSum uint64 = 1142
	sum := sumDuplicatedNumbers(numbers)
	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}
}

func TestSplitString_IntoParts(t *testing.T) {
	var tests = []struct {
		stringInput string
		partCount   int
		want        []string
	}{
		{"123456", 1, []string{"123456"}},
		{"123456", 2, []string{"123", "456"}},
		{"123456", 3, []string{"12", "34", "56"}},
		{"123456", 6, []string{"1", "2", "3", "4", "5", "6"}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s-%d", tt.stringInput, tt.partCount)
		t.Run(testname, func(t *testing.T) {
			parts, err := getParts(tt.stringInput, tt.partCount)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(parts, tt.want) {
				t.Errorf("getParts(%s, %d): expected %v, got %v", tt.stringInput, tt.partCount, tt.want, parts)
			}
		})
	}
}

func TestSplitString_IntoParts_Errors(t *testing.T) {
	inputString := "12345"
	partCountTooLarge := 10
	partCountZero := 0
	partCountNotDivisible := 3
	_, err := getParts(inputString, partCountNotDivisible)
	if err == nil {
		t.Errorf("Expected error for partCount %d not dividing length of string %s", partCountNotDivisible, inputString)
	}
	_, err = getParts(inputString, partCountTooLarge)
	if err == nil {
		t.Errorf("Expected error for partCount %d greater than length of string %s", partCountTooLarge, inputString)
	}
	_, err = getParts(inputString, partCountZero)
	if err == nil {
		t.Errorf("Expected error for partCount %d less than or equal to 0", partCountZero)
	}
}

func Test_ArePartsEqual(t *testing.T) {
	var tests = []struct {
		parts []string
		want  bool
	}{
		{[]string{"12", "12", "12"}, true},
		{[]string{"12", "34", "12"}, false},
		{[]string{"1", "1", "1", "1"}, true},
		{[]string{"1", "2", "1", "1"}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.parts)
		t.Run(testname, func(t *testing.T) {
			result := arePartsEqual(tt.parts)
			if result != tt.want {
				t.Errorf("arePartsEqual(%v): expected %v, got %v", tt.parts, tt.want, result)
			}
		})
	}
}
