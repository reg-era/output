package ascciartoutput

import (
	"fmt"
	"os"
	"strings"
)

func CheckCharacter(text string) []string {
	for _, char := range text {
		if rune(char) < ' ' || rune(char) > '~' {
			Error("Please enter a valid character\n")
		}
	}
	return strings.Split(text, "\\n")
}

func Error(specification string) {
	fmt.Println(specification + "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}
