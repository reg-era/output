package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	asciiartfs "asciiartfs/functions"
)

func main() {
	text := ""
	checkcharacter := asciiartfs.CheckCharacter(text)
	readFile := map[rune][]string{}
	if len(os.Args) != 2 && len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatalln("err: you shoud enter two or three or foor arguments")
	}
	output := asciiartfs.CheckFlag(os.Args[1])
	args := os.Args[1:]
	banner := "standard"
	// text = ""
	if len(args) == 3 {
		text = args[1]
		banner = args[2]
	} else if len(args) == 2 {
		text = args[0]
		banner = args[1]
	} else if len(args) == 1 {
		text = args[0]
	}
	if output != "" {
		if !strings.HasSuffix(banner, ".txt") {
			readFile = asciiartfs.ReadFile("./sources/" + banner + ".txt")
		} else {
			log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
		}
		// check if the input is among ascii manual
		checkcharacter := asciiartfs.CheckCharacter(text)

		// go print my argument
		result := asciiartfs.FindAndPrint(checkcharacter, readFile)
		file, err := os.Create(output)
		if err != nil {
			fmt.Println("err")
			os.Exit(0)
		}
		defer file.Close()
		file.WriteString(result)
	} else {
		// if it's only two or three arguments 
		result_2 := asciiartfs.AsciiArtFsParametre(text, banner, checkcharacter, readFile)
		fmt.Print(result_2)
	}
	 
}
