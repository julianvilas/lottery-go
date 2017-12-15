package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "lottery-go",
	Short: "Checks Spain Christmas Lottery Results using EL PAIS API",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "prints verbose information during command execution")
}

func stringToInt(s ...string) (res []int, err error) {
	for _, str := range s {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("not a number: %v", str)
		}

		res = append(res, n)
	}

	return res, nil
}
