package conf

import (
	"jgttech/ypm/errors"
	"jgttech/ypm/utils"
	"path"
)

type InitPackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (conf InitPackageJson) Write(dir string) {
	exists := utils.PathExists(dir)

	if !exists {
		errors.PathNotFound(dir)
	} else {
		utils.WriteJson(path.Join(dir, "package.json"), conf)
	}
}
