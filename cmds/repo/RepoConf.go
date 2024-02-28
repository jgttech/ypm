package repo

import (
	"log"
	"strings"

	"golang.org/x/mod/semver"
)

type RepoConf struct {
	Name    string
	Version string
	Path    string
}

func NewRepoConf(conf string) *RepoConf {
	idx := strings.LastIndex(conf, "@")
	name := conf
	version := "0.0.1"

	if idx != -1 && idx != 0 {
		name = conf[0:idx]
		confVersion := conf[idx+1:]
		isValid := semver.IsValid("v" + confVersion)

		if !isValid {
			log.Fatal("Invalid version. Must be a valid semver format. Did you have an extra \"@\" character somewhere?")
		}

		version = confVersion
	}

	return &RepoConf{
		Name:    name,
		Version: version,
	}
}
