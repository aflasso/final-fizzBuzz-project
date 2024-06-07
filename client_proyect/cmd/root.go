/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"client_proyect/util"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Data struct {
	Problem     string
	CantData    int64
	MinNumber   int64
	MaxNumber   int64
	OutPut_file bool
	OutPut_cmd  bool
	Kill        bool
	TestMode    bool
}

type Result struct {
	Result []string
	Error  string
}

var problem string
var cantData int64
var minNumber int64
var maxNumber int64
var outPut_file bool
var outPut_cmd bool
var testMode bool

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

			Problem:     problem,
			CantData:    cantData,
			MinNumber:   minNumber,
			MaxNumber:   maxNumber,
			OutPut_file: outPut_file,
			OutPut_cmd:  outPut_cmd,
			Kill:        false,
			TestMode:    testMode,
		}

		// if testMode {
		// 	jsonData, err := json.Marshal(data)
		// 	if err != nil {
		// 		return fmt.Errorf("error marshaling data: %v", err)
		// 	}
		// 	fmt.Fprintln(cmd.OutOrStdout(), string(jsonData))
		// 	// return nil
		// }

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

		var result Result

		err = json.Unmarshal(data_recived, &result)

		if result.Error != "" {
			return fmt.Errorf("error: %v", result.Error)
		}

		if err != nil {
			return fmt.Errorf("error deserializing JSON: %v", err)
		}

		if outPut_file {

			util.Write_by_arr(result.Result, "files/result.txt")

		}

		if outPut_cmd || (!outPut_file && !outPut_cmd) {

			for _, item := range result.Result {

				fmt.Println(item)

			}

		}

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
	rootCmd.Flags().StringVarP(&problem, "problem", "p", "Prime", "Problem to be resolved")
	rootCmd.Flags().Int64VarP(&cantData, "amount", "a", 1000, "amount of numbers")
	rootCmd.Flags().Int64VarP(&minNumber, "min", "x", 0, "minimum number")
	rootCmd.Flags().Int64VarP(&maxNumber, "max", "y", 100, "maximum number")
	rootCmd.Flags().BoolVarP(&outPut_file, "file", "f", false, "output file")
	rootCmd.Flags().BoolVarP(&outPut_cmd, "cmd", "c", false, "output cmd")

	rootCmd.Flags().BoolVar(&testMode, "test-mode", false, "enable test mode") // Añadir esta línea

}
