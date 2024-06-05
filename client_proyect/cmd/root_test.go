package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// mockExecuteCommand is a helper function to execute a Cobra command and capture its output.
func mockExecuteCommand(t *testing.T, cmd *cobra.Command, args []string) (string, error) {
	t.Helper()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	fmt.Println(args)

	cmd.SetArgs(args)

	// Execute the command
	err := cmd.Execute()
	return buf.String(), err
}

func TestRootCmd(t *testing.T) {
	tests := []struct {
		args         []string
		expectedData Data
	}{
		{
			args: []string{"--test-mode"},
			expectedData: Data{
				Problem:     "Prime",
				CantData:    1000,
				MinNumber:   0,
				MaxNumber:   100,
				OutPut_file: false,
				OutPut_cmd:  false,
				Kill:        false,
				TestMode:    true,
			},
		},

		{
			args: []string{"--problem", "Factorial", "--amount", "100", "--min", "0", "--max", "10", "--cmd", "--test-mode"},
			expectedData: Data{
				Problem:     "Factorial",
				CantData:    100,
				MinNumber:   0,
				MaxNumber:   10,
				OutPut_file: false,
				OutPut_cmd:  true,
				Kill:        false,
				TestMode:    true,
			},
		},
		{
			args: []string{"--problem", "Sum", "--amount", "500", "--min", "1", "--max", "1000", "--file", "--cmd=false", "--test-mode"},
			expectedData: Data{
				Problem:     "Sum",
				CantData:    500,
				MinNumber:   1,
				MaxNumber:   1000,
				OutPut_file: true,
				OutPut_cmd:  false,
				Kill:        false,
				TestMode:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			// Execute the command with the test arguments
			output, err := mockExecuteCommand(t, rootCmd, tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Unmarshal the command output into a Data struct
			var outputData Data
			if err := json.Unmarshal([]byte(output), &outputData); err != nil {
				t.Fatalf("error unmarshaling output data: %v", err)
			}

			// Compare the output data with the expected data
			if outputData != tt.expectedData {
				t.Errorf("expected %+v, got %+v", tt.expectedData, outputData)
			}
		})
	}
}
