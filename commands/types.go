package commands

type Range struct {
	Start int
	End   int
}

const minArgsLength = 6

var minuteRange = Range{0, 59}
var hourRange = Range{0, 23}
var dayRange = Range{1, 31}
var weekRange = Range{1, 7}
var monthRange = Range{1, 12}
