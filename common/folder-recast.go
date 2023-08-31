package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/logger"
)

func Refactor(dst_recipe_dirpath string, templateMap map[string]string, patterns ...string) error {
	return filepath.Walk(dst_recipe_dirpath, refactorFunc(templateMap,  patterns))
}

func refactorFunc(templateMap map[string]string, filePatterns []string) filepath.WalkFunc {
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

				path_substitute := substitute(path, templateMap)

				msg := fmt.Sprintf("Refactoring: %s ", path)
				logger.Log.Debug(msg)


				// newContents := strings.Replace(string(read), old, new, -1)
				newContents := substitute(string(read), templateMap)

				// Create parent directory for substituted path if it doesn't exist
				parentDir := filepath.Dir(path_substitute)
				if !IsDir(parentDir) {
					err = os.MkdirAll(parentDir, 0777)
					if err != nil {
						return err
					}
					err = os.RemoveAll(path)
					if err != nil {
						return err
					}
				}

				err = os.WriteFile(path_substitute, []byte(newContents), 0777)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func substitute(in_text string, templateMap map[string]string) string {
	out_text := in_text
	for old, new := range templateMap {
		out_text = strings.Replace(out_text, old, new, -1)
	}
	return out_text
}
