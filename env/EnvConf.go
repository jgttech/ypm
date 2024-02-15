package env

type envConfRepofile struct {
	Name    string
	Version string
}

type envConfLockfile struct {
	Name    string
	Version string
}

type EnvConf struct {
	Repofile envConfRepofile
	Lockfile envConfLockfile
}

// A static configuration that gets compiled into
// the binary configuring how YPM functions. This
// approach means there is not need to read a config
// file to setup the package manager.
//
// Potentially, in the future it might be better to
// have a user configuration in a $HOME/.ypm directory
// that allows more customization about how YPM works.
func GetEnv() *EnvConf {
	return &EnvConf{
		Repofile: envConfRepofile{
			Name:    "package.json",
			Version: "0.0.1",
		},
		Lockfile: envConfLockfile{
			Name:    "package.lock",
			Version: "1.0.0",
		},
	}
}
