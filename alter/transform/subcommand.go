package transform

import (
	"fmt"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

var BuildSubcommand = func(templateMap map[string]string) error {

	var getFirstLineOfFirstFile = func(templateMap map[string]string) (string, error) {
		storePath := filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
		if !common.IsDir(storePath) {
			err := fmt.Errorf("store-path %s does not exist", storePath)
			return "", err
		}

		firstFile, err := common.GetFirstFileInRootDir(storePath)
		if err != nil {
			return "", err
		}

		wordList, err := common.GetWordsFromFile(firstFile)
		if err != nil {
			return "", err
		}
		firstWordInFirstFile := wordList[0]

		return firstWordInFirstFile, nil
	}
	firstWordInFirstFile, err := getFirstLineOfFirstFile(templateMap)
	if err != nil {
		return err
	}
	logger.Log.Debug(firstWordInFirstFile)
	templateMap[globals.KEY_FIRST_WORD_IN_FIRST_FILE] = firstWordInFirstFile

	err = common.BuildAlterInfrastucture(templateMap, TransformQueryTemplate, TransformControlTemplate)
	return err
}
