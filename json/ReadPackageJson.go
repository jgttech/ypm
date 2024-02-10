package json

import (
	"encoding/json"
	"jgttech/ypm/models"
	"jgttech/ypm/utils"
	"os"
	"path/filepath"
)

func ReadPackageJson(dirPath string) models.PackageJson {
	data, err := os.ReadFile(filepath.Join(dirPath, "package.json"))
	utils.Error(err)

	jsonData := models.PackageJson{}

	err = json.Unmarshal(data, &jsonData)

	return jsonData
}
