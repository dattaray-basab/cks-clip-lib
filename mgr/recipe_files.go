package mgr

import (
	"path/filepath"

	"github.com/otiai10/copy"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func SetupRecipeFiles(absPathToSource string,  src_recipe_dirpath string) (string, bool, error) {
	dst_recipe_dirpath := filepath.Join(absPathToSource, globals.RECIPE_ROOT_DIR_)
	err := copy.Copy(src_recipe_dirpath, dst_recipe_dirpath)
	if err != nil {
		return "", true, err
	}
	return dst_recipe_dirpath, false, nil
}
