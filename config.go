package config

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func FindFileInFileSystem(fileSystem fs.FS, basePath string, currentDir string, filename string) (string, error) {
	var base string
	if basePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting home dir: %w", err)
		}
		base = homeDir
	} else {
		base = basePath
	}
	for {
		rel, _ := filepath.Rel(base, currentDir)
		if rel == "." {
			break
		}
		filePath := filepath.Join(currentDir, filename)
		fmt.Fprintf(os.Stderr, "checking %s\n", filePath)
		if _, err := fileSystem.Open(filePath); err == nil {
			return filePath, nil
		}

		currentDir += "/.."
	}

	return "", fmt.Errorf("file %s not found", filename)
}
