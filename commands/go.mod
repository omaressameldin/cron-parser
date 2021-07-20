module commands

require (
	github.com/spf13/cobra v1.2.1
	utils v0.0.0
)

replace utils => ../utils

go 1.16
