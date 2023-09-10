package pick

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string, content string, queryFilePath string) error {
	// fullQueryIdWithoutQuotes := strings.ReplaceAll(fullQueryId, globals.QUOTE, "")
	// idParts := strings.Split(fullQueryIdWithoutQuotes, ".")
	// queryId := idParts[len(idParts)-1]
	// queryIdWithQuotes := globals.QUOTE + queryId + globals.QUOTE

	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content: content,
// 			Content: `
// {
//   "__CONTENT": [
// 	{
// 	  "id": ` + queryIdWithQuotes + `,
// 	  "kind": "multiselect",
// 	  "prompt": "enter ...",
// 	  "selector": [
// 		0,
// 		1
// 	  ],
// 	  "children": {
// 		"kind": "literal",
// 		"value": [
// 		  "_app_js",
// 		  "_document_js"
// 		]
// 	  }
// 	}
//   ]
// }
// `,
		},
	}

	err := common.CreateFiles(queryTokenScaffold)
	if err != nil {
		return err
	}

	return nil
}
