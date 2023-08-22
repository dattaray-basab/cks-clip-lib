package mgr

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func CreateRecipe(absPathToSource string, absPathToRecipeParent string, overwrite bool) error {
	err := checkInputs(absPathToRecipeParent, absPathToSource)
	if err != nil {
		return err
	}
	// recipe_dirpath := filepath.Join(absPathToRecipeParent, globals.RECIPE_ROOT_DIR_)
	// cwd, _ := os.Getwd()
	// fmt.Println("cwd: ", cwd)
	cwd, _ := os.Getwd()
	recipe_dirpath := filepath.Join(cwd, globals.RECIPE_ROOT_DIR_)

	err = common.CopyDir(recipe_dirpath, recipe_dirpath)
	if err != nil {
		return err
	}

	return nil
}

func CopyDir(recipe_dirpath1, recipe_dirpath2 string) {
	panic("unimplemented")
}

func checkInputs(absPathToSource string, absPathToRecipeParent string) error {
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
		err := errors.New("mgr/recipe.go::checkInputs: " + "recipe folder already exists: " + recipe_dirpath)
		log.Printf("%s", err)
		return err
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
