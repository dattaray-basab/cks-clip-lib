package alter

import (
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var BuildTransformControl = func(templateMap map[string]string) error {
	scaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: filepath.Join(templateMap[globals.KEY_CONTROL_JSON_PATH]),
			Content: `
	[
	{
		"op": "transform"
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
