package cmd

import (
	"fmt"

	lottery "github.com/julianvilas/lottery-go"
	"github.com/spf13/cobra"
)

// check represents the check command
var checkCmd = &cobra.Command{
	Use:   "check <number>",
	Short: "Checks if number has been awarded",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("incorrect number of args, want 1, got %v", len(args))
		}

		n, err := stringToInt(args[0])
		if err != nil {
			return fmt.Errorf("not a number: %v", args[0])
		}

		res, err := lottery.CheckNumbers(n...)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", res)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
