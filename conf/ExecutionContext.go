package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/utils"
)

type ExecutionContext struct {
	RepoPath    string
	WorkingPath string
	Env         *env.EnvConf
	Config      *YpmConf
}

func (etx *ExecutionContext) Load() {
	env := env.GetEnv()
	cwd := utils.Cwd()
	repoPath := utils.DetectRepoPath(cwd)
	ypmConf := &YpmConf{}
	ypmConf.Read(repoPath)

	etx.RepoPath = repoPath
	etx.WorkingPath = cwd
	etx.Env = env
	etx.Config = ypmConf
}
