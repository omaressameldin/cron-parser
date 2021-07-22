package parser

import (
	"reflect"
	"testing"
)

type parserTest struct {
	timeInput         [5]string
	commandInput      []string
	minutes           []int
	hours             []int
	days              []int
	months            []int
	weekDays          []int
	command           string
	shouldReturnError bool
}

func createParserTestTable() []parserTest {
	return []parserTest{
		{
			timeInput:         [5]string{"*/15", "0", "1,15", "*", "1-5"},
			commandInput:      []string{"/usr/bin/find"},
			shouldReturnError: false,
			minutes:           []int{0, 15, 30, 45},
			hours:             []int{0},
			days:              []int{1, 15},
			months:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			weekDays:          []int{1, 2, 3, 4, 5},
			command:           "/usr/bin/find",
		},
		{
			timeInput:         [5]string{"10-20/4", "*", "*", "*", "*"},
			commandInput:      []string{"/usr/bin/find", "test"},
			shouldReturnError: false,
			minutes:           []int{10, 14, 18},
			hours:             []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			days:              []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			months:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			weekDays:          []int{1, 2, 3, 4, 5, 6, 7},
			command:           "/usr/bin/find test",
		},
		{
			timeInput:         [5]string{"10-20/65", "*", "*", "*", "*"},
			commandInput:      []string{"/usr/bin/find", "test"},
			shouldReturnError: true,
			minutes:           []int{},
			hours:             []int{},
			days:              []int{},
			months:            []int{},
			weekDays:          []int{},
			command:           "",
		},
		{
			timeInput:         [5]string{"10-20/5", "15-40/0", "*", "*", "*"},
			commandInput:      []string{"/usr/bin/find", "test"},
			shouldReturnError: true,
			minutes:           []int{},
			hours:             []int{},
			days:              []int{},
			months:            []int{},
			weekDays:          []int{},
			command:           "",
		},
		{
			timeInput:         [5]string{"10,20,30", "15", "1-5/2", "4", "*"},
			commandInput:      []string{"/usr/bin/find", "test"},
			shouldReturnError: false,
			minutes:           []int{10, 20, 30},
			hours:             []int{15},
			days:              []int{1, 3, 5},
			months:            []int{4},
			weekDays:          []int{1, 2, 3, 4, 5, 6, 7},
			command:           "/usr/bin/find test",
		},
		{
			timeInput:         [5]string{"10", "15", "1", "4/2", "1"},
			commandInput:      []string{"/usr/bin/find", "test"},
			shouldReturnError: false,
			minutes:           []int{10},
			hours:             []int{15},
			days:              []int{1},
			months:            []int{4, 6, 8, 10, 12},
			weekDays:          []int{1},
			command:           "/usr/bin/find asd",
		},
	}
}

func TestParser(t *testing.T) {
	for _, row := range createParserTestTable() {
		res, err := Init(row.timeInput, row.commandInput)
		if row.shouldReturnError && err != nil {
			return
		} else if row.shouldReturnError && err == nil {
			t.Errorf("expected error but got (%v,<nil>)", res)
		} else if !row.shouldReturnError && err != nil {
			t.Errorf(
				"expected output(mins: %v hours: %v days: %v months: %v weekDays: %v, command: %s) but got error %s",
				row.minutes,
				row.hours,
				row.days,
				row.months,
				row.weekDays,
				row.command,
				err.Error(),
			)
		} else if res.GetCommand() != row.command ||
			!reflect.DeepEqual(res.GetMinute(), row.minutes) ||
			!reflect.DeepEqual(res.GetDay(), row.days) ||
			!reflect.DeepEqual(res.GetHour(), row.hours) ||
			!reflect.DeepEqual(res.GetDay(), row.days) ||
			!reflect.DeepEqual(res.GetMonth(), row.months) ||
			!reflect.DeepEqual(res.GetWeek(), row.weekDays) {
			t.Errorf(
				"expected output(mins: %v hours: %v days: %v months: %v weekDays: %v, command: %s) but got error %v",
				row.minutes,
				row.hours,
				row.days,
				row.months,
				row.weekDays,
				row.command,
				res,
			)
		}

	}
}
