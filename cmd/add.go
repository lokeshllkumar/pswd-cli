package cmd

import (
	"fmt"

	"github.com/lokeshllkumar/pswd-cli/models"
	"github.com/lokeshllkumar/pswd-cli/utils"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "The 'add' subcommand will add a new username: password pair for a specific service to the database",
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		encryptedPassword, err := utils.EncryptPassword(password)
		if err != nil {
			fmt.Println("Error encrypting password", err)
			return
		}

		db, err := models.InitDB()
		if err != nil {
			fmt.Println("Error, database connection failed", err)
			return
		}

		if err := db.AddPassword(service, username, encryptedPassword); err != nil {
			fmt.Println("Error adding password: ", err)
			return
		}

		db.CloseDB()
	},
}

