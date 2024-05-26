/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

type Data struct {
	Problem   string
	CantData  int64
	MinNumber int64
	MaxNumber int64
	OutPut    string
	Kill      bool
}

var problem string
var cantData int64
var minNumber int64
var maxNumber int64
var outPut string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client_proyect",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		data := Data{

			Problem:   problem,
			CantData:  cantData,
			MinNumber: minNumber,
			MaxNumber: maxNumber,
			OutPut:    outPut,
			Kill:      false,
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client_proyect.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&problem, "problem", "p", "", "Problem to be resolved")
	rootCmd.Flags().Int64VarP(&cantData, "amount", "a", 10, "amount of numbers")
	rootCmd.Flags().Int64VarP(&minNumber, "min", "x", 10, "minimum number")
	rootCmd.Flags().Int64VarP(&maxNumber, "max", "y", 10, "maximum number")
	rootCmd.Flags().StringVarP(&outPut, "output", "o", "", "output file")
}
