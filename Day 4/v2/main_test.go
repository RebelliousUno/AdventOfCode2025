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
	count, _ := accessibleRolls(parseInput(input))
	if expected != count {
		t.Errorf("Expected %d, to equal %d", count, expected)
	}
}

func TestMarkedForRemoval(t *testing.T) {
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
	expectedNewString := `..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.`
	expectedNewMap := parseInput(expectedNewString)
	count, newMap := accessibleRolls(parseInput(input))
	if expected != count {
		t.Errorf("Expected %d, to equal %d", count, expected)
	}
	if !reflect.DeepEqual(newMap, expectedNewMap) {
		t.Errorf("Expected %v to equal %v", newMap, expectedNewMap)
	}
}
