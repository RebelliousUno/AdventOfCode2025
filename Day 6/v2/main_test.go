package main

import (
	"reflect"
	"testing"
)

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
