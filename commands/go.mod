module commands

require (
	github.com/spf13/cobra v1.2.1
	parser v0.0.0
	utils v0.0.0
)

replace utils => ../utils

replace parser => ../parser

go 1.16
