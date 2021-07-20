package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:  "cron-parser",
	Long: "a GOlang cli to parse a cron string to show the times at which it will run",
}
