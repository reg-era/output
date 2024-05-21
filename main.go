package main

import (
	"fmt"
	"os"

	ascciartoutput "ascciartoutput/functions"
)

func main() {
	/*****Function the Check The Arguments*****/
	args := os.Args[1:]
	if len(args) > 3 || len(args) == 0 {
		ascciartoutput.Error("You shoud enter one or two or three arguments\n")
	}
	/*****Function Return File Name The Flag*****/
	file_Name := ascciartoutput.CheckFlag(args[0])
	/*****Function the Is Output*****/
	if file_Name != "" {
		ascciartoutput.IsOutPut(args, file_Name)
	} else {
		//*****if it's ascii-art or fs parametre*****/
		fmt.Print(ascciartoutput.AsciiArtFsParametre(args))
	}
}
