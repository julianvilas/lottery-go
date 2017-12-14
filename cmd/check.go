package cmd

import (
	"fmt"
	"strconv"

	lottery "github.com/julianvilas/lottery-go"
	"github.com/spf13/cobra"
)

// check represents the check command
var checkCmd = &cobra.Command{
	Use:   "check number",
	Short: "Checks if number has been awarded",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("incorrect number of args, want 1, got %v", len(args))
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("not a number: %v", args[0])
		}

		p, err := lottery.CheckNumber(n)
		if err != nil {
			return err
		}

		fmt.Println(p)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
