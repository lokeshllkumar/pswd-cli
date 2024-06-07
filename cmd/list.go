package cmd

import (
	"fmt"
	"log"
	_"time"

	"encoding/json"

	"github.com/lokeshllkumar/pswd-cli/models"
	"github.com/lokeshllkumar/pswd-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "The 'list' subcommand will display all stored password records that exist in the database",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := models.InitDB()
		if err != nil {
			fmt.Printf("Error, database connection failed: %v", err)
			return
		}

		data, err := db.ListPasswords()
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
	},
}
