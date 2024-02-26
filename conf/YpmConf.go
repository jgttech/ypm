package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/utils"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type YpmConfPackage struct {
	Name    string
	Version string
	Path    string
}

type YpmConf struct {
	LockfileVersion string           `yaml:"lockfileVersion,omitempty"`
	Packages        []YpmConfPackage `yaml:",omitempty"`
}

func (conf *YpmConf) Write(dir string) {
	fileName := env.GetEnv().Lockfile.Name
	confPath := path.Join(dir, fileName)

	utils.WriteYaml(confPath, conf)
}

func (conf *YpmConf) Read(dir string) {
	fileName := env.GetEnv().Lockfile.Name
	confPath := path.Join(dir, fileName)
	exists := utils.PathExists(confPath)

	if exists {
		bytes, err := os.ReadFile(confPath)
		utils.Check(err)

		err = yaml.Unmarshal(bytes, conf)
		utils.Check(err)
	}
}
