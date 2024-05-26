/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

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

			log.Fatalf("Error", err)

		}

		conn, err := net.Dial("tcp", "localhost:65432")

		if err != nil {

			return fmt.Errorf("Error connecting to server not able to connect")
		}
		defer conn.Close()

		_, err = conn.Write(jsonData)

		if err != nil {
			return fmt.Errorf("Error sendind data: %v", err)
		}

		fmt.Println(string(jsonData))

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)

		if err != nil {

			fmt.Println("Error", err)
			return nil

		}

		response := string(buffer[:n])
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
}
