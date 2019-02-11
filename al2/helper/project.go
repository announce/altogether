package helper

import (
	"errors"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

const ProjectPath = "src/github.com/announce/altogether"

func MustProjectPath(subdir string) string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	path, err := existingPath(strings.Split(goPath, ":"), subdir)
	if err != nil {
		panic(err)
	}
	return path
}

func existingPath(paths []string, subdir string) (string, error) {
	for _, path := range paths {
		fullPath := filepath.Join(path, ProjectPath, subdir)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", errors.New("project path was not found")
}
