package conf

import (
	"jgttech/ypm/env"
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
		fileName := env.GetEnv().Repofile.Name
		utils.WriteJson(path.Join(dir, fileName), conf)
	}
}
