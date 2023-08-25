package mgr

import (
	"errors"

	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/otiai10/copy"
)

func CreateRecipe(templateMap map[string]string, targetDirpath string, recipeDirpath string, overwrite bool) error {
	var checkInputs = func(dst_recipe_dirpath string, absPathToRecipeParent string, overwrite bool) error {
		var success bool
		success = common.IsDir(dst_recipe_dirpath)
		if !success {
			err := errors.New("mgr/recipe.go::checkInputs: " + "recipe folder does not exist: " + dst_recipe_dirpath)
			if err != nil {
				return err
			}
		}
		success = common.IsDir(absPathToRecipeParent)
		if !success {
			err := os.Mkdir(absPathToRecipeParent, os.ModePerm)
			if err != nil {
				err = errors.New("mgr/recipe.go::checkInputs: " + "could not create recipe parent folder: " + absPathToRecipeParent)
				return err
			}
		}
		return nil
	}
	var processBlueprint = func(templateMap map[string]string, recipePath string, srcTargetPath string, overwrite bool) error {

		err := common.Refactor(recipePath, templateMap, "*.json")
		if err != nil {
			println(err)
		}
		shouldReturn, returnValue := common.CleanuupSubstitutedDirectories(recipePath)
		if shouldReturn {
			return returnValue
		}
		err = common.Rename(recipePath, templateMap)
		if err != nil {
			println(err)
		}
		return nil
	}
	var processBlockCode = func(templateMap map[string]string, recipePath string, srcTargetPath string) error {
		var err error
		code_block := templateMap["{{code_block}}"]
		target := templateMap["{{target}}"]

		target_code_path := filepath.Join(recipePath, target, globals.RECIPE_ROOT_DIR_, globals.CODE_BLOCK_ROOT, code_block)
		if common.IsDir(target_code_path) {
			err := os.RemoveAll(target_code_path)
			if err != nil {
				return err
			}

			err = os.MkdirAll(target_code_path, os.ModePerm)
			if err != nil {
				return err
			}
		}

		err = copy.Copy(srcTargetPath, target_code_path)
		if err != nil {
			return err
		}
		return nil
	}

	err := checkInputs(recipeDirpath, targetDirpath, overwrite)
	if err != nil {
		return err
	}

	// dst_recipe_dirpath := filepath.Join(absPathToSource, globals.RECIPE_ROOT_DIR_)
	err = processBlueprint(templateMap, recipeDirpath, targetDirpath, overwrite)
	if err != nil {
		return err
	}

	err = processBlockCode(templateMap, recipeDirpath, targetDirpath)
	return err
}
