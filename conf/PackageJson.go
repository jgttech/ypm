package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/errors"
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
		fileName := env.GetEnv().Repofile.Name
		confPath := path.Join(dir, fileName)
		utils.WriteJson(confPath, conf)
	}
}

func (conf PackageJson) Read(dir string) *PackageJson {
	exists := utils.PathExists(dir)

	if !exists {
		errors.PathNotFound(dir)
	} else {
		fileName := env.GetEnv().Repofile.Name
		confPath := path.Join(dir, fileName)
		confData := utils.ReadJson[PackageJson](confPath)

		if confData != nil {
			return confData
		}
	}

	return nil
}
