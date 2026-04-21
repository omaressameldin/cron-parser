package main

import (
	"commands"
	"utils"
)

func main() {
	utils.Must([]error{commands.RootCmd.Execute()})
}
