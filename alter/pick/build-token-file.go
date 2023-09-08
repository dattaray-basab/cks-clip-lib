package pick

import (

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string,  queryFilePath string, fullQueryId string) error {
	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content: `
[
	{
	  ' + fullQueryId + ': {',
	  "kind": "multiselect",
	  "prompt": "enter ...",
	  "selector": [
		0,
		1
	  ],
	  "children": {
		"kind": "literal",
		"value": [
		  "_app_js",
		  "_document_js"
		]
	  }
	}
]
`,
		},
	}

	err := common.CreateFiles(queryTokenScaffold)
	if err != nil {
		return err
	}

	return nil
}
