package mgr

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
)

func CreateRecipe(absPathToSource string, absPathToRecipeParent string, overwrite bool) string {
	pathToRecipe := prolog(absPathToRecipeParent, absPathToSource)
	// err := CreatePathIfAbsent(pathToRecipe)
	// if err != nil {
	// 	return err
	// }
	return pathToRecipe
}

func prolog(absPathToRecipeParent string, absPathToSource string) string {
	fmt.Println()
	fmt.Println("<<< CreateRecipe/n")
	fmt.Println("absPathToRecipeParent::", absPathToRecipeParent)
	fmt.Println("absPathToSource::", absPathToSource)
	absPathToRecipe := filepath.Join(absPathToRecipeParent, "__recipe")
	pathToRecipe := common.GetRecipePath(absPathToRecipe, true)
	fmt.Println("pathToRecipe::", pathToRecipe)
	fmt.Println(">>> CreateRecipe")
	fmt.Println()
	return pathToRecipe
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
