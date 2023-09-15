package filegen

import (
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

func CreatePhaseAndMiscFiles(baseDirpath string, tokenFileName string) error {
	err := os.RemoveAll(baseDirpath)
	if err != nil {
		return err
	}
	rootPath := filepath.Dir(baseDirpath)
	tokenFilePath := filepath.Join(rootPath, globals.TOKENS_DIRNAME, globals.QUERY_DIRNAME, tokenFileName+globals.JSON_EXT)
	logger.Log.Debug(tokenFilePath)

	recipeScaffold := globals.ScaffoldInfoTListT{
		{
			Filepath: tokenFilePath,
			Content: `
{
  "__CONTENT": [
  ]
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, globals.MISC_DIRNAME, globals.DIRECTIVES_JSON),
			Content: `
{
  "LOG_LEVEL": 10,
  "MODE_LOG": true,
  "RESPONSE_CONTEXT": "1_response_context",
  "OPERATION_FOLDER_PREFIX": "__"
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, globals.PHASES_DIRNAME, "{{phase-name}}"+globals.JSON_EXT),
			Content: `
{
  "__CODE_BLOCK": "{{code-block-name}}",
  "ops_pipeline": [
	{
	  "remove": {
		"filter": "[*]"
	  }
	},
	{
	  "copy": {
		"filter": "~[__*]|~[.DS_Store, .gitignore]"
	  }
	}
  ]
}
		`,
		},
		{
			Filepath: filepath.Join(rootPath, globals.RUN_PY),
			Content: `
from code_transformer.src.main._main_generator import fn_start

error, _ = fn_start(__file__)
if error is not None:
    exit(1)
		`,
		},
		{
			Filepath: filepath.Join(rootPath, globals.RECIPE_CONFIG_),
			Content: `
[
  "../../__BLUEPRINTS"
]
		`,
		},
	}

	err = common.CreateFiles(recipeScaffold)
	return err
}
