package alter

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var UpdatePhaseFile = func(templateMap map[string]string, moveItems []string, phasePath string, phaseName, lastPhase string) error {
	phaseFilePath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	log.Println(phaseFilePath)
	_, err := os.Create(phaseFilePath)
	if err != nil {
		return err
	}
	return nil
}
