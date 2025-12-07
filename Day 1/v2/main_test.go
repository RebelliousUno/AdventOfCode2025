package main

import "testing"

func TestDial_TurnOnZero(t *testing.T) {

	dial := &Dial{CurrentValue: 0, PasswordCounter: 0}
	dial.TurnRight(100)
	if dial.CurrentValue != 0 {
		t.Errorf("Expected CurrentValue to be 0, got %d", dial.CurrentValue)
	}
	if dial.PasswordCounter != 1 {
		t.Errorf("Expected PasswordCounter to be 1, got %d", dial.PasswordCounter)
	}
	dial = &Dial{CurrentValue: 0, PasswordCounter: 0}

	dial.TurnRight(200)
	if dial.CurrentValue != 0 {
		t.Errorf("Expected CurrentValue to be 0, got %d", dial.CurrentValue)
	}
	if dial.PasswordCounter != 2 {
		t.Errorf("Expected PasswordCounter to be 2, got %d", dial.PasswordCounter)
	}
}

func TestDial_TurnLeftOnZero(t *testing.T) {

	dial := &Dial{CurrentValue: 0, PasswordCounter: 0}
	dial.TurnLeft(100)
	if dial.CurrentValue != 0 {
		t.Errorf("Expected CurrentValue to be 0, got %d", dial.CurrentValue)
	}
	if dial.PasswordCounter != 1 {
		t.Errorf("Expected PasswordCounter to be 1, got %d", dial.PasswordCounter)
	}
	dial = &Dial{CurrentValue: 0, PasswordCounter: 0}

	dial.TurnLeft(200)
	if dial.CurrentValue != 0 {
		t.Errorf("Expected CurrentValue to be 0, got %d", dial.CurrentValue)
	}
	if dial.PasswordCounter != 2 {
		t.Errorf("Expected PasswordCounter to be 2, got %d", dial.PasswordCounter)
	}
}

func TestDial_PasswordCounter(t *testing.T) {
	dial := &Dial{CurrentValue: 95, PasswordCounter: 0}
	dial.TurnRight(1000)
	if dial.PasswordCounter != 10 {
		t.Errorf("Expected PasswordCounter to be 10, got %d", dial.PasswordCounter)
	}
}
func TestDial_TurnLeft5(t *testing.T) {
	dial := &Dial{CurrentValue: 5, PasswordCounter: 0}
	dial.TurnLeft(10)
	if dial.CurrentValue != 95 {
		t.Errorf("Expected CurrentValue to be 95, got %d", dial.CurrentValue)
	}
	dial.TurnRight(5)
	if dial.CurrentValue != 0 {
		t.Errorf("Expected CurrentValue to be 0, got %d", dial.CurrentValue)
	}
	if dial.PasswordCounter != 2 {
		t.Errorf("Expected PasswordCounter to be 2, got %d", dial.PasswordCounter)
	}
}

func TestDial_TurnLeft_Normal(t *testing.T) {
	dial := &Dial{CurrentValue: 50}
	dial.TurnLeft(10)
	if dial.CurrentValue != 40 {
		t.Errorf("Expected CurrentValue to be 40, got %d", dial.CurrentValue)
	}
}

func TestDial_TurnRight_Normal(t *testing.T) {
	dial := &Dial{CurrentValue: 0}
	dial.TurnRight(10)
	if dial.CurrentValue != 10 {
		t.Errorf("Expected CurrentValue to be 10, got %d", dial.CurrentValue)
	}
}

func TestDial_TurnLeft_WrapAround(t *testing.T) {
	dial := &Dial{CurrentValue: 5}
	dial.TurnLeft(10)
	if dial.CurrentValue != 95 {
		t.Errorf("Expected CurrentValue to be 95, got %d", dial.CurrentValue)
	}
}

func TestDial_TurnRight_WrapAround(t *testing.T) {
	dial := &Dial{CurrentValue: 95}
	dial.TurnRight(10)
	if dial.CurrentValue != 5 {
		t.Errorf("Expected CurrentValue to be 5, got %d", dial.CurrentValue)
	}
}

func TestDial_PasswordCounter_Left(t *testing.T) {
	dial := &Dial{CurrentValue: 5, PasswordCounter: 0}
	dial.TurnLeft(5)
	if dial.PasswordCounter != 1 {
		t.Errorf("Expected PasswordCounter to be 1, got %d", dial.PasswordCounter)
	}
}

func TestDial_PasswordCounter_Right(t *testing.T) {
	dial := &Dial{CurrentValue: 95, PasswordCounter: 0}
	dial.TurnRight(5)
	if dial.PasswordCounter != 1 {
		t.Errorf("Expected PasswordCounter to be 1, got %d", dial.PasswordCounter)
	}
}

func TestDial_NoPasswordCounter_Left(t *testing.T) {
	dial := &Dial{CurrentValue: 10, PasswordCounter: 0}
	dial.TurnLeft(5)
	if dial.PasswordCounter != 0 {
		t.Errorf("Expected PasswordCounter to be 0, got %d", dial.PasswordCounter)
	}
}

func TestDial_NoPasswordCounter_Right(t *testing.T) {
	dial := &Dial{CurrentValue: 90, PasswordCounter: 0}
	dial.TurnRight(5)
	if dial.PasswordCounter != 0 {
		t.Errorf("Expected PasswordCounter to be 0, got %d", dial.PasswordCounter)
	}
}

func TestDial_WithInputCommands(t *testing.T) {
	dial := &Dial{CurrentValue: 50, PasswordCounter: 0}
	commands := []string{"R50", "L20", "R30", "L40"}
	err := dial.ParseInputCommands(commands)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedValue := 70
	if dial.CurrentValue != expectedValue {
		t.Errorf("Expected CurrentValue to be %d, got %d", expectedValue, dial.CurrentValue)
	}
	expectedCounter := 3
	if dial.PasswordCounter != expectedCounter {
		t.Errorf("Expected PasswordCounter to be %d, got %d", expectedCounter, dial.PasswordCounter)
	}
}

func TestDial_WithInputCommands_FromSite(t *testing.T) {
	dial := &Dial{CurrentValue: 50, PasswordCounter: 0}
	commands := []string{"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82"}
	err := dial.ParseInputCommands(commands)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedPasswordCounter := 6
	if dial.PasswordCounter != expectedPasswordCounter {
		t.Errorf("Expected PasswordCounter to be %d, got %d", expectedPasswordCounter, dial.PasswordCounter)
	}
}
func TestDial_FromExample(t *testing.T) {
	dial := &Dial{CurrentValue: 11, PasswordCounter: 0}
	dial.TurnRight(8)
	expectedValue := 19
	if dial.CurrentValue != expectedValue {
		t.Errorf("Expected CurrentValue to be %d, got %d", expectedValue, dial.CurrentValue)
	}
	dial.TurnLeft(19)
	expectedValue = 0
	if dial.CurrentValue != expectedValue {
		t.Errorf("Expected CurrentValue to be %d, got %d", expectedValue, dial.CurrentValue)
	}
	dial.TurnLeft(1)
	expectedValue = 99
	if dial.CurrentValue != expectedValue {
		t.Errorf("Expected CurrentValue to be %d, got %d", expectedValue, dial.CurrentValue)
	}
	dial.TurnRight(2)
	expectedValue = 1
	if dial.CurrentValue != expectedValue {
		t.Errorf("Expected CurrentValue to be %d, got %d", expectedValue, dial.CurrentValue)
	}
}
