package alter

import (
	// "log"
	"log"
	"path/filepath"

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

		log.Println("templateMap[globals.KEY_STORE_DIR_PATH]: ", templateMap[globals.KEY_STORE_DIR_PATH])
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
