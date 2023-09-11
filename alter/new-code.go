package alter

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/alter/pick"
	"github.com/dattaray-basab/cks-clip-lib/alter/recast"
	"github.com/dattaray-basab/cks-clip-lib/alter/transform"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

var BuildNewAlterDir = func(templateMap map[string]string) error {

	templateMap[globals.KEY_STORE_DIR_PATH] = filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
	templateMap[globals.KEY_CONTROL_JSON_PATH] = filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.CONTROL_JSON_FILE)

	var buildStoreDir = func(templateMap map[string]string) error {

		var matches = func(list []string, item string) bool {
			for _, listItem := range list {
				if listItem == item {
					return true
				}
			}
			return false
		}

		move_items := strings.Split(templateMap[globals.KEY_MOVE_ITEMS], ":")
		for i, item := range move_items {
			move_items[i] = strings.TrimSpace(item)
		}

		logger.Log.Debug(move_items)

		files, err := os.ReadDir(templateMap[globals.KEY_CODE_BLOCK_PATH])
		if err != nil {
			return err
		}
		foundCount := 0
		for _, item := range files {
			if matches(move_items, item.Name()) {
				item_path := filepath.Join(templateMap[globals.KEY_CODE_BLOCK_PATH], item.Name())
				// if common.IsDir(item_path) {
				parentDir := filepath.Dir(item_path)
				alterName := filepath.Base(templateMap[globals.KEY_ALTER_PATH])
				newParent := filepath.Join(parentDir, alterName, globals.STORE_DIRNAME)

				err := os.MkdirAll(newParent, os.ModePerm)
				if err != nil {
					return err
				}
				new_dirpath := filepath.Join(newParent, item.Name())
				err = os.Rename(item_path, new_dirpath)
				if err != nil {
					return err
				}
				foundCount++
			}
		}
		if foundCount ==0 || len(move_items) != foundCount {
			msg := "*** FAILED ***: atleast one move item was not found"
			logger.Log.Error(msg)
			return errors.New(msg)
		}
		return nil
	}

	err := buildStoreDir(templateMap)
	if err != nil {
		return err
	}

	switch templateMap[globals.KEY_ALTER_SUB_COMMAND] {
	case globals.TRANSFORM:
		err = transform.BuildSubcommand(templateMap)
		if err != nil {
			return err
		}
	case globals.RECAST:
		err = recast.BuildSubcommand(templateMap)
		if err != nil {
			return err
		}
	case globals.PICK:
		err = pick.BuildSubcommand(templateMap)
		if err != nil {
			return err
		}
	default:
		msg := "*** FAILED ***: unknown alter sub-command"
		logger.Log.Error(msg)
		return errors.New(msg)
	}

	return nil
}
