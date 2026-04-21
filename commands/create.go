package commands

import (
	"log"
	"parser"
	"utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const createCmdArgsLength = 0



var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a cron string given minute hour day month day of month, month and command",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Must([]error{utils.ValidateMinLength(createCmdArgsLength, args)})
		flags := cmd.Flags()

		parser := parseFlags(flags)

		log.Println(parser.GenerateCronString())
	},
}

func init() {
	createCmd.Flags().String("command", "c", "Command run by cron")
	createCmd.Flags().IntSlice("hours", nil, "hours to run the command")
	createCmd.Flags().IntSlice("minutes", nil, "minutes to run the command")
	createCmd.Flags().IntSlice("days", nil, "days of month to run the command")
	createCmd.Flags().IntSlice("months", nil, "months to run the command")
	createCmd.Flags().IntSlice("weekDays", nil, "days of week to run the command")
	RootCmd.AddCommand(createCmd)
}

func createCron(cmd *cobra.Command, args []string) {

}

func parseFlags(flags *pflag.FlagSet) *parser.Parser {
	var errors []error
	cmd, err := flags.GetString("command")
	if err != nil {
		errors = append(errors, err)
	}

	hour, err := flags.GetIntSlice("hours")
	if err != nil {
		errors = append(errors, err)
	}

	day, err := flags.GetIntSlice("days")
	if err != nil {
		errors = append(errors, err)
	}
	month, err := flags.GetIntSlice("months")
	if err != nil {
		errors = append(errors, err)
	}
	weekDays, err := flags.GetIntSlice("weekDays")
	if err != nil {
		errors = append(errors, err)
	}
	min, err := flags.GetIntSlice("minutes")
	if err != nil {
		errors = append(errors, err)
	}

	utils.Must(errors)

	data, errors :=  parser.CreateRawParser(min, hour, day, weekDays, month, cmd)
	utils.Must(errors)

	return data
}
