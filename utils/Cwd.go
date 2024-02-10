package utils

import (
	"log"
	"os"
)

func Cwd() string {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	return cwd
}
