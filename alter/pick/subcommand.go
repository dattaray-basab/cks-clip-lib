package pick

import "github.com/dattaray-basab/cks-clip-lib/common"

var BuildSubcommand = func(templateMap map[string]string) error {
	_, err := common.BuildAlterInfrastucture(templateMap, QueryTemplate, ControlTemplate)
	return err
}
