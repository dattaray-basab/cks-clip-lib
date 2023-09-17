package common

import (
	"io/fs"
	"path/filepath"
)

var GetFirstFileInRootDir = func(filePath string) (string, error) {
	firstFile := ""
	err := filepath.WalkDir(filePath, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			firstFile = s
			return filepath.SkipDir
		}
		return nil
	})
	return firstFile, err
}
