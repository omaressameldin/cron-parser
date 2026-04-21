package parser

import (
	"fmt"
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

func FindBestCronIntConversion(values []int, maxRange Range) string {
	if len(values) == 0 {
		return "*"
	}
	sortedValues := sort.IntSlice(values)

	ranges := []string{}
	rangeStart := sortedValues[0]
	prevValue := rangeStart
	step := 0
	for i := 1; i < len(sortedValues); i++ {
		curValue := sortedValues[i]
		if curValue == prevValue {
			continue
		}

		if step == 0 {
			step = curValue - prevValue
		} else if curValue-prevValue != step {
			ranges = append(ranges, convertRangeToString(rangeStart, prevValue, step, maxRange))
			rangeStart = curValue
			step = 0
		}
		prevValue = curValue
	}
	ranges = append(ranges, convertRangeToString(rangeStart, prevValue, step, maxRange))

	return strings.Join(ranges, ",")
}

func convertRangeToString(rangeStart int, rangeEnd int, step int, maxRange Range) string {
			rangeStr := ""
			if rangeStart == maxRange.Start && rangeEnd == maxRange.End {
				rangeStr = "*"
			} else if rangeEnd == maxRange.End || rangeStart == rangeEnd  {
				rangeStr = strconv.Itoa(rangeStart)
			} else {
				rangeStr = fmt.Sprintf("%d-%d", rangeStart, rangeEnd)
			}

			if step > 1 {
				rangeStr = fmt.Sprintf("%s/%d", rangeStr, step)
			}

			return rangeStr
}
