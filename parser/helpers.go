package parser

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"utils"
)

func createError(value interface{}, err error) error {
	return fmt.Errorf("invalid value (%v), reason: %s", value, err.Error())
}

func getTimeValues(value string, rng Range) ([]int, error) {
	timeRanges := strings.Split(value, ",")
	totalValues := make(map[int]bool)
	if len(timeRanges) == 0 {
		return nil, fmt.Errorf("value is empty")
	}
	for _, timeRange := range timeRanges {
		runTime, err := parseTimeRange(timeRange, rng)
		if err != nil {
			return nil, err
		}

		isStarValue, err := addStarValue(totalValues, rng, runTime)
		if err != nil {
			return nil, err
		}
		if isStarValue {
			continue
		}

		isRangeValue, err := addRangeValue(totalValues, rng, runTime)
		if err != nil {
			return nil, err
		}
		if isRangeValue {
			continue
		}

		_, err = addSingleValue(totalValues, rng, runTime)
		if err != nil {
			return nil, err
		}
	}

	return utils.ConvertMapToSortedArr(totalValues), nil
}

func parseTimeRange(timeRange string, rng Range) (*RunTime, error) {
	rangeValues := strings.Split(timeRange, "/")
	runTime := &RunTime{
		step:      1,
		hasStep:   false,
		timeValue: rangeValues[0],
	}

	if len(rangeValues) > 2 {
		return nil, fmt.Errorf("multiple / found in one range")
	}
	if len(rangeValues) == 2 {
		step, err := strconv.Atoi(rangeValues[1])
		if err != nil {
			return nil, err
		}
		if step > rng.End || step < rng.Start {
			return nil, fmt.Errorf("%d is not a valid increment for %v", step, rng)
		}
		runTime.step = step
		runTime.hasStep = true
	}

	return runTime, nil
}

func addStarValue(totalValues map[int]bool, rng Range, runTime *RunTime) (bool, error) {
	if runTime.timeValue != "*" {
		return false, nil
	}

	timeValues, err := utils.CreateArrFrom(rng.Start, rng.End, runTime.step)
	if err != nil {
		return false, err
	}
	utils.AddToMap(totalValues, timeValues)

	return true, nil
}

func addRangeValue(totalValues map[int]bool, oldRange Range, runTime *RunTime) (bool, error) {
	rangeValue := strings.Split(runTime.timeValue, "-")
	if len(rangeValue) == 1 {
		return false, nil
	} else if len(rangeValue) > 2 {
		return false, fmt.Errorf("too many - found")
	}

	start, err := strconv.Atoi(rangeValue[0])
	if err != nil {
		return false, err
	}
	if start < oldRange.Start {
		return false, fmt.Errorf("values is out of range %v", oldRange)
	}

	end, err := strconv.Atoi(rangeValue[1])
	if err != nil {
		return false, err
	}
	if end > oldRange.End {
		return false, fmt.Errorf("value is out of range %v", oldRange)
	}

	timeValues, err := utils.CreateArrFrom(start, end, runTime.step)
	if err != nil {
		return false, err
	}
	utils.AddToMap(totalValues, timeValues)

	return true, nil
}

func addSingleValue(totalValues map[int]bool, rng Range, runTime *RunTime) (bool, error) {
	singleValue, err := strconv.Atoi(runTime.timeValue)
	if err != nil {
		return false, err
	}
	if singleValue < rng.Start || singleValue > rng.End {
		return false, fmt.Errorf("value is out of range %v", rng)
	}

	if runTime.hasStep {
		timeValues, err := utils.CreateArrFrom(singleValue, rng.End, runTime.step)
		if err != nil {
			return false, err
		}
		utils.AddToMap(totalValues, timeValues)
	} else {
		totalValues[singleValue] = true
	}

	return true, nil
}

func validateInRange(value []int, rng Range) error {
	for _, v := range value {
		if v < rng.Start || v > rng.End {
			return fmt.Errorf("%d is out of range %v", v, rng)
		}
	}
	return nil
}

