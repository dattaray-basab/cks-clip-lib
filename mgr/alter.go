package mgr

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

const prefix = "/"

var codeBlockPath string

func AddAlter(
	target string,
	recipeDirpath,
	alterDirPath,
	alterName string,
	itemListToAlter []string,
	phaseName string,
	lastPhase string,
	codeBlockName string,
	force bool,
) error {

	var calcAlterPath = func() (string, error) {
		var joinAlterDirPath = func(baseDir string, frags []string) string {
			for _, frag := range frags {
				baseDir = filepath.Join(baseDir, frag)
			}
			return baseDir
		}

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
		if !force {
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
		files, err := os.ReadDir(phasePath)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !file.IsDir() {
				filePath := filepath.Join(phasePath, file.Name())
				log.Println(filePath)
			}
		}
		return nil
	}

	alterPath, err := calcAlterPath()
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
