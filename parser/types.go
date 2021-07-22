package parser

type Parser struct {
	minute  []int
	hour    []int
	day     []int
	month   []int
	week    []int
	command string
}

type Range struct {
	Start int
	End   int
}

type RunTime struct {
	timeValue string
	step      int
	hasStep   bool
}

const minArgsLength = 6

var minuteRange = Range{0, 59}
var hourRange = Range{0, 23}
var dayRange = Range{1, 31}
var weekRange = Range{1, 7}
var monthRange = Range{1, 12}
