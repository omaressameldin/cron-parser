module cron-parser

require (
	commands v0.0.0
	github.com/spf13/cobra v1.2.1 // indirect
	utils v0.0.0
)

replace commands => ./commands

replace utils => ./utils

go 1.16
