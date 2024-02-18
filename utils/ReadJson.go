package utils

import (
	"encoding/json"
	"jgttech/ypm/exceptions"
	"os"

	"github.com/mitchellh/mapstructure"
)

func ReadJson[T any](filePath string) *T {
	exists := PathExists(filePath)

	if !exists {
		exceptions.PathNotFound(filePath)
	} else {
		var result T
		file, err := os.ReadFile(filePath)
		Check(err)

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		Check(err)

		err = mapstructure.Decode(data, &result)
		Check(err)

		return &result
	}

	return nil
}
