package recast

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	var replaceFirstMoveItemFileNameWithToken = func(templateMap map[string]string) error {
		firstMoveItem, err  := common.GetFirstMoveItem(templateMap)
		if err != nil {
			return err
		}

		storePath := filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
		if !common.IsDir(storePath) {
			err := fmt.Errorf("store-path %s does not exist", storePath)
			return err
		}
		templateMap[globals.KEY_STORE_PATH] = storePath


		oldPath := filepath.Join(storePath, firstMoveItem)
		newPath := filepath.Join(storePath, "{{name}}")

		err = os.Rename(oldPath, newPath)
		if err != nil {
			return err
		}
		return nil
	}
	err := replaceFirstMoveItemFileNameWithToken(templateMap)
	if err != nil {
		return err
	}

	_, err = common.BuildAlterInfrastucture(templateMap, QueryTemplate, ControlTemplate)
	return err
}
