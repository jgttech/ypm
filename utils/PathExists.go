package utils

import (
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	exists := err == nil

	return exists
}
