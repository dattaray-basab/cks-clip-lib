package common

import (
	"os"
	"path/filepath"
	"strings"
)

func Rename(dst_recipe_dirpath string, templateMap map[string]string) error {
	return filepath.Walk(dst_recipe_dirpath, renameFunc(templateMap))
}

func renameFunc(templateMap map[string]string) filepath.WalkFunc {
	var substitute = func(in_text string, templateMap map[string]string) string {
		out_text := in_text
		for old, new := range templateMap {
			out_text = strings.Replace(out_text, old, new, -1)
		}
		return out_text
	}

	return filepath.WalkFunc(func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			dirname := filepath.Base(path)
			if strings.HasPrefix(dirname, "{{") && strings.HasSuffix(dirname, "}}") {
				dirpath_substitute := substitute(path, templateMap)
				err := os.Rename(path, dirpath_substitute)
				if err != nil {
					return err
				}
				println(dirpath_substitute)

			}

			return nil
		}


		return nil
	})
}


