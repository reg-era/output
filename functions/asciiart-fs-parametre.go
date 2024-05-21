package ascciartoutput


// this func takes arguments as parametre and return the final result to print
func AsciiArtFsParametre(args []string) string {
	text := args[0]
	readfile := map[rune][]string{}
	if len(args) == 1 {
		readfile = ReadFile("./sources/standard.txt")
		// if it's three arguments
	} else if len(args) == 2 && (args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		readfile = ReadFile("./sources/" + args[1] + ".txt")
	} else {
		Error("The banner should be one of these file names [standard] [shadow] [thinkertoy]\n")
	}
	// go print my argument
	result := FindAndPrint(CheckCharacter(text), readfile)
	return result
}
