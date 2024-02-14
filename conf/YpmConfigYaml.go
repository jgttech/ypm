package conf

import (
	"io"
	"jgttech/ypm/utils"
	"os"

	"gopkg.in/yaml.v3"
)

type YpmConfigWorkspace struct {
	Name string
	Path string
}

type YpmConfigYaml struct {
	Workspaces []YpmConfigWorkspace `yaml:",omitempty"`
}

func (conf *YpmConfigYaml) Write(path string) {
	data, err := yaml.Marshal(conf)
	utils.Check(err)

	file, err := os.Create(path)
	utils.Check(err)

	defer file.Close()

	_, err = io.WriteString(file, string(data))
	utils.Check(err)
}
