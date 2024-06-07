package cmd

import (
	"fmt"
	"log"
	_ "time"

	"encoding/json"

	"github.com/lokeshllkumar/pswd-cli/models"
	"github.com/lokeshllkumar/pswd-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "The 'get' subcommand will display the stored password for the specified service: username pair or all the stored passwords for a specific service that exists in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			service, _ := cmd.Flags().GetString("service")
			username, _ := cmd.Flags().GetString("username")

			db, err := models.InitDB()
			if err != nil {
				fmt.Println("Error, database connection failed")
				return
			}

			data, err := db.GetPassword(service, username)
			if err != nil {
				fmt.Println("Error, data could not be fetched")
				return
			}

			var res []models.PasswordEntry
			if err := json.Unmarshal([]byte(data), &res); err != nil {
				log.Fatal("Error retrieving data")
			}

			for _, record := range res {
				decryptedPassword, err := utils.DecryptPassword(record.Password)
				if err != nil {
					fmt.Println("Error decrypting stored password")
					continue
				}

				fmt.Printf("Service:%s |\tUsername:%s |\tPassword:%s\n", record.Service, record.Username, decryptedPassword)
			}

			db.CloseDB()
		} else {
			service, _ := cmd.Flags().GetString("service")

			db, err := models.InitDB()
			if err != nil {
				fmt.Println("Error, database connection failed")
				return
			}

			data, err := db.GetPasswords(service)
			if err != nil {
				fmt.Println("Error, data could not be fetched")
				return
			}

			var res []models.PasswordEntry
			if err := json.Unmarshal([]byte(data), &res); err != nil {
				log.Fatal("Error retrieving data")
			}

			for _, record := range res {
				decryptedPassword, err := utils.DecryptPassword(record.Password)
				if err != nil {
					fmt.Println("Error decrypting stored password")
					continue
				}

				fmt.Printf("Service: %s |\tUsername: %s |\tPassword: %s\n", record.Service, record.Username, decryptedPassword)
			}

			db.CloseDB()
		}
	},
}