func FindBestCronIntConversion(values []int, maxRange Range ) string {
	if len(values) == 0 {
		return "*"
	}
		sort.Ints(values)
	sortedValues := slices.Compact(values)
	leastRanges := FindLeastCronRanges(sortedValues)
	sort.Slice(leastRanges, func(i, j int) bool {
        return leastRanges[i].LongestRange[0] < leastRanges[j].LongestRange[0]
	})
	cronRanges := []string{}
	for i := 0; i< len(leastRanges); i++ {
		cronRanges = append(cronRanges, convertRangeToString(leastRanges[i].LongestRange[0], leastRanges[i].LongestRange[len(leastRanges[i].LongestRange)-1], leastRanges[i].Step, maxRange))
	}

	return strings.Join(cronRanges, ",")
}

func FindLeastCronRanges(sortedValues []int) []GreedyRangeSplit {
	if len(sortedValues) == 0 {
		return []GreedyRangeSplit{}
	}

	maxPossibleStep := sortedValues[len(sortedValues)-1] - sortedValues[0]
	bestRangeSplit := GreedyRangeSplit {
		LongestRange: []int{sortedValues[0]},
		RemainingValues: sortedValues[1:],
		NumberOfRanges: len(sortedValues),
		Step: 1,
	}
	for step := 1; step <= maxPossibleStep; step++ {
		rangeSplitforStep := greedyRangesSplit(sortedValues, step)
		if rangeSplitforStep.NumberOfRanges < bestRangeSplit.NumberOfRanges || (rangeSplitforStep.NumberOfRanges == bestRangeSplit.NumberOfRanges && len(rangeSplitforStep.LongestRange) > len(bestRangeSplit.LongestRange)) {
			bestRangeSplit = rangeSplitforStep
		}
	}

	return append([]GreedyRangeSplit{bestRangeSplit}, FindLeastCronRanges(bestRangeSplit.RemainingValues)...)
}



func greedyRangesSplit(sortedValues []int, step int) GreedyRangeSplit {
	firstMismatch := -1
	longestRange := 1
	longestRangeStart := 0
	curRange := 1
	curRangeStart := 0
	numberOfRanges := 1
	prev := sortedValues[0]
	for i := 1; i < len(sortedValues); i++ {
		if sortedValues[i]-step == prev {
			prev = sortedValues[i]
			curRange ++
			if curRange > longestRange {
				longestRange = curRange
				longestRangeStart = curRangeStart
			}
		} else if firstMismatch == -1{
			firstMismatch = i
		}
		if i+1 == len(sortedValues) && firstMismatch != -1{
			numberOfRanges ++
			i = firstMismatch+1
			curRangeStart = firstMismatch
			curRange = 1
			prev = sortedValues[firstMismatch]
			firstMismatch = -1
		}
	}

	longestRangeValuesMap := make(map[int]bool)
	for i:= 0; i < longestRange; i++ {
		nextNumberInLongestRange := sortedValues[longestRangeStart]+i*step
		longestRangeValuesMap[nextNumberInLongestRange] = true
	}
	longestRangeValues := []int{}
	remainingValues := []int{}
	for i := 0; i < len(sortedValues); i++ {
		if longestRangeValuesMap[sortedValues[i]] {
			longestRangeValues = append(longestRangeValues, sortedValues[i])
		} else {
			remainingValues = append(remainingValues, sortedValues[i])
		}
	}

	return  GreedyRangeSplit {
		LongestRange: longestRangeValues,
		RemainingValues: remainingValues,
		NumberOfRanges: numberOfRanges,
		Step: step,
	}
}

func convertRangeToString(rangeStart int, rangeEnd int, step int, maxRange Range) string {
			rangeStr := ""
			isOneNumber := rangeStart == rangeEnd
			hasStepAndMaxEnd := rangeEnd == maxRange.End && step > 1
			canEliminateEnd := hasStepAndMaxEnd || isOneNumber
			if rangeStart == maxRange.Start && rangeEnd == maxRange.End {
				rangeStr = "*"
			} else if canEliminateEnd  {
				rangeStr = strconv.Itoa(rangeStart)
			} else {
				rangeStr = fmt.Sprintf("%d-%d", rangeStart, rangeEnd)
			}

			if step > 1 {
				rangeStr = fmt.Sprintf("%s/%d", rangeStr, step)
			}

			return rangeStr
}
