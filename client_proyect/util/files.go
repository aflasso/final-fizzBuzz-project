package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Write(password string) {

	file, err := os.Create("files/password.txt")

	if err != nil {

		fmt.Println("error creating or opening file:", err)
		return

	}

	defer file.Close()

	_, err = io.WriteString(file, password)

	if err != nil {

		fmt.Println("error at writing on the file: ", err)
		return

	}

}

func Read() string {

	file, err := os.Open("files/password.txt")

	if err != nil {

		fmt.Println("error opening the file:", err)
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	first_line := scanner.Text()

	return first_line
}
