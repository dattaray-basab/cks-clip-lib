package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildNewPhaseFile = func(templateMap map[string]string) (string, error)  {
	// lastPhase := templateMap[globals.KEY_LAST_PHASE]
	phaseName := templateMap[globals.KEY_PHASE_NAME]
	phasesPath := templateMap[globals.KEY_PHASES_PATH]

	phaseFilepath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)
	scaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(phaseFilepath),
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
		  "{{full-alter-rel-path}}"
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
		return phaseFilepath, err
	}

	err = common.ReplaceUsingTemplateMap(templateMap, phasesPath)
	if err != nil {
		return phaseFilepath, err
	}
	return phaseFilepath, err
}
