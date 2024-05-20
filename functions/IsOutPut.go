package ascciartoutput

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func IsOutPut(args []string, file_Name string) {
	text := ""
	banner := "standard"
	readFile := map[rune][]string{}

	if len(args) == 3 {
		text = args[1]
		banner = args[2]
	} else if len(args) == 2 {
		text = args[1]
		// banner = args[1]
	} else if len(args) == 1 {
		text = args[1]
	}
	if file_Name != "" {
		// fmt.Println(banner)
		if !strings.HasSuffix(banner, ".txt") {
			readFile = ReadFile("./sources/" + banner + ".txt")
		} else {
			log.Fatalln("err: the third argument should be one of these file names (standard),(shadow),(thinkertoy)")
		}
		// check if the input is among ascii manual
		checkcharacter := CheckCharacter(text)
		fmt.Println(checkcharacter)

		// if os.Args[0] == file_Name {
		if strings.HasSuffix(args[0], file_Name) {
			result := FindAndPrint(checkcharacter, readFile)
			file, err := os.Create(file_Name)
			if err != nil {
				fmt.Println("err")
				os.Exit(0)
			}
			defer file.Close()
			file.WriteString(result)
		}
	}
}
