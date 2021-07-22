package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

func createError(value interface{}, err error) error {
	return fmt.Errorf("invalid value %v, reason: %s", value, err.Error())
}

func getTimeValues(value string, rng Range) ([]int, error) {
	valuesArr := strings.Split(value, "/")
	if len(valuesArr) == 0 {
		return nil, fmt.Errorf("value is empty")
	}
	if valuesArr[0] == "*" {
		timeValues, err := utils.CreateArrFrom(rng.Start, rng.End)
		if err != nil {
			return nil, err
		}

		return timeValues, nil
	}
	newRange, err := getRangeIfRangeValue(valuesArr[0], rng)
	if err != nil {
		return nil, err
	}
	if newRange != nil {
		timeValues, err := utils.CreateArrFrom(newRange.Start, newRange.End)
		if err != nil {
			return nil, err
		}
		return timeValues, nil
	}

	return getCommaSeparatedValues(valuesArr[0], rng)
}

func getRangeIfRangeValue(value string, oldRange Range) (*Range, error) {
	rangeValue := strings.Split(value, "-")
	if len(rangeValue) == 1 {
		return nil, nil
	} else if len(rangeValue) > 2 {
		return nil, fmt.Errorf("too many - found")
	}

	start, err := strconv.Atoi(rangeValue[0])
	if err != nil {
		return nil, err
	}
	if start < oldRange.Start {
		return nil, fmt.Errorf("values is out of range %v", oldRange)
	}

	end, err := strconv.Atoi(rangeValue[1])
	if err != nil {
		return nil, err
	}
	if end > oldRange.End {
		return nil, fmt.Errorf("value is out of range %v", oldRange)
	}

	return &Range{
		Start: start,
		End:   end,
	}, nil
}

func getCommaSeparatedValues(value string, oldRange Range) ([]int, error) {
	values := strings.Split(value, ",")
	valuesMap := make(map[int]bool)

	for _, value := range values {
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		if valuesMap[parsedValue] {
			return nil, fmt.Errorf("duplicate value found")
		}
		if parsedValue < oldRange.Start || parsedValue > oldRange.End {
			return nil, fmt.Errorf("value is out of range %v", oldRange)
		}
		valuesMap[parsedValue] = true
	}

	parsedValues := make([]int, 0, len(valuesMap))
	for val := range valuesMap {
		parsedValues = append(parsedValues, val)
	}
	//sorting because hashmap will not preserve insertion order, linkedHashMap can be used instead.
	sort.Ints(parsedValues)
	return parsedValues, nil
}

func getStep(value string, rng Range) (int, error) {
	timeValues := strings.Split(value, "/")
	if len(timeValues) > 2 {
		return 0, fmt.Errorf("multiple / found")
	}

	if len(timeValues) < 2 {
		return 1, nil
	}

	step, err := strconv.Atoi(timeValues[1])
	if err != nil {
		return 0, err
	}
	if step < 1 || step > rng.End {
		return 0, fmt.Errorf("step value is out of range: 1 -> %d", rng.End)
	}

	return step, nil
}
