package utils

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func WriteYaml(path string, data any) {
	bytes, err := yaml.Marshal(data)
	Check(err)

	file, err := os.Create(path)
	Check(err)

	defer file.Close()

	_, err = io.WriteString(file, string(bytes))
	Check(err)
}
