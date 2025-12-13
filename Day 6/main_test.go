package main

import (
	"reflect"
	"testing"
)

func Test_calcSumsDay2(t *testing.T) {
	input := [][]string{
		{"1", "2", "3", " ", "3", "2", "8", " ", " ", "5", "1", " ", "6", "4", " "},
		{" ", "4", "5", " ", "6", "4", " ", " ", "3", "8", "7", " ", "2", "3", " "},
		{" ", " ", "6", " ", "9", "8", " ", " ", "2", "1", "5", " ", "3", "1", "4"},
		{"*", " ", " ", " ", "+", " ", " ", " ", "*", " ", " ", " ", "+", " ", " "},
	}
	expectedNumbers := [][]int{
		{1, 24, 356},
		{8, 248, 369},
		{32, 175, 581},
		{4, 431, 623},
	}
	expectedOps := []string{"*", "+", "*", "+"}
	numbers, ops, err := makeDay2Sums(input)
	doneMaths := doTheMaths(numbers, ops)
	sum := sumTheMaths(doneMaths)
	expectedSum := 3263827
	if err != nil {
		t.Error("Expected no error to be thrown but was", err)
	}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected %v to equal %v", numbers, expectedNumbers)
	}
	if !reflect.DeepEqual(ops, expectedOps) {
		t.Errorf("Expected %v to equal %v", ops, expectedOps)
	}
	if sum != expectedSum {
		t.Errorf("Expected %d to be %d", sum, expectedSum)
	}
}

func Test_makeSumsDay2(t *testing.T) {
	input := [][]string{
		{"1", "2", "3", " ", "3", "2", "8", " ", " ", "5", "1", " ", "6", "4", " "},
		{" ", "4", "5", " ", "6", "4", " ", " ", "3", "8", "7", " ", "2", "3", " "},
		{" ", " ", "6", " ", "9", "8", " ", " ", "2", "1", "5", " ", "3", "1", "4"},
		{"*", " ", " ", " ", "+", " ", " ", " ", "*", " ", " ", " ", "+", " ", " "},
	}
	expectedNumbers := [][]int{
		{1, 24, 356},
		{8, 248, 369},
		{32, 175, 581},
		{4, 431, 623},
	}
	expectedOps := []string{"*", "+", "*", "+"}
	numbers, ops, err := makeDay2Sums(input)
	if err != nil {
		t.Error("Expected no error to be thrown but was", err)
	}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected %v to equal %v", numbers, expectedNumbers)
	}
	if !reflect.DeepEqual(ops, expectedOps) {
		t.Errorf("Expected %v to equal %v", ops, expectedOps)
	}
}

func Test_parseDay2(t *testing.T) {
	input := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  "}
	expectedOutPut := [][]string{
		{"1", "2", "3", " ", "3", "2", "8", " ", " ", "5", "1", " ", "6", "4", " "},
		{" ", "4", "5", " ", "6", "4", " ", " ", "3", "8", "7", " ", "2", "3", " "},
		{" ", " ", "6", " ", "9", "8", " ", " ", "2", "1", "5", " ", "3", "1", "4"},
		{"*", " ", " ", " ", "+", " ", " ", " ", "*", " ", " ", " ", "+", " ", " "},
	}
	day2Output := parseDay2(input)
	if !reflect.DeepEqual(day2Output, expectedOutPut) {
		t.Errorf("%v not %v", day2Output, expectedOutPut)
	}
}

func Test_sumTheMaths(t *testing.T) {
	input := []int{33210, 490, 4243455, 401}
	expectedOutput := 4277556
	sum := sumTheMaths(input)
	if sum != expectedOutput {
		t.Errorf("Expected %d to be %d", sum, expectedOutput)
	}
}

func Test_doTheMaths(t *testing.T) {
	numbersInput := [][]int{
		{123, 328, 51, 64}, {45, 64, 387, 23}, {6, 98, 215, 314},
	}
	opsInput := []string{"*", "+", "*", "+"}
	expectedOutput := []int{33210, 490, 4243455, 401}
	mathsDone := doTheMaths(numbersInput, opsInput)
	if !reflect.DeepEqual(mathsDone, expectedOutput) {
		t.Errorf("Expected %v to equal %v", mathsDone, expectedOutput)
	}
}

func Test_parseInput(t *testing.T) {
	input := []string{"123 328  51 64",
		"45 64  387 23",
		"6 98  215 314",
		"*   +   *   +  "}
	expectedNumbers := [][]int{
		{123, 328, 51, 64}, {45, 64, 387, 23}, {6, 98, 215, 314},
	}
	expectedOps := []string{"*", "+", "*", "+"}
	numbers, ops, err := parseInput(input)
	if err != nil {
		t.Error("Expected no error to be thrown but was", err)
	}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected %v to equal %v", numbers, expectedNumbers)
	}
	if !reflect.DeepEqual(ops, expectedOps) {
		t.Errorf("Expected %v to equal %v", ops, expectedOps)
	}
}
