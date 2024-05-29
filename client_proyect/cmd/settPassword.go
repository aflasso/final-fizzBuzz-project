/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"client_proyect/util"
	"fmt"

	"github.com/spf13/cobra"
)

var new_password string

// settPasswordCmd represents the settPassword command
var settPasswordCmd = &cobra.Command{
	Use:   "settPassword",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("settPassword called")

		if new_password == "" {

			return fmt.Errorf("a password is require use -p 'password'")

		}

		password := util.HashPassword(new_password)

		util.Write(password)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(settPasswordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// settPasswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// settPasswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	settPasswordCmd.Flags().StringVarP(&new_password, "password", "p", "", "new password for the system")
}
