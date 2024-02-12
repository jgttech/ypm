package conf

import (
	"jgttech/ypm/errors"
	"jgttech/ypm/fsutil"
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
		fsutil.WriteJson(path.Join(dir, "package.json"), conf)
	}
}
