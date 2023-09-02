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
	Command LocationT `json:"alter"`
}
type LocationT struct {
	Location []string `json:"location"`
}

var UpdatePhaseFile = func(templateMap map[string]string) error {

	var getAlterJson = func(item string) CommandT {

		command := CommandT{
			Command: LocationT{
				Location: []string{item},
			},
		}
		// js, _ := json.Marshal(command)
		// fmt.Printf("%s" js)
		return command
	}

	// lastPhase := templateMap[globals.KEY_LAST_PHASE]
	phaseName := templateMap[globals.KEY_PHASE_NAME]
	phasesPath := templateMap[globals.KEY_PHASES_PATH]
	fullPhasePath := filepath.Join(phasesPath, phaseName+globals.JSON_EXT)

	fullAlterPathWithQuotes := fmt.Sprintf("\"%s\"", templateMap[globals.KEY_FULL_ALTER_PATH_WITH_QUOTES])
	logger.Log.Debug("fullAlterPathWithQuotes: ", fullAlterPathWithQuotes)

	command := getAlterJson(templateMap[globals.KEY_FULL_ALTER_PATH_WITH_QUOTES])
	logger.Log.Debug("command: ", command)

	jsonifiedCommand, _ := json.Marshal(command)
	logger.Log.Debug("json-command: ", string(jsonifiedCommand))

	if !common.IsFile(fullPhasePath) {
		err := errors.New("phase file does not exist")
		return err
	}
	// read the phase file
	result, err := common.ReadJsonFile(fullPhasePath)

	if err != nil {
		msg := fmt.Sprintf("phase file could not be read: %s", err.Error())
		err = errors.New(msg)
		return err
	}

	ops_pipeline := result["ops_pipeline"].([]interface{})
	pipelineLen := len(ops_pipeline)
	logger.Log.Debug("pipelineLen: ", pipelineLen)
	// pipeline_entry :=

	// 	logger.Log.Debug(ops)

	// ensure that the alter location is not already present
	// add the alter location to the phase file
	// write the modified phase file
	// generate the alter directory at the right location

	return nil
}
