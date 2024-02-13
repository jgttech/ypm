package utils

import (
	"log"
	"strings"

	"golang.org/x/mod/semver"
)

type PackageArgs struct {
	Name    string
	Version string
}

func ParsePackageArgs(project string) *PackageArgs {
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

	packageArgs := &PackageArgs{
		Name:    name,
		Version: version,
	}

	return packageArgs
}
