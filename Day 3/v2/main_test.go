package main

import (
	"fmt"
	"testing"
)

func Test_findLargestNumberInString(t *testing.T) {
	var tests = []struct {
		input     string
		wantNum   int
		wantIndex int
	}{
		{"123456789", 9, 8},
		{"987654321", 9, 0},
		{"555555555", 5, 0},
		{"012345678", 8, 8},
		{"000000009", 9, 8},
		{"123459876", 9, 5},
		{"176843210", 8, 3},
	}
	for _, tt := range tests {
		num, index := findLargestNumberInString(tt.input)
		if num != tt.wantNum || index != tt.wantIndex {
			t.Error("findLargestNumberInString(" + tt.input + "): expected (" + fmt.Sprint(tt.wantNum) + ", " + fmt.Sprint(tt.wantIndex) + "), got (" + fmt.Sprint(num) + ", " + fmt.Sprint(index) + ")")
		}
	}
}

func Test_findJoltageForBank(t *testing.T) {
	var tests = []struct {
		bank string
		want int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}
	for _, tt := range tests {
		result := findJoltageForBank(tt.bank)
		if result != tt.want {
			t.Error("findJoltageForBank(" + tt.bank + "): expected " + fmt.Sprint(tt.want) + ", got " + fmt.Sprint(result))
		}
	}
}

func Test_sumJoltageFromBanks(t *testing.T) {
	banks := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	expectedTotal := 98 + 89 + 78 + 92
	result := sumJoltageFromBanks(banks)
	if result != expectedTotal {
		t.Error("sumJoltageFromBanks: expected " + fmt.Sprint(expectedTotal) + ", got " + fmt.Sprint(result))
	}
}

func Test_findJoltageForBankWithLength(t *testing.T) {
	var tests = []struct {
		bank string
		want string
	}{
		{"987654321111111", "987654321111"},
		{"811111111111119", "811111111119"},
		{"234234234234278", "434234234278"},
		{"818181911112111", "888911112111"},
	}
	for _, tt := range tests {
		result := findJoltageForBankWithLength(tt.bank, 12)
		if result != tt.want {
			t.Error("findJoltageForBank(" + tt.bank + "): expected " + fmt.Sprint(tt.want) + ", got " + fmt.Sprint(result))
		}
	}
}
