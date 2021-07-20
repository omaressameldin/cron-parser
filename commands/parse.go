package commands

import (
	"log"
	"strings"
	"utils"

	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parses a cron task and shows times it will run",
	Run:   parseCron,
}

func init() {
	RootCmd.AddCommand(parseCmd)
}

func parseCron(cmd *cobra.Command, args []string) {
	utils.Must(utils.ValidateMinLength(minArgsLength, args))

	rows := utils.EqualizeStringsSizes([]string{
		"minute:",
		"hour:",
		"day of month:",
		"month:",
		"day of week:",
		"command:",
	})
	// taking all of args except command (1 or more arg) as command does not need adjusting
	outputs := createRanges(args[:minArgsLength-1])
	outputs = append(outputs, createCommandFromArgs(args[minArgsLength-1:]))

	printExpandedTable(rows, outputs)
}

func printExpandedTable(rows []string, outputs []string) {
	for i := 0; i < len(rows); i++ {
		log.Printf("%s %s\n", rows[i], outputs[i])
	}
}

func createCommandFromArgs(command []string) string {
	return strings.Join(command, " ")
}

func createRanges(rangeValues []string) []string {
	outputs := make([]string, minArgsLength-1)

	return outputs
}
