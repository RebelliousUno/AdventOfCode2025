package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	zLog "github.com/rs/zerolog/log"
)

type RangeList struct {
	ranges []Range
}

type Range struct {
	min int
	max int
}

func (r *RangeList) isFresh(id int) bool {
	for _, rng := range r.ranges {
		if rng.isInRange(id) {
			return true
		}
	}
	return false
}

func (r *Range) isInRange(id int) bool {
	return id >= r.min && id <= r.max
}

func parseToRange(value string) (Range, error) {
	s := strings.Split(value, "-")
	if len(s) != 2 {
		return Range{}, fmt.Errorf("value %s, not in correct format for range", value)
	}
	start, err := strconv.Atoi(s[0])
	end, errEnd := strconv.Atoi(s[1])
	if err != nil {
		return Range{}, err
	}
	if errEnd != nil {
		return Range{}, errEnd
	}
	if start > end {
		return Range{}, fmt.Errorf("start %d must be less than end %d", start, end)
	}
	return Range{start, end}, nil
}

func parseToRangeList(inputRanges []string) (RangeList, error) {
	rangeList := make([]Range, 0)
	for _, v := range inputRanges {
		r, err := parseToRange(v)
		if err != nil {
			return RangeList{}, err
		}
		rangeList = append(rangeList, r)
	}
	return RangeList{rangeList}, nil
}

func parseInputToRangesAndIds(input []string) ([]string, []string) {
	ranges := make([]string, 0)
	ingredients := make([]string, 0)
	parsingRanges := true
	for _, v := range input {
		if len(strings.TrimSpace(v)) == 0 {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			ranges = append(ranges, v)
		} else {
			ingredients = append(ingredients, v)
		}
	}
	return ranges, ingredients
}

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

func convertIngredientIds(ingredients []string) ([]int, error) {
	ids := make([]int, 0)
	for _, v := range ingredients {
		id, err := strconv.Atoi(v)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func countFreshIngredients(rl RangeList, ids []int) int {
	freshCount := 0
	for _, v := range ids {
		if rl.isFresh(v) {
			freshCount++
		}
	}
	return freshCount
}

func main() {
	database := readInput()
	ranges, ids := parseInputToRangesAndIds(database)
	rangeList, errRangeList := parseToRangeList(ranges)
	if errRangeList != nil {
		zLog.Fatal().Err(errRangeList).Msg("Problem with RangeList")
	}
	ingredientsList, errIngredients := convertIngredientIds(ids)
	if errIngredients != nil {
		zLog.Fatal().Err(errIngredients).Msg("Problem with Ingredients List")
	}
	fresh := countFreshIngredients(rangeList, ingredientsList)
	zLog.Info().Msgf("Count of fresh Ingredients %d", fresh)

}
