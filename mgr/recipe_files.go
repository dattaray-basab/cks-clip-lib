package mgr

import (

	"github.com/otiai10/copy"

)

func SetupRecipeFiles(templateMap map[string]string, dst_recipe_dirpath string,  src_recipe_dirpath string) ( error) {

	err := copy.Copy(src_recipe_dirpath, dst_recipe_dirpath)
	if err != nil {
		return  err
	}
	return  nil
}
