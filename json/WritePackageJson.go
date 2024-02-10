package json

import (
	"encoding/json"
	"jgttech/ypm/utils"
	"os"
	"path"
)

func WritePackageJson(basePath string, data any) {
	pkgPath := path.Join(basePath, "package.json")
	file, err := json.MarshalIndent(data, "", "  ")
	utils.Error(err)

	os.WriteFile(pkgPath, file, os.FileMode(0644))
}
