package helper

import (
	"errors"
	"github.com/announce/altogether/al2/domain"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const ProjectPath = "src/github.com/announce/altogether"

type DataDir map[domain.Type]string

var dataDir = DataDir{
	domain.Alfred: "testdata/Alfred.alfredpreferences",
	domain.Albert: "testdata/albert",
}
var tmpDir = ""

func EnsureDataPath(p domain.Type, config domain.TypedConfig) string {
	return ensureTmpDataDirs(dataDir[p], config.Path(p))
}

func MustRemoveTmpDir() {
	if tmpDir == "" {
		return
	}
	if err := os.RemoveAll(tmpDir); err != nil {
		panic(err)
	}
	tmpDir = ""
}
func mustTmpDir() string {
	if tmpDir == "" {
		name, err := ioutil.TempDir("", "")
		if err != nil {
			panic(err)
		}
		tmpDir = name
	}
	return tmpDir
}

func ensureTmpDataDirs(srcDir string, srcPath string) string {
	dir := filepath.Join(mustTmpDir(), srcDir, filepath.Dir(srcPath))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(filepath.Join(EnsureProjectPath(srcDir), srcPath))
	if err != nil {
		panic(err)
	}
	file, err := os.Create(filepath.Join(dir, filepath.Base(srcPath)))
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(data); err != nil {
		panic(err)
	}
	return filepath.Join(mustTmpDir(), srcDir)
}

func EnsureProjectPath(subdir string) string {
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
