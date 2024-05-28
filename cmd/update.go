package cmd

import (
	"fmt"

	"github.com/lokeshllkumar/pswd-cli/models"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "The 'update' subcommand updates the stored password for a username for a specific service",
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		username, _ := cmd.Flags().GetString("username")
		newPassword, _ := cmd.Flags().GetString("newPassword")

		db, err := models.InitDB()
		if err != nil {
			fmt.Println("Error, database connection failed")
			return
		}

		if err := db.UpdatePassword(service, username, newPassword); err != nil {
			fmt.Println("Error updating password")
			return
		}

		db.CloseDB()
	},
}