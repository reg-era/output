package funcs

import "fmt"

func PrintOutput(getInput []string, readfile map[rune][]string) {
	is_printed := false
	for idx, word := range getInput {
		if word != "" {
			i := 0
			for i < 8 {
				for _, char := range word {
					line := readfile[char][i]
					fmt.Print(line)
				}
				fmt.Println()
				i++
				is_printed = true
			}
		} else {
			if idx == len(getInput)-1 && !is_printed {
				continue
			} else {
				fmt.Println()
			}
		}
	}

}
