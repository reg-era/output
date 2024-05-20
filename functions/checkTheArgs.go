package ascciartoutput

import (
	"log"
)

func CheckTheArgs(args []string) {
	if len(args) != 1 && len(args) != 2 && len(args) != 3 {
		log.Fatalln("err: you shoud enter two or three or foor arguments")
	}
}
