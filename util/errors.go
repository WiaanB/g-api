package util

import "log"

func FatalErrorWrapper(err error, reason string) {
	if err != nil {
		log.Fatalf("%s due to %s. Exiting...\n", reason, err.Error())
	}
}

func ErrorWrapper(err error, reason string) {
	if err != nil {
		log.Printf("%s due to %s\n", reason, err.Error())
	}
}
