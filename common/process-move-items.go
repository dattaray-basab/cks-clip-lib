package common

import (
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var GetMoveItemMap = func(templateMap map[string]string) map[string]globals.MoveItemDetailsT {
	moveItemMap := make(map[string]globals.MoveItemDetailsT)
	moveItems := templateMap[globals.KEY_MOVE_ITEMS]
	moveItemParts := strings.Split(moveItems, ":")
	lastIndexOfMoveItems := len(moveItemParts) -1
	index := 0
	for _, moveItemVal := range moveItemParts {
		isLastItem := false
		if index == lastIndexOfMoveItems {
			isLastItem = true
		} 
	
		moveItemKey := strings.Replace(moveItemVal, ".", "_", -1)
		MoveItemDetails := globals.MoveItemDetailsT{Key: moveItemKey, Index: index, IsLastItem: isLastItem}

		moveItemMap[moveItemKey] = MoveItemDetails
		index++
	}
	return moveItemMap
}
