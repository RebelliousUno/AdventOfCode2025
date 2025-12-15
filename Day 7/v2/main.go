package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
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

func (t *TacyonMap) countTimelines3() (int, error) {
	output := make([][]int, 0)
	timelines := 0

	for i := range t.currentState {
		row := make([]int, len(t.currentState[i]))
		output = append(output, row)
		for ii, vi := range t.currentState[i] {
			switch vi {
			case "S":
				output[i][ii] = 1
			case ".":
				output[i][ii] = output[i][ii]
			case "|":				
				output[i][ii] += output[i-1][ii]
			case "^":
				//might need to increment by 1
				output[i][ii-1] += output[i-1][ii]
				output[i][ii+1] += output[i-1][ii]
			}
		}
	}
	for _, v:= range output[len(output)-1] {
		timelines += v
	}
	return timelines, nil
}

func (t *TacyonMap) countTimelinesBottomUp() (int, error) {

	strippedMap := make([][]string, 0)
	for _, v := range t.currentState {
		if strings.Contains(strings.Join(v, ""), "^") {
			strippedMap = append(strippedMap, v)
		}
	}

	vals := strippedMap
	timelines := 0
	//bottom row first

	vals[len(vals)-1] = strings.Split(strings.ReplaceAll(strings.Join(vals[len(vals)-1], ""), "|", "1"), "")

	for i := len(vals) - 2; i >= 0; i-- {
		for iRow, vRow := range vals[i] {
			if vRow == "|" {
				bottomLeft := "."
				if iRow > 0 {
					bottomLeft = vals[i+1][iRow-1]
				}
				// bottom := vals[i+1][iRow]
				bottomRight := "."
				if iRow < len(vals[i])-1 {
					bottomRight = vals[i+1][iRow+1]
				}
				bottomLeftInt, _ := strconv.Atoi(bottomLeft)
				// bottomInt, _ := strconv.Atoi(bottom)
				bottomRightInt, _ := strconv.Atoi(bottomRight)
				vals[i][iRow] = strconv.Itoa(bottomLeftInt + bottomRightInt)
			}
		}
	}
	return timelines, nil
}

func (t *TacyonMap) countTimelines2() (int, error) {
	// strip rows without ^
	// strippedMap := make([][]string, 0)
	// for _, v := range t.currentState {
	// 	if strings.Contains(strings.Join(v, ""), "^") {
	// 		strippedMap = append(strippedMap, v)
	// 	}
	// }
	// first row replace | with 1
	strippedMap := t.currentState
	for i, v := range strippedMap[0] {
		if v == "S" {
			// if v == "|" {
			strippedMap[0][i] = "1"
		}
	}
	for i := range strippedMap {
		if i == 0 {
			//skip first row
			continue
		}
		for ii, v := range strippedMap[i] {
			if v == "|" {
				topLeft := "0"
				above := "0"
				topRight := "0"

				if ii > 0 {
					topLeft = strippedMap[i-1][ii-1]
				}
				above = strippedMap[i-1][ii]
				if ii < len(strippedMap[i])-1 {
					topRight = strippedMap[i-1][ii+1]
				}
				//convert to numbers
				topLeftInt, errTL := strconv.Atoi(topLeft)
				if topLeft != "." && topLeft != "^" && errTL != nil {
					return 0, errTL
				}
				aboveInt, errTop := strconv.Atoi(above)
				if above != "." && above != "^" && errTop != nil {
					return 0, errTop
				}
				topRightInt, errTR := strconv.Atoi(topRight)
				if topRight != "." && topRight != "^" && errTR != nil {
					return 0, errTR
				}
				strippedMap[i][ii] = strconv.Itoa(topLeftInt + aboveInt + topRightInt)

			}
		}
	}
	// each row replace | with sum of numbers above
	// final row sum numbers
	timelines := 0
	for _, v := range strippedMap[len(strippedMap)-1] {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		timelines += val
	}
	return timelines, nil
}

func (t *TacyonMap) countTimelines() int {
	timelines := 0
	for _, v := range t.currentState {
		line := strings.Join(v, "")
		if strings.Contains(line, "^") {
			timelines += strings.Count(line, "|")
		}
	}
	return timelines
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
	timelines, _ := t.countTimelines3()
	zLog.Info().Int("FinalStep", t.currentStep).Int("Splits", t.numSplits).Int("Timelines", timelines).Msg("")
}
