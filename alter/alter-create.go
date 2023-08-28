package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var CreatePhaseFile = func(templateMap map[string]string, moveItems []string, phasePath string, phaseName, lastPhase string) error {
	var buildNewPhaseFile = func(phasePath string, phaseName, lastPhase string) error {
		baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
		recipeScaffold := globals.ScaffoldInfoTListT{

			{
				Filepath: filepath.Join(baseDirpath),
				Content: `
{
  "__DEPENDS_ON": [
	{{last-phase-name}}
  ],
  "ops_pipeline": [
	{
	  {{alter-name}}: {
		"locator": [
		  {{alter-dir-path}}
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

	err := buildNewPhaseFile(phasePath, phaseName, lastPhase)
	if err != nil {
		return err
	}

	fullPhasePath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)

	// substitute the templateMap values
	err = common.Refactor(fullPhasePath, templateMap, "*.*")
	if err != nil {
		return err
	}

	// now add to the CODE_BlOCK: storage (from item-list) + control files (use recast or transform as they will just copy the files from storage)

	return nil
}
