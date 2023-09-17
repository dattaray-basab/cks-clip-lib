package filegen

import (
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/logger"
)

func CreatePhaseAndMiscFilesAndRun(baseDirpath string, tokenFileName string) error {
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
  "MODE_LOG": false,
  "RESPONSE_CONTEXT": "1_response_context",
  "OPERATION_FOLDER_PREFIX": "__",
  "APP_PREFIX": "__cks_codegen.",
  "ENFORCED_APP_NAME": "__CKS_codegen.nextjs"
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
from cks_codegen.main._main_generator import fn_start

error, _ = fn_start(__file__)
if error is not None:
	print(error)
    exit(1)
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, globals.RECIPE_LOCATOR),
			Content: `
[
]
		`,
		},
	}

	err = common.CreateFiles(recipeScaffold)
	return err
}
