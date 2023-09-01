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
