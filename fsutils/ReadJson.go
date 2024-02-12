package fsutils

import (
	"encoding/json"
	"jgttech/ypm/errors"
	"jgttech/ypm/utils"
	"os"

	"github.com/mitchellh/mapstructure"
)

func ReadJson[T any](filePath string) *T {
	exists := utils.PathExists(filePath)

	if !exists {
		errors.PathNotFound(filePath)
	} else {
		var result T
		file, err := os.ReadFile(filePath)
		utils.Check(err)

		data := make(map[string]any)
		err = json.Unmarshal(file, &data)

		mapstructure.Decode(data, &result)

		return &result
	}

	return nil
}
