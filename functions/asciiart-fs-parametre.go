package ascciartoutput

import (
	"log"
)

func AsciiArtFsParametre(args []string) string {
	text := args[0]
	readfile := map[rune][]string{}
	if len(args) == 1 {
		readfile = ReadFile("./sources/standard.txt")
		// if it's three arguments
	} else if len(args) == 2 && (args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		readfile = ReadFile("./sources/" + args[1] + ".txt")
	} else {
		log.Fatalln("err: the third argument should be one of these file names (standard),(shadow),(thinkertoy) 0001111")
	}
	// check if the input is among ascii manual

	checkcharacter := CheckCharacter(text)

	// go print my argument
	result := FindAndPrint(checkcharacter, readfile)
	return result
}
