package utils

import (
	"os"
)

func Cwd() string {
	cwd, err := os.Getwd()
	Check(err)

	return cwd
}
