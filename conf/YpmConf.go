package conf

type YpmConfWorkspace struct {
	Name    string
	Version string
	Path    string
}

type YpmConf struct {
	Version    string             `yaml:",omitempty"`
	Workspaces []YpmConfWorkspace `yaml:",omitempty"`
}
