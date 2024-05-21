package main

import (
	"fmt"
	"os"
	"strings"

	"funcs/funcs"
)

func main() {
	var filepath string
	var getInput []string
	// checkening if the user entered more than 4 or less than 2 args
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
	// looking for the banner and making it the path and deleting the last arg (the banner)
	if os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "shadow" || os.Args[len(os.Args)-1] == "thinkertoy" {
		filepath = ("./banners/" + os.Args[len(os.Args)-1] + ".txt")
		os.Args = os.Args[:len(os.Args)-1]
		// if the last arg isn't a banner the path would be by default standard
	} else if os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy" && !strings.HasSuffix(os.Args[len(os.Args)-1], ".txt") {
		filepath = ("./banners/standard.txt")
	}
	readfile := funcs.Readfile(filepath)
	if len(os.Args) == 2 {

		getInput = funcs.GetInput(os.Args[1])
		fmt.Print(funcs.PrintOutput(getInput, readfile))

	} else if len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--output=") {
		getInput = funcs.GetInput(os.Args[2])
		filename := os.Args[1]
		file, err := os.Create(filename[9:])
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			os.Exit(0)
		}
		output := funcs.PrintOutput(getInput, readfile)
		file.Write([]byte(output))
		if err != nil {
			fmt.Println("error")
		}
		defer file.Close()

	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
}
