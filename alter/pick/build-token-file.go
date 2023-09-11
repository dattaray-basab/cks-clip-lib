package pick

import (
	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string, content string, queryFilePath string) error {

	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content: content,
		},
	}

	err := common.CreateFiles(queryTokenScaffold)
	if err != nil {
		return err
	}

	return nil
}
