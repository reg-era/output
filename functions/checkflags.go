package ascciartoutput

import (
	"strings"
)

func CheckFlag(arg string) string {
	file_Name := ""
	if strings.HasPrefix(arg, "--output=") {
		file_Name = arg[9:]
		if !strings.HasSuffix(file_Name, ".txt") || file_Name == ".txt" || strings.Contains(file_Name, "/") {
			Error("Check your file option\n")
		}
	}
	return file_Name
}
