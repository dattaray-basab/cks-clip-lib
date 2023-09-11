package pick

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/templates"
)



var BuildSubcommand = func(templateMap map[string]string) error {
	pickQueryTemplate := templates.PickQueryTemplate
	pickControlTemplate := templates.PickControlTemplate

	err := common.BuildAlterInfrastucture(templateMap, pickQueryTemplate, pickControlTemplate)
	return err
}
