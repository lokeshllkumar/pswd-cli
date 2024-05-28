package cmd

import (
	_ "github.com/spf13/cobra"
)


func init() {
	AddCmd.Flags().StringP("service", "s", "", "Service name")
	AddCmd.Flags().StringP("username", "u", "", "Username")
	AddCmd.Flags().StringP("password", "p", "", "Password")

	GetCmd.Flags().StringP("service", "s", "", "Service name")
	GetCmd.Flags().StringP("username", "u", "", "Username")

	UpdateCmd.Flags().StringP("service", "s", "", "Service name")
	UpdateCmd.Flags().StringP("username", "u", "", "Username")
	UpdateCmd.Flags().StringP("newPassword", "p", "", "New password")

	DeleteCmd.Flags().StringP("service", "s", "", "Service name")
	DeleteCmd.Flags().StringP("username", "u", "", "Username")

	// no flags for the "list" subcommand

	AddCmd.MarkFlagRequired("service")
	AddCmd.MarkFlagRequired("username")
    AddCmd.MarkFlagRequired("password")

	GetCmd.MarkFlagRequired("service")

	UpdateCmd.MarkFlagRequired("service")
	UpdateCmd.MarkFlagRequired("username")
    UpdateCmd.MarkFlagRequired("newPassword")

	DeleteCmd.MarkFlagRequired("service")
}