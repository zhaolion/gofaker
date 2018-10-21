package gofaker

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// return current file path
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// return data path that contain random data files
func dataPath() string {
	return filepath.Join(getCurrentPath(), "data")
}

// return data file that contain random data
func localeFilePath(locale string) string {
	return filepath.Join(dataPath(), locale+".yml")
}

// return data source file name/path map
// eg.
// en: ./data/en.yml
func dataFiles() map[string]string {
	folder := dataPath()
	files := make(map[string]string)
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yml") {
			files[strings.Replace(info.Name(), ".yml", "", 1)] = path
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func containString(list []string, target string) bool {
	for _, l := range list {
		if l == target {
			return true
		}
	}

	return false
}
