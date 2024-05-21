package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func Readfile(filename string) map[rune][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
	defer file.Close()
	output := make(map[rune][]string)
	s := bufio.NewScanner(file)
	i := 31
	line := 0
	for s.Scan() {
		if s.Text() == "" {
			i++
			continue
		}
		output[rune(i)] = append(output[rune(i)], s.Text())
		line++
		if i == 127 {
			break
		}
	}
	return output
}
