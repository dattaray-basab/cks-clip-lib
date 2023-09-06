package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildNewPhaseFile = func(templateMap map[string]string, phasePath string, phaseName, lastPhase string) error {

	baseDirpath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	scaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(baseDirpath),
			Content: `
{
  "__CODE_BLOCK": "{{code-block-name}}",
  "__DEPENDS_ON": [
	"{{depends-on-phase}}"
  ],
  "ops_pipeline": [
	{
	  "alter": {
		"locator": [
		  "{{full-alter-path}}"
		]
	  }
	}
  ]
}
		`,
		},
	}

	err := common.CreateFiles(scaffold)
	if err != nil {
		return err
	}

	err = common.ReplaceUsingTemplateMap(templateMap, phasePath)
	if err != nil {
		return err
	}
	return err
}
