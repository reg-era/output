package ascciartoutput

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) map[rune][]string {
	file, err := os.Open(fileName)
	if err != nil {
		Error("Check your file")
	}
	defer file.Close()
	data := map[rune][]string{}
	scaner := bufio.NewScanner(file)
	i := 31
	for scaner.Scan() {
		if scaner.Text() == "" {
			i++
			continue
		}
		data[rune(i)] = append(data[rune(i)], scaner.Text())
		if i == 127 {
			break
		}
	}
	return data
}
