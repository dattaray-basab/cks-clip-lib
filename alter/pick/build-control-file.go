package pick

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeControlFile = func(templateMap map[string]string, moveItemMap map[string]globals.MoveItemDetailsT, fullQueryId string) error {

	jsonControlFileScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: templateMap[globals.KEY_CONTROL_JSON_PATH],
			Content: `
[
  {
	"op": "pick",
	"directives": {
	  "token_id": ` + fullQueryId + `,
	  "options": [
		{
		  "rel_paths": [
			"_app.js"
		  ],
		  "sift": "_app_js"
		},
		{
		  "rel_paths": [
			"_document.js"
		  ],
		  "sift": "_document_js"
		}
	  ]
	}
  }
]
`,
		},
	}

	err := common.CreateFiles(jsonControlFileScaffold)
	if err != nil {
		return err
	}
	return nil
}
