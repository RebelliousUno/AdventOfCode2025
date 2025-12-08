package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	expected := [][]string{
		{".", ".", "@", "@", ".", "@", "@", "@", "@", "."},
		{"@", "@", "@", ".", "@", ".", "@", ".", "@", "@"},
		{"@", "@", "@", "@", "@", ".", "@", ".", "@", "@"},
		{"@", ".", "@", "@", "@", "@", ".", ".", "@", "."},
		{"@", "@", ".", "@", "@", "@", "@", ".", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", ".", "@"},
		{".", "@", ".", "@", ".", "@", ".", "@", "@", "@"},
		{"@", ".", "@", "@", "@", ".", "@", "@", "@", "@"},
		{".", "@", "@", "@", "@", "@", "@", "@", "@", "."},
		{"@", ".", "@", ".", "@", "@", "@", ".", "@", "."},
	}
	output := parseInput(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v to equal %v", output, expected)
	}
}

func Test_countAccessibleRolls(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	expected := 13
	count := accessibleRolls(parseInput(input))
	if expected != count {
		t.Errorf("Expected %d, to equal %d", count, expected)
	}

}
