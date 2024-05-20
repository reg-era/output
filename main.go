package main

import (
	"fmt"
	"os"

	ascciartoutput "ascciartoutput/functions"
)

func main() {
	/*****Function the Check The Arguments*****/
	args := os.Args[1:]
	ascciartoutput.CheckTheArgs(args)

	/*****Function Return File Name The Flag*****/
	file_Name := ascciartoutput.CheckFlag(args[0])
	// fmt.Println(file_Name)

	/*****Function the Is Output*****/
	if file_Name != "" {
		ascciartoutput.IsOutPut(args, file_Name)
	} else {
		//*****if it's ascii-art or fs parametre*****/
		result_2 := ascciartoutput.AsciiArtFsParametre(args)
		fmt.Print(result_2)
	}
}
