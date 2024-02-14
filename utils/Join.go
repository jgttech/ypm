package utils

import "path"

func Join(paths ...string) string {
	cwd := []string{Cwd()}
	args := append(cwd, paths...)

	return path.Join(args...)
}
