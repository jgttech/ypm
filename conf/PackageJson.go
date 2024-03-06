package conf

type PackageJson struct {
	Name            string            `json:"name,omitempty"`
	Version         string            `json:"version,omitempty"`
	Private         bool              `json:"private,omitempty"`
	Scripts         map[string]string `json:"scripts,omitempty"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
}
