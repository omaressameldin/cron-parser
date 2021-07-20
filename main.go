package main

import (
	"commands"
	"utils"
)

func main() {
	utils.Must(commands.RootCmd.Execute())
}
