package filegen

import (
	"errors"
	"log"
	"os"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func CreateOrUpdatePhaseFile(phasePath, phaseName, lastPhase string) error {

	var checkDependsonPhaseFileName = func(phasePath string) (bool, error) {
		// var getPhaseName = func(jsonMap map[string]interface{}) (string, error) {
		// 	if phaseName, ok := jsonMap[globals.CODE_BLOCK_]; ok {
		// 		return phaseName.(string), nil
		// 	}
		// 	return "", nil
		// }

		files, err := os.ReadDir(phasePath)
		if err != nil {
			return false, err
		}
		for _, file := range files {
			if !file.IsDir() {
				// filePath := filepath.Join(phasePath, file.Name())
				// log.Println(filePath)
				// jsonMap, err := common.ReadJsonFile(filePath)
				// if err != nil {
				// 	return false, err
				// }
				// log.Println(jsonMap)
				// phaseNameFromFile, err := getPhaseName(jsonMap)
				// if err != nil {
				// 	return "", err
				// }

				phaseNameFromFile := file.Name()
				lastPhaseFileFromDirective := lastPhase + globals.JSON_EXT
				if phaseNameFromFile == lastPhaseFileFromDirective {
					return true, nil
				}

			}
		}
		return false, nil

	}

	log.Println(phasePath)
	success, err := checkDependsonPhaseFileName(phasePath)
	if err != nil {
		return err
	}
	if !success {		
		errNew := errors.New("The phase name " + lastPhase + " does not exist")
		return errNew
	}

	// does phase already exist?

	// if so read the file and add to the end ops_pipeline

	// if not create a new file and add under __PHASES

	return nil
}
