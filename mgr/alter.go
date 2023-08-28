package mgr

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/filegen"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

const prefix = "/"

var codeBlockPath string

func AddAlter(
	templateMap map[string]string,
) error {

	var calcAlterPath = func(templateMap map[string]string) (string, error) {
		var joinAlterDirPath = func(baseDir string, frags []string) string {
			for _, frag := range frags {
				baseDir = filepath.Join(baseDir, frag)
			}
			return baseDir
		}
		alterDirPath := templateMap[globals.KEY_ALTER_DIR_PATH]
		codeBlockName := templateMap[globals.KEY_CODE_BLOCK_NAME]
		recipeDirpath := templateMap[globals.KEY_RECIPE_PATH]
		alterName := templateMap[globals.KEY_ALTER_NAME]
		force := templateMap[globals.KEY_FORCE]

		if !strings.HasPrefix(alterDirPath, prefix) {
			err := fmt.Errorf("alter-dir-path %s must start with %s", alterDirPath, prefix)
			return "", err
		}
		cutAlterDirPath := strings.TrimPrefix(alterDirPath, prefix)
		cutAlterDirParts := strings.Split(cutAlterDirPath, prefix)
		codeBlockPath = filepath.Join(recipeDirpath, "__CODE", codeBlockName)
		codeBlockPath = joinAlterDirPath(codeBlockPath, cutAlterDirParts)
		prefixedAlterName := globals.SPECIAL_DIR_PREFIX_ + alterName
		fullAlterPath := filepath.Join(codeBlockPath, prefixedAlterName)
		if force == "F" {
			if common.IsDir(fullAlterPath) {
				err := fmt.Errorf("full-alter-path %s already exists", fullAlterPath)
				return fullAlterPath, err
			}
		} else {
			err := os.RemoveAll(fullAlterPath)
			if err != nil {
				return fullAlterPath, err
			}
		}
		err := os.MkdirAll(fullAlterPath, os.ModePerm)
		if err != nil {
			err := fmt.Errorf("could not create full-alter-path %s", fullAlterPath)
			return fullAlterPath, err
		}

		return fullAlterPath, nil
	}

	var createPhase = func() error {
		var getTargetNameFromBlueprint = func(blueprintPath string) (string, error) {
			files, err := os.ReadDir(blueprintPath)
			if err != nil {
				return "", err
			}

			targetFromBlueprint := ""
			for _, file := range files {
				if file.IsDir() {
					if file.Name() != globals.TOKENS_DIRNAME {
						targetFromBlueprint = file.Name()
						return targetFromBlueprint, nil
					}
				}
			}
			return "", nil
		}

		recipeDirpath := templateMap[globals.KEY_RECIPE_PATH]
		target := templateMap[globals.KEY_TARGET]


		blueprintPath := filepath.Join(
			recipeDirpath,
			globals.BLUEPRINTS_DIRNAME)

		targetFromBlueprint, err := getTargetNameFromBlueprint(blueprintPath)
		if err != nil {
			return err
		}

		targetName := targetFromBlueprint
		if len(target) > 0 {
			targetName = target
		}

		phasePath := filepath.Join(blueprintPath, targetName, globals.PHASES_DIRNAME)

		log.Println(phasePath)
		err = filegen.CreateOrUpdatePhaseFile(templateMap)
		if err != nil {
			return err
		}
		return nil
	}

	alterPath, err := calcAlterPath(templateMap)
	if err != nil {
		return err
	}
	log.Println(alterPath)

	err = createPhase()
	if err != nil {
		return err
	}

	// create new phase

	return nil
}
