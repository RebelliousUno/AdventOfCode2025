package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	zLog "github.com/rs/zerolog/log"
)

func readInput() []string {
	// Placeholder for reading input from a file
	path := filepath.Join("input.txt")
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

type TacyonMap struct {
	currentState     [][]string
	currentStep      int
	emitterPositions [][]int
	numSplits        int
}

func (t *TacyonMap) parseDiagramToMap(input []string) {
	output := make([][]string, len(input))
	for i := range input {
		output[i] = strings.Split(input[i], "")
	}
	t.currentState = output
	t.currentStep = 0
	t.emitterPositions = make([][]int, len(input))
	t.numSplits = 0
}

func (t *TacyonMap) doStep() bool {
	if t.currentStep == 0 {
		//Do Emit
		//Find the position of the S
		//grab row
		row := t.currentState[t.currentStep]
		for i := range row {
			if row[i] == "S" {
				t.emitterPositions[t.currentStep] = []int{i}
				//we've found an emmision point
				t.currentStep++
				return true
			}
			//we've not found an emission point
		}
		return false
	}
	if t.currentStep < len(t.currentState) {
		//Do Emit
		// look at the emission position of the previous row
		// replace . with |
		// if ^ at position replace Left and right with |

		for _, pos := range t.emitterPositions[t.currentStep-1] {
			switch t.currentState[t.currentStep][pos] {
			case ".":
				t.currentState[t.currentStep][pos] = "|"
			case "^":
				t.numSplits++
				if pos > 0 {
					//replace left
					t.currentState[t.currentStep][pos-1] = "|"
				} //replace left and right with |
				if pos < len(t.currentState[t.currentStep])-1 {
					t.currentState[t.currentStep][pos+1] = "|"
				}
			}
		}
		// record |s as new emission positions
		newEmissions := make([]int, 0)
		for i, v := range t.currentState[t.currentStep] {
			if v == "|" {
				newEmissions = append(newEmissions, i)
			}
		}
		t.emitterPositions[t.currentStep] = newEmissions
		t.currentStep++
		return true
	}

	return false
}

func main() {
	readMap := readInput()
	t := TacyonMap{}
	t.parseDiagramToMap(readMap)
	for {
		ok := t.doStep()
		if !ok {
			break
		}
	}
	zLog.Info().Int("FinalStep", t.currentStep).Int("Splits", t.numSplits).Msg("")
}
