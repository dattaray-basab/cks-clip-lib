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

	var checkParams = func() error {
		var calcBlockPath = func() (string, error) {
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
			codeBlockPath = filepath.Join(codeBlockPath, alterName)
			if common.IsDir(codeBlockPath) {
				err := fmt.Errorf("code-block-path %s already exists", codeBlockPath)
				return codeBlockPath, err
			}
			err := os.MkdirAll(codeBlockPath, os.ModePerm)
			if err != nil {
				err := fmt.Errorf("could not create code-block-path %s", codeBlockPath)
				return codeBlockPath, err
			}

			return codeBlockPath, nil
		}
		codeBlockPath, err := calcBlockPath()
		if err != nil {
			return err
		}
		println(codeBlockPath)

		return nil
	}

	err := checkParams()
	if err != nil {
		return err
	}

	return nil
}
