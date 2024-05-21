package funcs

func PrintOutput(getInput []string, readfile map[rune][]string) string {
	var Content string
	isappaneded := false
	for index, word := range getInput {
		if word != "" {
			i := 0
			for i < 8 {
				for _, char := range word {
					line := readfile[char][i]
					Content += line
				}
				Content += "\n"
				i++
				isappaneded = true
			}
		} else {
			if index == len(getInput)-1 && !isappaneded {
				continue
			} else {
				Content += "\n"
			}
		}
	}
	return Content
}
