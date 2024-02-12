package fsutil

import (
	"encoding/json"
	"jgttech/ypm/utils"
	"log"
	"os"
)

func WriteJson(file string, data any) {
	exists := utils.PathExists(file)

	if exists {
		log.Fatal("File already exists: " + "\"" + file + "\"")
	} else {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		utils.Check(err)

		os.WriteFile(file, jsonData, os.FileMode(0644))
	}
}
