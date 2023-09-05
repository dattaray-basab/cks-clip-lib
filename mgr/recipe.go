package mgr

import (
	"errors"
	"fmt"
	"strconv"

	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/filegen"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
	"github.com/otiai10/copy"
)

func CreateRecipe(templateMap map[string]string, srcAppPath string, recipePath string, tokenFileName string) error {

	forceAsString := templateMap[globals.KEY_FORCE]
	force, err := strconv.ParseBool(forceAsString)
	msg := fmt.Sprintf("force: %v", force)
	logger.Log.Debug(msg)
	if err != nil {
		force = false
	}

	var checkInputs = func(dst_recipe_dirpath string, absPathToRecipeParent string) error {
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
	var processBlueprint = func(templateMap map[string]string, recipePath string, srcTargetPath string) error {

		err := filegen.CreateRecipeFiles(recipePath, tokenFileName)
		if err != nil {
			return err
		}

		err = common.SubstituteContentsFromTemplate(templateMap, recipePath)
		if err != nil {
			return err
		}
		err = common.Cleanup(recipePath)
		if err != nil {
			return err
		}

		err = common.SubstitutePathsFromTemplate(recipePath, templateMap)
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

	err = checkInputs(recipePath, srcAppPath)
	if err != nil {
		return err
	}

	// dst_recipe_dirpath := filepath.Join(absPathToSource, globals.RECIPE_ROOT_DIR_)
	err = processBlueprint(templateMap, recipePath, srcAppPath)
	if err != nil {
		return err
	}

	err = processBlockCode(templateMap, recipePath, srcAppPath)

	logger.Log.Info("SUCCESS: create recipe")
	return err
}
