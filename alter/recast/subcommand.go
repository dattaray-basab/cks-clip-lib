package recast

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	storePath := filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
	if !common.IsDir(storePath) {
		err := fmt.Errorf("store-path %s does not exist", storePath)
		return err
	}
	templateMap[globals.KEY_STORE_PATH] = storePath
	moveItems := templateMap[globals.KEY_MOVE_ITEMS]
	moveItemParts := strings.Split(moveItems, ":")
	if len(moveItemParts) == 0 {
		err := fmt.Errorf("no move-item is available")
		return err
	}
	firstMoveItem := moveItemParts[0]
	fmt.Println(firstMoveItem)

	changeMap := make(map[string]string)
	changeMap["{{name}}"] = firstMoveItem

	err := common.ReplaceUsingTemplateMap(changeMap, storePath)
	if err != nil {
		return err
	}

	err = common.BuildAlterInfrastucture(templateMap, QueryTemplate, ControlTemplate)
	return err
}
