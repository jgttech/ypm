package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/exceptions"
	"jgttech/ypm/utils"
	"log"
	"path"
	"strings"

	"golang.org/x/mod/semver"
)

type InitPackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (conf *InitPackageJson) Load(project string) {
	idx := strings.LastIndex(project, "@")
	name := project
	version := "0.0.1"

	if idx != -1 && idx != 0 {
		name = project[0:idx]
		confVersion := project[idx+1:]
		isValid := semver.IsValid("v" + confVersion)

		if !isValid {
			log.Fatal("Invalid version. Must be a valid semver format. Did you have an extra \"@\" character somewhere?")
		}

		version = confVersion
	}

	conf.Name = name
	conf.Version = version
}

func (conf *InitPackageJson) Write(dir string) {
	exists := utils.PathExists(dir)

	if !exists {
		exceptions.PathNotFound(dir)
	} else {
		fileName := env.GetEnv().Repofile.Name
		utils.WriteJson(path.Join(dir, fileName), conf)
	}
}
