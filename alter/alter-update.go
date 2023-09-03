package alter

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

type CommandT struct {
	Command LocatorT `json:"alter"`
}
type LocatorT struct {
	Locator []string `json:"locator"`
}

// ensure that the alter location is not already present
// add the alter location to the phase file
// write the modified phase file
// generate the alter directory at the right location
var UpdatePhaseFile = func(templateMap map[string]string) error {

	var getPhaseData = func(templateMap map[string]string) (map[string]interface{}, error) {
		phaseName := templateMap[globals.KEY_PHASE_NAME]
		phasesPath := templateMap[globals.KEY_PHASES_PATH]
		fullPhasePath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)

		if !common.IsFile(fullPhasePath) {
			err := errors.New("phase file does not exist")
			return nil, err
		}

		if !common.IsFile(fullPhasePath) {
			err := errors.New("phase file does not exist")
			return nil, err
		}
		// read the phase file
		result, err := common.ReadJsonFile(fullPhasePath)
		if err != nil {
			msg := fmt.Sprintf("phase file could not be read: %s", err.Error())
			err = errors.New(msg)
			return nil, err
		}
		return result, err
	}

	var getAlterJson = func(templateMap map[string]string) map[string]interface{} {

		jsonStr := `	{
	  "alter": {
		"locator": [
		  "/components/__pick"
		]
	  }
	}`

		var jsonFrame map[string]interface{}

		json.Unmarshal([]byte(jsonStr), &jsonFrame)
		return jsonFrame
	}

	var checkForDuplicateAlter = func(opsPipeline []interface{}, alterCommand interface{}) error {
		// check if the alter command is already present
		for _, alterCommandInPipeline := range opsPipeline {
			alterCommandInPipelineStr := fmt.Sprintf("%v", alterCommandInPipeline)
			alterCommandStr := fmt.Sprintf("%v", alterCommand)
			if alterCommandInPipelineStr == alterCommandStr {
				return errors.New("alter command already present")
			}
		}
		return nil

	}

	phaseContent, err := getPhaseData(templateMap)
	if err != nil {
		return err
	}

	alterCommand := getAlterJson(templateMap)
	logger.Log.Debug("alter command to add: ", alterCommand)

	opsPipeline := phaseContent["ops_pipeline"].([]interface{})
	err = checkForDuplicateAlter(opsPipeline, alterCommand)
	if err != nil {
		return err
	}
	
	// opsPipeline = opsPipeline.( alterCommand)
	phaseContent["ops_pipeline"] = opsPipeline

	jsonifiedPhase, _ := json.MarshalIndent(phaseContent, "", "  ")
	logger.Log.Debug(string(jsonifiedPhase))

	fmt.Println(string(jsonifiedPhase))

	if err != nil {
		msg := fmt.Sprintf("phase file could not be read: %s", err.Error())
		err = errors.New(msg)
		return err
	}

	ops_pipeline := phaseContent["ops_pipeline"].([]interface{})
	pipelineLen := len(ops_pipeline)
	logger.Log.Debug("pipelineLen: ", pipelineLen)

	return nil
}
