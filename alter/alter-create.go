package alter

import (
	"path/filepath"	

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var CreatePhaseFile = func(templateMap map[string]string) error {



	lastPhase := templateMap[globals.KEY_LAST_PHASE]
	phaseName := templateMap[globals.KEY_PHASE_NAME]

	phasesPath := templateMap[globals.KEY_PHASES_PATH]
	err := BuildNewPhaseFile(phasesPath, phaseName, lastPhase)
	if err != nil {
		return err
	}

	fullPhasePath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)

	// substitute the templateMap values
	err = common.Refactor(fullPhasePath, templateMap, "*.*")
	if err != nil {
		return err
	}

	err = BuildNewAlterDir(templateMap)
	if err != nil {
		return err
	}


	return nil
}
