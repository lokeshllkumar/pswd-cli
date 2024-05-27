package main

import (
	"fmt"
	"os"

	"github.com/lokeshllkumar/pswd-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pswd-cli",
		Short: "A CLI password managing service",
	}

	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.GetCmd)
	rootCmd.AddCommand(cmd.DisplayCmd)
	rootCmd.AddCommand(cmd.UpdateCmd)
	rootCmd.AddCommand(cmd.DeleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}