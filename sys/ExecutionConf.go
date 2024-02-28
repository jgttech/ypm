package sys

type packageConf struct {
	Name    string
	Version string
}

type repoConf struct {
	Name    string
	Version string
}

type ExecutionConf struct {
	Sep     string
	Package packageConf
	Repo    repoConf
}

// Returns a reference to a new instance of the
// execution configuration. This configuration is
// something that might live in a configuration file.
// However, a configuration file has to exist on the
// file system, and this approach builds the environment
// configuration into the binary.
//
// Keeping the configuration separate from the context
// to ensures that, if changes do happen, that the
// configuration can be managed separately from the
// execution context.
func NewExecutionConf() *ExecutionConf {
	return &ExecutionConf{
		Package: packageConf{
			Name:    "package.json",
			Version: "0.0.1",
		},
		Repo: repoConf{
			Name:    "repository.lock",
			Version: "1.0.0",
		},
	}
}
