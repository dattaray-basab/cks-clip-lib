package pick

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
	"op": "pick",
	"directives": {
	  "token_id": "unknown_token_file.unkown_token_id",
	  "options": [
		{
		  "rel_paths": [
		  ],
		  "sift": ""
		}
	  ]
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
