package filegen

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
)

func CreatePhaseFile(phasePath string) error {
	var getPhaseName = func(jsonMap map[string] interface{}) (string, error) {
		return "", nil
	}

	log.Println(phasePath)
	files, err := os.ReadDir(phasePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(phasePath, file.Name())
			log.Println(filePath)
			jsonMap, err := common.ReadJsonFile(filePath)
			if err != nil {
				return err
			}
			log.Println(jsonMap)
			phaseName, err := getPhaseName(jsonMap)
			if err != nil {
				return err
			}
			log.Println(phaseName)
		}
	}
	return nil
}
