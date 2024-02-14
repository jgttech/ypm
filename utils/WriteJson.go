package utils

import (
	"encoding/json"
	"log"
	"os"
)

func WriteJson(file string, data any) {
	exists := PathExists(file)

	if exists {
		log.Fatal("File already exists: " + "\"" + file + "\"")
	} else {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		Check(err)

		os.WriteFile(file, jsonData, os.FileMode(0644))
	}
}
