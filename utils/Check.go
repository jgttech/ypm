package utils

import "log"

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
