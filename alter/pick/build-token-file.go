package pick

import (
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string, queryFilePath string, fullQueryId string) error {
	idParts := strings.Split(fullQueryId, ".")
	queryId := idParts[len(idParts)-1]

	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content: `
{
  "__CONTENT": [
	{
	  "id": ` + queryId + `,
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
}
`,
		},
	}

	err := common.CreateFiles(queryTokenScaffold)
	if err != nil {
		return err
	}

	return nil
}
