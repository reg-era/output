package ascciartoutput

func FindAndPrint(checkcharacter []string, readfile map[rune][]string) string {
	result := ""
	is_Printed := false
	for idx, word := range checkcharacter {
		if word != "" {
			for i := 0; i < 8; i++ {
				line := ""
				for _, char := range word {
					line = readfile[char][i]
					result += line
				}
				result += "\n"
				is_Printed = true
			}
		} else {
			if idx != len(checkcharacter)-1 || is_Printed {
				result += "\n"
			}
		}
	}
	return result
}
