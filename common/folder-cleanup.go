package common

import (
	"os"
	"path/filepath"
	"strings"
)

func Cleanup(dst_recipe_dirpath string, templateMap map[string]string) error {
	return filepath.Walk(dst_recipe_dirpath, cleanupFunc(templateMap))
}

func cleanupFunc(templateMap map[string]string) filepath.WalkFunc {

	return filepath.WalkFunc(func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			dirname := filepath.Base(path)
			if strings.HasPrefix(dirname, "{{") && strings.HasSuffix(dirname, "}}") {
				err := os.RemoveAll(path)
				if err != nil {
					return err
				}

			}

			return nil
		}


		return nil
	})
}


