package filegen

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/alter"
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func CreateOrUpdatePhaseFile(templateMap map[string]string) error {
	lastPhase := templateMap[globals.KEY_LAST_PHASE]

	var checkDependsonPhaseFileName = func(phasePath string) (bool, error) {

		files, err := os.ReadDir(phasePath)
		if err != nil {
			return false, err
		}
		for _, file := range files {
			if !file.IsDir() {
				phaseNameFromFile := file.Name()
				lastPhaseFileFromDirective := lastPhase + globals.JSON_EXT
				if phaseNameFromFile == lastPhaseFileFromDirective {
					return true, nil
				}
			}
		}
		return false, nil
	}

	templateMap[globals.KEY_PHASES_PATH] = filepath.Join(templateMap[globals.KEY_BLUEPRINTS_PATH], templateMap[globals.KEY_TARGET], globals.PHASES_DIRNAME)
	phasesPath := templateMap[globals.KEY_PHASES_PATH]

	log.Println(phasesPath)
	success, err := checkDependsonPhaseFileName(phasesPath)
	if err != nil {
		return err
	}
	if !success {
		errNew := errors.New("The phase name " + lastPhase + " does not exist")
		return errNew
	}

	phaseName := templateMap[globals.KEY_PHASE_NAME]

	// does phase already exist?
	currentPhaseFilePath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)
	isFile := common.IsFile(currentPhaseFilePath)
	log.Println(isFile)
	if isFile {
		// if so update the file
		err = alter.UpdatePhaseFile(templateMap)
		if err != nil {
			return err
		}
	} else {
		// if not create a new file
		err = alter.CreatePhaseFile(templateMap)
		if err != nil {
			return err
		}
	}

	return nil
}
