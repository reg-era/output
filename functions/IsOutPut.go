package ascciartoutput

import (
	"os"
	"strings"
)

func IsOutPut(args []string, file_Name string) {
	readFile := map[rune][]string{}
	text := CheckCharacter(args[1])

	if !strings.HasSuffix(args[2], ".txt") {
		readFile = ReadFile("./sources/" + args[2] + ".txt")
	} else {
		Error("The banner should be one of these file names [standard] [shadow] [thinkertoy]\n")
	}

	result := FindAndPrint(text, readFile)
	file, err := os.Create(file_Name)
	if err != nil {
		Error("Check your file")
	}
	defer file.Close()
	file.WriteString(result)
}
