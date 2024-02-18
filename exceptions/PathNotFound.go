package exceptions

import "log"

func PathNotFound(path string) {
	log.Fatal("Path not found: " + "\"" + path + "\"")
}
