package common

import (
	"os"
	"path/filepath"
	"strings"
)

func SubstitutePathsFromTemplate(templateMap map[string]string, dirpath string) error {
	const CONTENT_START = "{{"
	const CONTENT_END = "}}"

	var substituteDir = func(templateMap map[string]string) filepath.WalkFunc {
		var substitute = func(in_text string, templateMap map[string]string) string {
			out_text := in_text
			for old, new := range templateMap {
				out_text = strings.Replace(out_text, old, new, -1)
			}
			return out_text
		}

		return filepath.WalkFunc(func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if fileInfo.IsDir() {
				dirname := filepath.Base(path)
				if strings.HasPrefix(dirname, CONTENT_START) && strings.HasSuffix(dirname, CONTENT_END) {
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

	var cleanupDir = func(templateMap map[string]string) filepath.WalkFunc {

		return filepath.WalkFunc(func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if fileInfo.IsDir() {
				dirname := filepath.Base(path)
				if strings.HasPrefix(dirname, CONTENT_START) && strings.HasSuffix(dirname, CONTENT_END) {

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

	err := filepath.Walk(dirpath, substituteDir(templateMap))
	if err != nil {
		return err
	}

	err = filepath.Walk(dirpath, cleanupDir(templateMap))

	return err
}
