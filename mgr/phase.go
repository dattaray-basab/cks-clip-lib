package mgr

import (
	"fmt"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
)

func CreatePhase(

	clipDirpath string,
	phaseName string,
	codeBlock string,
	dependentPhaseNames string) {

	fmt.Println("")
	fmt.Println("<<< CreatePhase/n")

	recipePath := common.GetRecipePath(clipDirpath, false)
	fmt.Println("clipDirpath::", clipDirpath)
	fmt.Println("recipePath::", recipePath)
	if clipDirpath != recipePath && !strings.HasPrefix(recipePath, clipDirpath) {
		fmt.Printf("Recipe path %s is not a subdirectory of %s", recipePath, clipDirpath)
	}
	fmt.Println(">>> CreatePhase")
	fmt.Println()
}
