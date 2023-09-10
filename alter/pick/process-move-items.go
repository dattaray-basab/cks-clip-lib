package pick

import (
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var GetMoveItemMap = func(templateMap map[string]string) map[string]globals.MoveItemDetailsT {
	moveItemMap := make(map[string]globals.MoveItemDetailsT)
	moveItems := templateMap[globals.KEY_MOVE_ITEMS]
	moveItemParts := strings.Split(moveItems, ":")
	index := 0
	for _, moveItemVal := range moveItemParts {
		moveItemKey := strings.Replace(moveItemVal, ".", "_", -1)
		MoveItemDetails := globals.MoveItemDetailsT{Key: moveItemVal, Index: index}

		moveItemMap[moveItemKey] = MoveItemDetails
		// moveItemMap[moveItemKey.Key] = moveItemVal

		// moveItemMap[moveItemKey].Index = index
		index++
	}
	// moveFile := templateMap[globals.KEY_ALTER_PATH]
	return moveItemMap
}
