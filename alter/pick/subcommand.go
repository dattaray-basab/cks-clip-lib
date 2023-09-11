package pick

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/templates"
)



var BuildSubcommand = func(templateMap map[string]string) error {
	err := common.BuildAlterInfrastucture(templateMap, templates.PickQueryTemplate, templates.PickControlTemplate)
	return err
}
