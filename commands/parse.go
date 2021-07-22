package commands

import (
	"log"
	"parser"
	"utils"

	"github.com/spf13/cobra"
)

const minArgsLength = 6

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
	var timeArgs [5]string
	copy(timeArgs[:], args[:minArgsLength-1])

	cronParser, err := parser.Init(timeArgs, args[minArgsLength-1:])
	utils.Must(err)

	printExpandedTable(rows, cronParser)
}

func printExpandedTable(rows []string, cronParser *parser.Parser) {
	separator := " "
	log.Printf("%s %s\n", rows[0], utils.ConvertIntArrToString(cronParser.GetMinute(), separator))
	log.Printf("%s %s\n", rows[1], utils.ConvertIntArrToString(cronParser.GetHour(), separator))
	log.Printf("%s %s\n", rows[2], utils.ConvertIntArrToString(cronParser.GetDay(), separator))
	log.Printf("%s %s\n", rows[3], utils.ConvertIntArrToString(cronParser.GetMonth(), separator))
	log.Printf("%s %s\n", rows[4], utils.ConvertIntArrToString(cronParser.GetWeek(), separator))
	log.Printf("%s %s\n", rows[5], cronParser.GetCommand())
}
