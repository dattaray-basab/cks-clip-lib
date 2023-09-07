package recast

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	scaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(templateMap[globals.KEY_CONTROL_JSON_PATH]),
			Content: `
	[
		{
			"op": "recast",
				"directives": {
				"declare": {
					"name": "unknown-token-file.id_unknown_token",
				}
			}
		}
	]
			`,
		},
	}

	err := common.CreateFiles(scaffold)
	if err != nil {
		return err
	}

	return nil
}
