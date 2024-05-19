package funcs

import (
	"fmt"
	"os"
	"strings"
)

func Writefile(fileName string) {
	if strings.HasPrefix(fileName, "--output=") && strings.HasSuffix(fileName, ".txt") {
		fileName = fileName[9:]
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("error: couldn't create a file")
		}
		defer file.Close()
		// file = os.WriteFile(fileName, , 0644)
	}
}
