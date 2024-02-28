package sys

import (
	"jgttech/ypm/utils"
	"os"
	"path"
	"runtime"
	"strings"
)

type ExecutionContext struct {
	PkgExists  bool
	RepoExists bool
	RepoPath   string
	Cwd        string
	Sep        string
	Conf       *ExecutionConf
}

func NewExecutionContext() *ExecutionContext {
	cwd, err := os.Getwd()
	utils.Check(err)

	conf := NewExecutionConf()
	sep := "/"

	if runtime.GOOS == "windows" {
		sep = "\\"
	}

	tokens := strings.Split(cwd, sep)
	repoPath := cwd

	for idx := range tokens {
		segments := tokens[0 : len(tokens)-idx]
		subDir := strings.Join(segments, sep)

		if subDir == "" {
			subDir = sep
		}

		entries, err := os.ReadDir(subDir)
		utils.Check(err)

		for _, entry := range entries {
			fileName := entry.Name()

			if fileName == conf.Repo.Name {
				repoPath = subDir
				break
			}
		}
	}

	pkgFilePath := path.Join(repoPath, conf.Package.Name)
	repoFilePath := path.Join(repoPath, conf.Repo.Name)

	pkgExists := utils.PathExists(pkgFilePath)
	repoExists := utils.PathExists(repoFilePath)

	ctx := &ExecutionContext{
		PkgExists:  pkgExists,
		RepoExists: repoExists,
		RepoPath:   repoPath,
		Cwd:        cwd,
		Sep:        sep,
		Conf:       conf,
	}

	return ctx
}
