package exceptions

import "log"

func UnknownError(err error) {
	log.Fatal("Oops! Looks like something went wrong :( ...", err)
}
