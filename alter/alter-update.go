package alter

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var UpdatePhaseFile = func(phasePath string, phaseName, lastPhase string) error {
	phaseFilePath := filepath.Join(phasePath, phaseName+globals.JSON_EXT)
	log.Println(phaseFilePath)
	_, err := os.Create(phaseFilePath)
	if err != nil {
		return err
	}
	return nil
}
