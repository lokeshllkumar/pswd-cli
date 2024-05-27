package cmd

import (
	"fmt"

	"github.com/lokeshllkumar/pswd-cli/models"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command {
	Use:   "Delete",
	Short: "The 'delete' subcommand will delete an existing username: password pair for a specific service or all exisiting entries for a service",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			service, _ := cmd.Flags().GetString("service")
			username, _ := cmd.Flags().GetString("username")

			db, err := models.InitDB()
			if err != nil {
				fmt.Println("Error, database connection failed")
				return
			}

			if err := db.DeletePassword(service, username); err != nil {
				fmt.Println("Error deleting corresponding entry")
				return
			}
		} else if len(args) == 1 {
			service, _ := cmd.Flags().GetString("service")

			db, err := models.InitDB()
			if err != nil {
				fmt.Println("Error, database connection failed")
				return
			}

			if err := db.DeletePasswords(service); err != nil {
				fmt.Println("Error deleting corresponding entry/entries")
				return
			}
		} else {
			fmt.Println("Error, invalid number of arguments provided")
			return
		}
	},
}