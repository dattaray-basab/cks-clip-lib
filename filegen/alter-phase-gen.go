package filegen

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
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
	var createPhaseFile = func(phasePath string, phaseName, lastPhase string) error {
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

	var updatePhaseFile = func(phasePath string, phaseName, lastPhase string) error {
		phaseFilePath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
		log.Println(phaseFilePath)
		_, err := os.Create(phaseFilePath)
		if err != nil {
			return err
		}
		return nil
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
	currentPhaseFilePath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	isFile := common.IsFile(currentPhaseFilePath)
	log.Println(isFile)
	if isFile {
		// if so update the file
		err = updatePhaseFile(phasePath, phaseName, lastPhase)
		if err != nil {
			return err
		}
	} else {
		// if not create a new file
		err = createPhaseFile(phasePath, phaseName, lastPhase)
		if err != nil {
			return err
		}
	}

	return nil
}
