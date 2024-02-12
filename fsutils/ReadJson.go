package fsutils

import (
	"encoding/json"
	"jgttech/ypm/errors"
	"jgttech/ypm/utils"
	"os"
)

func ReadJson(filePath string) any {
	exists := utils.PathExists(filePath)

	if !exists {
		errors.PathNotFound(filePath)
	} else {
		file, err := os.ReadFile(filePath)
		utils.Check(err)

		data := make(map[any]any)
		err = json.Unmarshal(file, &data)

		return data
	}

	return nil
}
