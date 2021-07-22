module cron-parser

require (
	commands v0.0.0
	utils v0.0.0
)

replace commands => ./commands

replace parser => ./parser

replace utils => ./utils

go 1.16
