package fs

import (
	"os"
	"path/filepath"
)

func ReadFolder(path string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func ReadFolderWithFullPath(path string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fullPath := filepath.Join(path, entry.Name())
			files = append(files, fullPath)
		}
	}

	return files, nil
}

func ReadFolderRecursive(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, filePath)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
