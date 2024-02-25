package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/utils"
	"path"
)

type ExecutionContext struct {
	PkgExists   bool
	YpmExists   bool
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

	pkgExists := utils.PathExists(path.Join(cwd, env.Repofile.Name))
	ypmExists := utils.PathExists(path.Join(cwd, env.Lockfile.Name))

	etx.RepoPath = repoPath
	etx.WorkingPath = cwd
	etx.Env = env
	etx.Config = ypmConf
	etx.PkgExists = pkgExists
	etx.YpmExists = ypmExists
}
