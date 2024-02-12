package conf

import (
	"jgttech/ypm/errors"
	"jgttech/ypm/fsutils"
	"jgttech/ypm/utils"
	"path"
)

type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         string            `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func (conf PackageJson) Write(dir string) {
	exists := utils.PathExists(dir)

	if !exists {
		errors.PathNotFound(dir)
	} else {
		confPath := path.Join(dir, "package.json")
		fsutils.WriteJson(confPath, conf)
	}
}

func (conf PackageJson) Read(dir string) *PackageJson {
	exists := utils.PathExists(dir)

	if !exists {
		errors.PathNotFound(dir)
	} else {
		confPath := path.Join(dir, "package.json")
		jsonData := fsutils.ReadJson(confPath)

		if jsonData != nil {
			return jsonData.(*PackageJson)
		}
	}

	return nil
}
