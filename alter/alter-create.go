package alter

import (
	// "log"
	"path/filepath"
	// "strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var CreatePhaseFile = func(templateMap map[string]string) error {
	var buildNewPhaseFile = func(phasePath string, phaseName, lastPhase string) error {
		baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
		recipeScaffold := globals.ScaffoldInfoTListT{

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

		err := common.CreateFiles(recipeScaffold)
		return err
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
