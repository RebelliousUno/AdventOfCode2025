package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	zLog "github.com/rs/zerolog/log"
)

//Dial 0-99
//Current Value
//function to turn left
//function to turn right

type Dial struct {
	CurrentValue    int
	PasswordCounter int
}

func (d *Dial) TurnLeft(steps int) {
	for i := 0; i < steps; i++ {
		d.CurrentValue--
		if d.CurrentValue < 0 {
			d.CurrentValue = 99
		}
		if d.CurrentValue == 0 {
			d.PasswordCounter++
		}
	}
}

func (d *Dial) TurnRight(steps int) {
	for i := 0; i < steps; i++ {
		d.CurrentValue++
		if d.CurrentValue > 99 {
			d.CurrentValue = 0
		}
		if d.CurrentValue == 0 {
			d.PasswordCounter++
		}
	}
}

func readInput() []string {
	// Placeholder for reading input from a file
	path := filepath.Join("password_init.txt")
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

func (d *Dial) ParseInputCommands(input []string) error {
	for _, cmd := range input {
		// Process each command
		// split into command and value
		regex, err := regexp.Compile(`([R|L])(\d+)`)
		if err != nil {
			zLog.Error().Err(err).Msg("Failed to compile regex")
			return err
		}
		match := regex.FindStringSubmatch(cmd)
		steps, err := strconv.Atoi(match[2])
		if err != nil {
			zLog.Error().Err(err).Msg("Failed to convert steps to int")
			return err
		}
		switch match[1] {
		case "R":
			d.TurnRight(steps)
		case "L":
			d.TurnLeft(steps)
		}
	}
	return nil
}

func main() {
	input := readInput()
	fmt.Println(len(input))
	dial := &Dial{CurrentValue: 50, PasswordCounter: 0}
	err := dial.ParseInputCommands(input)
	if err != nil {
		zLog.Fatal().Err(err).Msg("Failed to parse input commands")
	}
	fmt.Printf("Final Dial Value: %d\n", dial.CurrentValue)
	fmt.Printf("Password Counter: %d\n", dial.PasswordCounter)
}
