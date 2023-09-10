package pick

import (
	"strings"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string, moveItemMap map[string]globals.MoveItemDetailsT, queryFilePath string, fullQueryId string) error {
	fullQueryIdWithoutQuotes := strings.ReplaceAll(fullQueryId, globals.QUOTE, "")
	idParts := strings.Split(fullQueryIdWithoutQuotes, ".")
	queryId := idParts[len(idParts)-1]
	queryIdWithQuotes := globals.QUOTE + queryId + globals.QUOTE

	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content: `
{
  "__CONTENT": [
	{
	  "id": ` + queryIdWithQuotes + `,
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
