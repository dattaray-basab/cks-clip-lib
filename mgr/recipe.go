package mgr

import (
	"errors"

	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/filegen"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/otiai10/copy"
)

func CreateRecipe(templateMap map[string]string, targetDirpath string, recipeDirpath string, queryToken string, force bool) error {
	var checkInputs = func(dst_recipe_dirpath string, absPathToRecipeParent string, force bool) error {
		var success bool

		success = common.IsDir(dst_recipe_dirpath)
		if !success {
			err := os.MkdirAll(dst_recipe_dirpath, os.ModePerm)
			if err != nil {
				err = errors.New("mgr/recipe.go::checkInputs: " + "could not create recipe folder: " + dst_recipe_dirpath)
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
	var processBlueprint = func(templateMap map[string]string, recipePath string, srcTargetPath string, force bool) error {

		err := common.Refactor(recipePath, templateMap, "*.*")
		if err != nil {
			return err
		}
		err = common.CleanuupSubstitutedDirectories(recipePath)
		if err != nil {
			return err
		}

		err = common.Rename(recipePath, templateMap)
		if err != nil {
			return err
		}
		return nil
	}
	var processBlockCode = func(templateMap map[string]string, recipePath string, srcTargetPath string) error {
		var err error
		code_block := templateMap["{{code_block}}"]

		target_code_path := filepath.Join(recipePath, globals.CODE_BLOCK_ROOT, code_block)

		err = os.RemoveAll(target_code_path)
		if err != nil {
			return err
		}

		err = os.MkdirAll(target_code_path, os.ModePerm)
		if err != nil {
			return err
		}

		err = copy.Copy(srcTargetPath, target_code_path)
		if err != nil {
			return err
		}
		return nil
	}

	err := checkInputs(recipeDirpath, targetDirpath, force)
	if err != nil {
		return err
	}

	err = filegen.CreateRecipeFiles(recipeDirpath, queryToken)
	if err != nil {
		return err
	}

	// dst_recipe_dirpath := filepath.Join(absPathToSource, globals.RECIPE_ROOT_DIR_)
	err = processBlueprint(templateMap, recipeDirpath, targetDirpath, force)
	if err != nil {
		return err
	}

	err = processBlockCode(templateMap, recipeDirpath, targetDirpath)
	return err
}
