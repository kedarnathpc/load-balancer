package utils

import (
	"fmt"
	"os"
)

// HandleErr handles errors by printing the error message and exiting the program.
func HandleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
