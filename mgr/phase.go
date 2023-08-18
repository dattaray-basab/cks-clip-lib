package mgr

import (
	"fmt"
	"strings"

	"github.com/dattaray-basab/cks-ckip-lib/common"
)

func CreatePhase(

	templaterDirpath string,
	phaseName string,
	codeBlock string,
	dependentPhaseNames string) {

	fmt.Println("")
	fmt.Println("<<< CreatePhase/n")

	recipePath := common.GetRecipePath(templaterDirpath, false)
	fmt.Println("templaterDirpath::", templaterDirpath)
	fmt.Println("recipePath::", recipePath)
	if templaterDirpath != recipePath && !strings.HasPrefix(templaterDirpath, recipePath) {
		_ = fmt.Errorf("Recipe path %s is not a subdirectory of %s", templaterDirpath, recipePath)
	}
	fmt.Println(">>> CreatePhase")
	fmt.Println()
}
