package funcs

import (
	"fmt"
	"os"
	"strings"
)

func GetInput(input string) []string {
	var splitedInput []string
	for _, runes := range input {
		if runes < 32 || runes > 126 {
			fmt.Println("invalid input")
			os.Exit(1)
		}
		splitedInput = strings.Split(input, "\\n")
	}
	return splitedInput
}
