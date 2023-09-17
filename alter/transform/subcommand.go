package transform

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	alterRecord, err := common.BuildAlterInfrastucture(templateMap, TransformQueryTemplate, TransformControlTemplate)
	if err != nil {
		return err
	}

	fullQueryId := alterRecord.FullQueryId
	quotedFullQueryId := globals.QUOTE + fullQueryId + globals.QUOTE


	prependString := "{%- set name = val(tokens, " + quotedFullQueryId + " -%}" + "\n"
	
	err = common.PrependToFile(alterRecord.FirstFilePath, prependString)
	return err
}
