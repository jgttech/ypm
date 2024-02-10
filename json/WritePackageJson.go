package json

import (
	"encoding/json"
	"jgttech/ypm/utils"
	"os"
	"path"
)

func (nvj NameVersionJson) WritePackageJson(basePath string) {
	pkgPath := path.Join(basePath, "package.json")
	file, err := json.MarshalIndent(nvj, "", "  ")
	utils.Check(err)

	os.WriteFile(pkgPath, file, os.FileMode(0644))
}
