package mgr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/filegen"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
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

		// if !strings.HasPrefix(alterDirPath, prefix) {
		// 	err := fmt.Errorf("alter-dir-path %s must start with %s", alterDirPath, prefix)
		// 	return "", err
		// }
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

		blueprintsPath := filepath.Join(
			recipeDirpath,
			globals.BLUEPRINTS_DIRNAME)
		templateMap[globals.KEY_BLUEPRINTS_PATH] = blueprintsPath

		templateMap[globals.KEY_CODE_BLOCK_ROOT_PATH] = filepath.Join(recipeDirpath, globals.CODE_BLOCK_ROOT)
		templateMap[globals.KEY_CODE_BLOCK_PATH] = codeBlockPath

		targetFromBlueprint, err := getTargetNameFromBlueprint(blueprintsPath)
		if err != nil {
			return err
		}

		targetName := targetFromBlueprint
		if len(target) > 0 {
			targetName = target
		}

		phasePath := filepath.Join(blueprintsPath, targetName, globals.PHASES_DIRNAME)

		logger.Log.Debug(phasePath)
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
	templateMap[globals.KEY_ALTER_PATH] = alterPath

	const QUOTE = "\""
	// KEY_CODE_BLOCK_NAME_WITH_QUOTES
	codeBlockNameWithQuotes := QUOTE + templateMap[globals.KEY_CODE_BLOCK_NAME] + QUOTE
	templateMap[globals.KEY_CODE_BLOCK_NAME_WITH_QUOTES] = codeBlockNameWithQuotes

	alterPathWithQuotes := QUOTE + templateMap[globals.KEY_ALTER_DIR_PATH] + QUOTE
	templateMap[globals.KEY_ALTER_PATH_WITH_QUOTES] = alterPathWithQuotes

	dependsOnPathWithQuotes := QUOTE + templateMap[globals.KEY_DEPENDS_ON_PHASE] + QUOTE
	templateMap[globals.KEY_DEPENDS_ON_PHASE_WITH_QUOTES] = dependsOnPathWithQuotes

	fullAlterPath := QUOTE + filepath.Join(templateMap[globals.KEY_ALTER_DIR_PATH], globals.SPECIAL_DIR_PREFIX_ + templateMap[globals.KEY_ALTER_NAME]) + QUOTE
	templateMap[globals.KEY_FULL_ALTER_PATH_WITH_QUOTES] = fullAlterPath

	err = createPhase()
	if err != nil {
		return err
	}

	// create new phase

	return nil
}
