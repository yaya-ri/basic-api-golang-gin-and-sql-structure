package common

import "log"

func Error(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
