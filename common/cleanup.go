package common

import (
	"os"
	"path/filepath"
)


func CleanuupSubstitutedDirectories(dst_recipe_dirpath string) (bool, error) {
	dir_to_cleanup := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "{{target}}")
	err := os.RemoveAll(dir_to_cleanup)
	if err != nil {
		return true, err
	}
	return false, nil
}