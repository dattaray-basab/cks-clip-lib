package transform

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	err := common.BuildAlterInfrastucture(templateMap, TransformQueryTemplate, TransformControlTemplate)
	return err
}
