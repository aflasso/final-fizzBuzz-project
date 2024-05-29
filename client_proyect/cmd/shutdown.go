/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"client_proyect/util"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var input_password string

// shutdownCmd represents the shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("shutdown called")

		if input_password == "" {
			return fmt.Errorf("a password is require use -p 'password'")
		}

		actual_password := util.Read()

		if actual_password == "" {

			return fmt.Errorf("the system doesnt have set a password, use command settPassword")

		}

		if !util.CheckPasswordHash(input_password, actual_password) {

			return fmt.Errorf("incorrect password")

		}

		data := Data{
			Problem:   "",
			CantData:  0,
			MinNumber: 0,
			MaxNumber: 0,
			OutPut:    "",
			Kill:      true,
		}

		jsonData, err := json.Marshal(data)

		if err != nil {

			log.Fatalf("error: %v", err)

		}

		port := "localhost:65432"

		conn, err := util.Start_server(port)

		if err != nil {

			return fmt.Errorf("error connecting to server not able to connect")
		}
		defer conn.Close()

		err = util.Send_data(jsonData, conn)

		if err != nil {
			return fmt.Errorf("error sendind data: %v", err)
		}

		data_recived, err := util.Recive_data(conn)

		if err != nil {

			return fmt.Errorf("error: %v", err)

		}

		response := string(data_recived)
		fmt.Println("Respuesta del servidor:", response)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(shutdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shutdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shutdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	shutdownCmd.Flags().StringVarP(&input_password, "password", "p", "", "the password to autorize shotdown")
}
