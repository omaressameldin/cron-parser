package parser

import (
	"fmt"
	"strings"
)

func Init(timeValues [5]string, command []string) (*Parser, error) {
	minute, err := parseRange(timeValues[0], minuteRange)
	if err != nil {
		return nil, createError(timeValues[0], err)
	}
	hour, err := parseRange(timeValues[1], hourRange)
	if err != nil {
		return nil, createError(timeValues[1], err)
	}
	day, err := parseRange(timeValues[2], dayRange)
	if err != nil {
		return nil, createError(timeValues[2], err)
	}
	month, err := parseRange(timeValues[3], monthRange)
	if err != nil {
		return nil, createError(timeValues[3], err)
	}
	week, err := parseRange(timeValues[4], weekRange)
	if err != nil {
		return nil, createError(timeValues[4], err)
	}

	parsedCommand, err := createCommand(command)
	if err != nil {
		return nil, createError(command, err)
	}
	return &Parser{
		minute:  minute,
		hour:    hour,
		day:     day,
		week:    week,
		month:   month,
		command: parsedCommand,
	}, nil
}

func parseRange(rangeValue string, rng Range) ([]int, error) {
	timeValues, err := getTimeValues(rangeValue, rng)
	if err != nil {
		return nil, err
	}

	return timeValues, nil
}

func (p *Parser) GetMinute() []int {
	return p.minute
}

func (p *Parser) GetHour() []int {
	return p.hour
}

func (p *Parser) GetDay() []int {
	return p.day
}

func (p *Parser) GetWeek() []int {
	return p.week
}

func (p *Parser) GetMonth() []int {
	return p.month
}

func (p *Parser) GetCommand() string {
	return p.command
}

func createCommand(command []string) (string, error) {
	if len(command) == 0 {
		return "", fmt.Errorf("command is empty")
	}

	return strings.Join(command, " "), nil
}
