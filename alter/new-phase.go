package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildNewPhaseFile = func(phasePath string, phaseName, lastPhase string) error {

	baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	scaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(baseDirpath),
			Content: `
{
  "__CODE_BLOCK": {{code-block-name-with-quotes}},
  "__DEPENDS_ON": [
	{{depends-on-phase-with-quotes}}
  ],
  "ops_pipeline": [
	{
	  "alter": {
		"locator": [
		  {{full-alter-path-with-quotes}}
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
