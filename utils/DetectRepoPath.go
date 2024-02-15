package utils

import (
	"jgttech/ypm/env"
	"os"
	"runtime"
	"strings"
)

func DetectRepoPath(ctxPath string) string {
	env := env.GetEnv()
	cwd := Cwd()
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
		Check(err)

		for _, entry := range entries {
			fileName := entry.Name()

			if fileName == env.Lockfile.Name {
				repoPath = subDir
				break
			}
		}
	}

	return repoPath
}
