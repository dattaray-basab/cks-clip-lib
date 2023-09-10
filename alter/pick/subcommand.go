package pick

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
	"github.com/dattaray-basab/cks-clip-lib/template_tests/expt1"
	"github.com/dattaray-basab/cks-clip-lib/template_tests/expt2"
	"github.com/dattaray-basab/cks-clip-lib/template_tests/expt3"
	"github.com/dattaray-basab/cks-clip-lib/templates"
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

	queryFilePath, err := getQueryFilePath(templateMap)
	if err != nil {
		return err
	}
	moveItemMap := common.GetMoveItemMap(templateMap)
	logger.Log.Debug(moveItemMap)
	QuotedFullQueryId, err := getQueryId(templateMap, queryFilePath)
	if err != nil {
		return err
	}

	QuotedShortQueryId := globals.QUOTE + templateMap[globals.KEY_ALTER_NAME] + globals.QUOTE

	substitutionTemplate :=
		globals.SubstitionTemplateT{
			FullQueryId:   QuotedFullQueryId,
			ShortQueryId:  QuotedShortQueryId,
			MoveItemsInfo: moveItemMap,
		}

	expt1.Expt1(templateMap, substitutionTemplate)
	expt2.Expt2(templateMap, substitutionTemplate)
	expt3.Expt3(templates.T3, templateMap, substitutionTemplate)

	err = MakeControlFile(templateMap, moveItemMap, QuotedFullQueryId)
	if err != nil {
		return err
	}

	err = MakeQueryTokenFile(templateMap, moveItemMap, queryFilePath, QuotedFullQueryId)
	if err != nil {
		return err
	}
	return nil
}
