module cron-parser

require (
	commands v0.0.0
	utils v0.0.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	parser v0.0.0 // indirect
)

replace commands => ./commands

replace parser => ./parser

replace utils => ./utils

go 1.21
