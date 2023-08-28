package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var CreatePhaseFile = func(phasePath string, phaseName, lastPhase string) error {
	baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	recipeScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(baseDirpath),
			Content: `
{
  "__DEPENDS_ON": [
	{{last-phase}}
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
