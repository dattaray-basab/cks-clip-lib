package mgr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
)

const prefix = "/"

var alterDirPathFull string
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
	var joinAlterDirPath = func(baseDir string, frags []string) string {
		for _, frag := range frags {
			baseDir = filepath.Join(baseDir, frag)
		}
		return baseDir
	}

	var checkParams = func(alterDirPath string) error {
		if !strings.HasPrefix(alterDirPath, prefix) {
			err := fmt.Errorf("alter-dir-path %s must start with %s", alterDirPath, prefix)
			return err
		}
		cutAlterDirPath := strings.TrimPrefix(alterDirPath, prefix)
		cutAlterDirParts := strings.Split(cutAlterDirPath, prefix)	
		codeBlockPath1 := filepath.Join(recipeDirpath, "__CODE", codeBlockName)
		codeBlockPath2 := joinAlterDirPath(codeBlockPath1, cutAlterDirParts)
		codeBlockPath = filepath.Join(codeBlockPath2, alterName)
		alterDirPathFull = filepath.Join(recipeDirpath, cutAlterDirPath)
		if common.IsDir(alterDirPathFull) {
			err := fmt.Errorf("alter-dir-path %s already exists", alterDirPathFull)
			return err
		}
		err := os.MkdirAll(alterDirPath, os.ModePerm)
		if err != nil {
			err := fmt.Errorf("could not create alter-dir-path %s", alterDirPath)
			return err
		}
		return nil
	}

	err := checkParams(alterDirPath)
	if err != nil {
		return err
	}

	return nil
}
