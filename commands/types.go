package commands

type Range struct {
	Start int
	End   int
}

const minArgsLength = 6

var minuteRange = Range{0, 59}
var hourRange = Range{0, 23}
var monthRange = Range{0, 12}
var dayRange = Range{0, 6}
