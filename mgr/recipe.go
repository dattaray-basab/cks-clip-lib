package mgr

import (
	"errors"
	"fmt"
	"log"
	"os"

	// "os/exec"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/otiai10/copy"
)

func CreateRecipe(src_recipe_dirpath string, templateMap map[string]string, absPathToSource string, absPathToRecipeParent string, overwrite bool) error {
	err := checkInputs(absPathToRecipeParent, absPathToSource, overwrite)
	if err != nil {
		return err
	}

	dst_recipe_dirpath := filepath.Join(absPathToSource, globals.RECIPE_ROOT_DIR_)
		err = copy.Copy(src_recipe_dirpath, dst_recipe_dirpath)
	if err != nil {
		return err
	}

	err = common.Refactor( dst_recipe_dirpath, templateMap, "{{code_block}}", "base", "*.json")
	if err != nil {
		println(err)
	}



	return nil
}

func CopyDir(recipe_dirpath1, recipe_dirpath2 string) {
	panic("unimplemented")
}

func checkInputs(absPathToSource string, absPathToRecipeParent string, overwrite bool) error {
	var success bool
	success = common.IsDir(absPathToSource)
	if !success {
		err := errors.New("mgr/recipe.go::checkInputs: " + "source folder does not exist: " + absPathToSource)
		log.Printf("%s", err)
		return err
	}
	success = common.IsDir(absPathToRecipeParent)
	if !success {
		err := os.Mkdir(absPathToRecipeParent, os.ModePerm)
		if err != nil {
			err = errors.New("mgr/recipe.go::checkInputs: " + "could not create recipe parent folder: " + absPathToRecipeParent)
			return err
		}
	}
	recipe_dirpath := filepath.Join(absPathToRecipeParent, globals.RECIPE_ROOT_DIR_)
	if common.IsDir(recipe_dirpath) {
		if overwrite {
			err := os.RemoveAll(recipe_dirpath)
			return err
		} else {
			err := errors.New("mgr/recipe.go::checkInputs: " + "recipe folder already exists: " + recipe_dirpath)
			log.Printf("%s", err)
			return err
		}
	}

	return nil

}

func CreatePathIfAbsent(recipePath string) error {
	if _, err := os.Stat(recipePath); errors.Is(err, os.ErrNotExist) {

		fmt.Println("recipe folder already exists: Overwrite by entering y/Y: ")

		// var then variable name then variable type
		var doOverride string

		// Taking input from user
		fmt.Scanln(&doOverride)

		if doOverride == "y" || doOverride == "Y" {
			fmt.Println("Overwriting recipe folder")

			err := os.Mkdir(recipePath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Exiting")
			return errors.New("Exiting")
		}
	} else {
		fmt.Println("recipe folder does not exist: Creating")
		err := os.Mkdir(recipePath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
