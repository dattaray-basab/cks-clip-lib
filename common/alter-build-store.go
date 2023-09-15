package common

import (
	"errors"

	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

var BuildStore = func(templateMap map[string]string) error {
	var buildStoreDir = func(templateMap map[string]string, writeMode bool) error {

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

		missingItems := []string{}
		missingMatches := []string{}
		for _, item := range files {
			if matches(move_items, item.Name()) {
				item_path := filepath.Join(templateMap[globals.KEY_CODE_BLOCK_PATH], item.Name())
				if _, err := os.Stat(item_path); os.IsNotExist(err) {
					missingItems = append(missingItems, item.Name())
				}
				parentDir := filepath.Dir(item_path)
				alterName := filepath.Base(templateMap[globals.KEY_ALTER_PATH])
				newParent := filepath.Join(parentDir, alterName, globals.STORE_DIRNAME)
				if writeMode {
					err := os.MkdirAll(newParent, os.ModePerm)
					if err != nil {
						return err
					}
					new_dirpath := filepath.Join(newParent, item.Name())
					err = os.Rename(item_path, new_dirpath)
					if err != nil {
						return err
					}
				}

			} else {
				missingMatches = append(missingMatches, item.Name())
			}
		}

		actualMisses := []string{}
		for _, moveItem := range move_items {
			found := false
			for _, item := range missingMatches {
				if item == moveItem {
					found = true
				}
			}
			if !found {
				actualMisses = append(actualMisses, moveItem)
			}
		}

		if len(missingItems) > 0 {
			msg := "FAILED: move items that do not exist: " + strings.Join(missingItems, ", ")
			logger.Log.Error(msg)
			return errors.New(msg)
		}

		if len(actualMisses) > 0 {
			msg := "FAILED: actual move items that do not match: " + strings.Join(actualMisses, ", ")
			logger.Log.Error(msg)

			return errors.New(msg)
		}

		return nil
	}

	err := buildStoreDir(templateMap, false)
	if err != nil {
		return err
	}
	err = buildStoreDir(templateMap, true)
	if err != nil {
		return err
	}
	return nil
}
