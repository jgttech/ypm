package conf

type YpmConfigYml struct {
	Workspaces []struct {
		Name string
		Path string
	}
}

func (conf YpmConfigYml) Read() {

}
