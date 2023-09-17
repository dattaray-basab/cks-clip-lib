package transform

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

var BuildSubcommand = func(templateMap map[string]string) error {

	var getFirstLineOfFirstFile = func(templateMap map[string]string) (string, error) {
		var getFirstFile = func(filePath string) (string, error) {
			firstFile := ""
			err := filepath.WalkDir(filePath, func(s string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					firstFile = s
					return filepath.SkipDir
				}
				return nil
			})
			return firstFile, err
		}

		var readFirstLine = func(filePath string) (string, error) {
			file, err := os.Open(filePath)
			if err != nil {
				return "", err
			}
			defer file.Close()
			rawBytes, err := ioutil.ReadAll(file)
			if err != nil {
				return "", err
			}
			lines := strings.Split(string(rawBytes), "\n")
			firstLine := lines[0]
			logger.Log.Debug(firstLine)
			return firstLine, nil
		}


		storePath := filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
		if !common.IsDir(storePath) {
			err := fmt.Errorf("store-path %s does not exist", storePath)
			return "", err
		}

		firstFile, err := getFirstFile(storePath)
		if err != nil {
			return "", err
		}
		logger.Log.Debug(firstFile)

		firstLine, err := readFirstLine(firstFile)
		if err != nil {
			return "", err
		}

		return firstLine, nil
	}
	firstline, err := getFirstLineOfFirstFile(templateMap)
	if err != nil {
		return err
	}
	firstWord := strings.Split(firstline, " ")[0]
	fmt.Println(firstWord)

	err = common.BuildAlterInfrastucture(templateMap, TransformQueryTemplate, TransformControlTemplate)
	return err
}
