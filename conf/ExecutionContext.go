package conf

import (
	"jgttech/ypm/env"
	"jgttech/ypm/utils"
	"path"
	"strings"

	"github.com/yargevad/filepathx"
)

type ExecutionContext struct {
	PkgExists   bool
	YpmExists   bool
	RepoPath    string
	WorkingPath string
	Env         *env.EnvConf
	Config      *YpmConf
}

func NewExecutionContext() *ExecutionContext {
	instance := &ExecutionContext{}

	// Loads the environment, not the lock file metadata.
	instance.LoadEnv()

	// Detects workspace metadata in the project and ensures
	// that the file configuration data matches the detected
	// repo and workspaces configuration. If something does
	// not match, it updates the configuration files.
	instance.Sync()

	// Loads the configuration file data (i.e. lock file).
	instance.LoadConfig()

	return instance
}

func (etx *ExecutionContext) LoadEnv() {
	env := env.GetEnv()
	cwd := utils.Cwd()
	repoPath := utils.DetectRepoPath(cwd)

	pkgExists := utils.PathExists(path.Join(repoPath, env.Repofile.Name))
	ypmExists := utils.PathExists(path.Join(repoPath, env.Lockfile.Name))

	etx.RepoPath = repoPath
	etx.WorkingPath = cwd
	etx.Env = env
	etx.PkgExists = pkgExists
	etx.YpmExists = ypmExists
}

func (etx *ExecutionContext) LoadConfig() {
	ypmConf := &YpmConf{}
	ypmConf.Read(etx.RepoPath)

	etx.Config = ypmConf
}

func (etx *ExecutionContext) Sync() {
	ypmPackages := []YpmConfPackage{}
	ypmConf := &YpmConf{}
	ypmConf.Read(etx.RepoPath)

	globPath := path.Join(etx.RepoPath, "**", etx.Env.Repofile.Name)
	matches, err := filepathx.Glob(globPath)
	utils.Check(err)

	for _, match := range matches {
		json := utils.ReadJson[PackageJson](match)
		relativePath := strings.Split(strings.Split(match, etx.RepoPath)[1], etx.Env.Sep+etx.Env.Repofile.Name)[0]

		if relativePath == etx.Env.Sep {
			relativePath = "."
		} else {
			relativePath = "." + relativePath
		}

		ypmPackages = append(ypmPackages, YpmConfPackage{
			Name:    json.Name,
			Version: json.Version,
			Path:    relativePath,
		})
	}

	ypmConf.Packages = ypmPackages
	ypmConf.Write(etx.RepoPath)
}

func (etx *ExecutionContext) IsValidPackageName(name string) bool {
	for _, pkg := range etx.Config.Packages {
		if pkg.Name == name {
			return false
		}
	}

	return true
}
