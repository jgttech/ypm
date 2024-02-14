package conf

import "jgttech/ypm/utils"

type YpmConfigWorkspace struct {
	Name string
	Path string
}

type YpmConfigYaml struct {
	Workspaces []YpmConfigWorkspace `yaml:",omitempty"`
}

func (conf *YpmConfigYaml) Write(filePath string) {
	exists := utils.PathExists(filePath)

	if !exists {
		utils.WriteYaml(filePath, conf)
	}
}
