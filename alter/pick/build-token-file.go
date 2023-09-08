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
        ` + fullQueryId + `: [
            "attractions",
            "restaurants",
            "events"
        ]
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
