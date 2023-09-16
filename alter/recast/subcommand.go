package recast

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
)



var BuildSubcommand = func(templateMap map[string]string) error {
	err := common.BuildAlterInfrastucture(templateMap, RecastQueryTemplate, RecastControlTemplate)
	return err
}
