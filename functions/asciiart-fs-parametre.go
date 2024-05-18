package asciiartfs

import (
	"log"
	"os"
	"strings"
)

func AsciiArtFsParametre( text string , banner string , checkcharacter []string, readfile map[rune][]string) string {
	readFile := map[rune][]string{}
	args := os.Args[1:]
	if len(args) == 2 {
		readFile = ReadFile("./sources/" + banner + ".txt")
		// if it's three arguments
	} else if !strings.HasSuffix(banner, ".txt") {
		readFile = ReadFile("./sources/" + banner + ".txt")
	} else {
		log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
	}
	// check if the input is among ascii manual

	// 	asciiartfs "asciiartfs/functions"
	checkcharacter = CheckCharacter(text)

	// go print my argument
	result := FindAndPrint(checkcharacter, readFile)
	return result
}
