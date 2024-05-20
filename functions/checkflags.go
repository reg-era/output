package ascciartoutput

import (
	"fmt"
	"os"
	"strings"
)

func CheckFlag(arg string) string {
	file_Name := ""
	if strings.HasPrefix(arg, "--output") {
		output := "--output="
		// if !strings.HasPrefix(arg, output) {
		// 	return ""
		// }
		file_Name = arg[len(output):]
		if !strings.HasSuffix(file_Name, ".txt") {
			fmt.Println("Error message")
			os.Exit(0)
		}
	}
	return file_Name
}
