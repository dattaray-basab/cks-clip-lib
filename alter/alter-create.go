package alter

import (
	// "log"

	"log"
	"os"
	"path/filepath"
	"strings"

	// "strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"

)

var CreatePhaseFile = func(templateMap map[string]string) error {
	var buildNewPhaseFile = func(phasePath string, phaseName, lastPhase string) error {

		baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
		scaffold := globals.ScaffoldInfoTListT{

			{
				Filepath: filepath.Join(baseDirpath),
				Content: `
{
  "__DEPENDS_ON": [
	{{depends-on-phase}}
  ],
  "ops_pipeline": [
	{
	  {{alter-name}}: {
		"locator": [
		  {{alter-path-with-quotes}}
		]
	  }
	}
  ]
}
		`,
			},
		}

		err := common.CreateFiles(scaffold)
		return err
	}

	var buildAlterDir = func(templateMap map[string]string) error {

		templateMap[globals.KEY_STORE_DIR_PATH] = filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
		templateMap[globals.KEY_CONTROL_JSON_PATH] = filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.CONTROL_JSON_FILE)

		var buildControlFile = func(templateMap map[string]string) error {
			scaffold := globals.ScaffoldInfoTListT{

				{
					Filepath: filepath.Join(templateMap[globals.KEY_CONTROL_JSON_PATH]),
					Content: `
	[
	{
		"op": "transform"
	}
	]
			`,
				},
			}

			err := common.CreateFiles(scaffold)
			if err != nil {
				return err
			}

			return nil
		}

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
			log.Println(move_items)

			files, err := os.ReadDir(templateMap[globals.KEY_CODE_BLOCK_PATH])
			if err != nil {
				return err
			}
			for _, item := range files {
				if matches(move_items, item.Name()) {
					item_path := filepath.Join(templateMap[globals.KEY_CODE_BLOCK_PATH], item.Name())
					if common.IsDir(item_path) {
						parentDir := filepath.Dir(item_path)
						newParent := filepath.Join(parentDir, globals.STORE_DIRNAME)		
						log.Println(newParent)
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
				}
			}
			return nil
		}

		err := buildStoreDir(templateMap)
		if err != nil {
			return err
		}

		err = buildControlFile(templateMap)
		if err != nil {
			return err
		}

		return nil
	}

	lastPhase := templateMap[globals.KEY_LAST_PHASE]
	phaseName := templateMap[globals.KEY_PHASE_NAME]

	phasesPath := templateMap[globals.KEY_PHASES_PATH]
	err := buildNewPhaseFile(phasesPath, phaseName, lastPhase)
	if err != nil {
		return err
	}

	fullPhasePath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)

	// substitute the templateMap values
	err = common.Refactor(fullPhasePath, templateMap, "*.*")
	if err != nil {
		return err
	}

	err = buildAlterDir(templateMap)
	if err != nil {
		return err
	}

	// relAlterPathFromPhase := templateMap["{{alter-dir-path}}"]
	// log.Println(relAlterPathFromPhase)
	// alterPathWithoutQuotes := strings.Trim(relAlterPathFromPhase, "\"")

	// relAlterPath := strings.TrimPrefix(alterPathWithoutQuotes, "/")
	// log.Println(relAlterPath)

	// files, err := os.ReadDir(fullPhasePath)
	// if err != nil {
	// 	return  err
	// }

	// now add to the CODE_BlOCK: storage (from item-list) + control files (use recast or transform as they will just copy the files from storage)

	return nil
}
