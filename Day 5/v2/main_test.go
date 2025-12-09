package main

import (
	"reflect"
	"testing"
)

func TestValueInRange(t *testing.T) {
	r := Range{123, 321}
	if !r.isInRange(123) {
		t.Errorf("Expected 123 to be in range inclusive")
	}
	if !r.isInRange(321) {
		t.Errorf("Expected 321 to be in range inclusive")
	}

	if !r.isInRange(222) {
		t.Errorf("Expected 222 to be in range")
	}

	if r.isInRange(500) {
		t.Errorf("Expected 500 to be out of range")
	}
}

func TestParseToRange(t *testing.T) {
	r, err := parseToRange("123-321")
	if err != nil {
		t.Error("Expected not to throw error but threw", err)
	}
	if r.min != 123 {
		t.Error("Expected min to be 123")
	}
	if r.max != 321 {
		t.Error("Expected max to be 321")
	}
}
func TestInvalidParseToRange(t *testing.T) {
	_, err := parseToRange("321-123")
	if err == nil {
		t.Error("Expected throwing an error")
	}
	_, err = parseToRange("asd-asd")
	if err == nil {
		t.Error("Expected throwing an error")
	}
	_, err = parseToRange("123-asd")
	if err == nil {
		t.Error("Expected throwing an error")
	}
	_, err = parseToRange("asd-123")
	if err == nil {
		t.Error("Expected throwing an error")
	}
	_, err = parseToRange("123321")
	if err == nil {
		t.Error("Expected throwing an error")
	}
}

func TestToRangeList(t *testing.T) {
	input := []string{"3-5", "10-14", "16-20", "12-18"}
	expected := RangeList{[]Range{Range{3, 5}, Range{10, 14}, Range{16, 20}, Range{12, 18}}}
	r, err := parseToRangeList(input)
	if err != nil {
		t.Error("Expected not to throw an error", err)
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected %v to equal %v", r, expected)
	}
}

func TestToRangesIngredients(t *testing.T) {
	input := []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"}
	expectedRanges := []string{"3-5", "10-14", "16-20", "12-18"}
	expectedIngedients := []string{"1", "5", "8", "11", "17", "32"}
	ranges, ingredients := parseInputToRangesAndIds(input)
	if !reflect.DeepEqual(ranges, expectedRanges) {
		t.Errorf("Expected %v to equal %v", ranges, expectedRanges)
	}
	if !reflect.DeepEqual(ingredients, expectedIngedients) {
		t.Errorf("Expected %v to equal %v", ingredients, expectedIngedients)
	}
}

func TestToRangesIngredientsInt(t *testing.T) {
	input := []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"}
	expectedRanges := []string{"3-5", "10-14", "16-20", "12-18"}
	expectedIngedients := []string{"1", "5", "8", "11", "17", "32"}
	expectedIngredientsInt := []int{1, 5, 8, 11, 17, 32}
	ranges, ingredients := parseInputToRangesAndIds(input)
	ids, err := convertIngredientIds(ingredients)
	if err != nil {
		t.Error("Expected no error but thrown", err)
	}
	if !reflect.DeepEqual(ranges, expectedRanges) {
		t.Errorf("Expected %v to equal %v", ranges, expectedRanges)
	}
	if !reflect.DeepEqual(ingredients, expectedIngedients) {
		t.Errorf("Expected %v to equal %v", ingredients, expectedIngedients)
	}
	if !reflect.DeepEqual(ids, expectedIngredientsInt) {
		t.Errorf("Epxected %v to equal %v", ids, expectedIngredientsInt)
	}

}
