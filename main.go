package main

import (
	"fmt"
	"os"
	"strings"

	"funcs/funcs"
)

func main() {
	var filepath string
	getInput := []string{}
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
	if os.Args[len(os.Args)-1] == "standard" || os.Args[len(os.Args)-1] == "shadow" || os.Args[len(os.Args)-1] == "thinkertoy" {
		filepath = ("./banners/" + os.Args[len(os.Args)-1] + ".txt")
		os.Args = os.Args[:len(os.Args)-1]
	} else if os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy" && !strings.HasSuffix(os.Args[len(os.Args)-1], ".txt") {
		filepath = ("./banners/standard.txt")
	}
	if len(os.Args) == 2 {
		getInput = funcs.GetInput(os.Args[1])
	}
	if len(os.Args) == 3 {
		getInput = funcs.GetInput(os.Args[2])
		funcs.Writefile(os.Args[1])
	}

	readfile := funcs.Readfile(filepath)
	funcs.PrintOutput(getInput, readfile)
}
