package mgr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
)

const prefix = "/"

var codeBlockPath string

func AddAlter(
	recipeDirpath,
	alterDirPath,
	alterName string,
	itemListToAlter []string,
	phaseName string,
	lastPhase string,
	codeBlockName string,
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
		fullAlterPath := filepath.Join(codeBlockPath, alterName)
		if common.IsDir(fullAlterPath) {
			err := fmt.Errorf("full-alter-path %s already exists", fullAlterPath)
			return fullAlterPath, err
		}
		err := os.MkdirAll(fullAlterPath, os.ModePerm)
		if err != nil {
			err := fmt.Errorf("could not create full-alter-path %s", fullAlterPath)
			return fullAlterPath, err
		}

		return fullAlterPath, nil
	}


	alterPath, err := calcAlterPath()
	if err != nil {
		return err
	}

	// create new phase
	
	println(alterPath)

	return nil
}
