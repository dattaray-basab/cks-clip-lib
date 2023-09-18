package transform

import (
	"io/ioutil"
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	alterRecord, err := common.BuildAlterInfrastucture(templateMap, TransformQueryTemplate, TransformControlTemplate)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(alterRecord.FirstFilePath)
	if err != nil {
		return err
	}
	text := string(data)
	changedText := strings.Replace(text, alterRecord.FirstWordInFirstFile, "{{name}}", 1)
	changedData := []byte(changedText)
	err = ioutil.WriteFile(alterRecord.FirstFilePath, changedData, 0644)
	if err != nil {
		return err
	}

	quotedFullQueryId := globals.QUOTE + alterRecord.FullQueryId + globals.QUOTE

	prependString := "{%- set name = val(tokens, " +  quotedFullQueryId + ") -%}" + "\n"

	err = common.PrependToFile(alterRecord.FirstFilePath, prependString)
	return err
}
