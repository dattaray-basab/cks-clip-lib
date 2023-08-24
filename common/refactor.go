package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)



func Refactor(srcDirPath string, templateMap map[string]string, old, new string, patterns ...string) error {
	return filepath.Walk(srcDirPath, refactorFunc(templateMap, old, new, patterns))
}

func refactorFunc(templateMap map[string]string, old, new string, filePatterns []string) filepath.WalkFunc {
	return filepath.WalkFunc(func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		var matched bool
		for _, pattern := range filePatterns {
			var err error
			matched, err = filepath.Match(pattern, fi.Name())
			if err != nil {
				return err
			}

			if matched {
				read, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				fmt.Println("Refactoring:", path)

				newContents := strings.Replace(string(read), old, new, -1)

				err = os.WriteFile(path, []byte(newContents), 0)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}