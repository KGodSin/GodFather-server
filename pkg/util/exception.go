package util

import "log"

func ErrorHandler(e error, eType string) {
	if e != nil {
		switch eType {
		case "panic":
			panic(e)
		case "fatal":
			log.Fatal(e)
		case "fatalln":
			log.Fatalln(e)
		}
	}
}
