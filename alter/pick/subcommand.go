package pick

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

var BuildSubcommand = func(templateMap map[string]string) error {


	var getQueryFilePath = func(templateMap map[string]string) (string, error) {
		dirpath := filepath.Join(templateMap[globals.KEY_BLUEPRINTS_PATH], globals.TOKENS_DIRNAME, globals.QUERY_DIRNAME)
		if !common.IsDir(dirpath) {
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
		fName := templateMap[globals.KEY_ALTER_NAME] + globals.JSON_EXT
		fPath := filepath.Join(dirpath, fName)
		filePathExists := common.IsFile(fPath)
		count := 0
		for filePathExists {
			fName = fName + strconv.Itoa(count) + globals.JSON_EXT
			fPath = filepath.Join(dirpath, fName)
			filePathExists = common.IsFile(fPath)
			count++
		}
		return fPath, nil
	}

	var getQueryId = func(templateMap map[string]string, queryFilePath string) (string, error) {
		queryFileName := filepath.Base(queryFilePath)
		queryName := queryFileName[:len(queryFileName)-len(globals.JSON_EXT)]
		logger.Log.Debug(queryName)
		suffix := 0
		queryId := "ID_" + strconv.Itoa(suffix)
		fullQueryId := globals.QUOTE + queryName + "." + queryId + globals.QUOTE

		return fullQueryId, nil
	}

	var getMoveMap = func(templateMap map[string]string) map[string]string {
		moveMap := make(map[string]string)
		moveItems := templateMap[globals.KEY_MOVE_ITEMS]
		moveItemParts := strings.Split(moveItems, ":")
		for _, moveItemVal := range moveItemParts {
			moveItemKey := strings.Replace(moveItemVal, ".", "_", -1)
			moveMap[moveItemKey] = moveItemVal
		}
		// moveFile := templateMap[globals.KEY_ALTER_PATH] 
		return moveMap
	}

	queryFilePath, err := getQueryFilePath(templateMap)
	if err != nil {
		return err
	}
	moveMap := getMoveMap(templateMap)
	logger.Log.Debug(moveMap)
	
	common.Junk(templateMap, moveMap)

	fullQueryId, err := getQueryId(templateMap, queryFilePath)
	if err != nil {
		return err
	}


	err = MakeControlFile(templateMap, moveMap, fullQueryId)
	if err != nil {
		return err
	}

	err = MakeQueryTokenFile(templateMap, queryFilePath, fullQueryId)
	if err != nil {
		return err
	}
	return nil
}
