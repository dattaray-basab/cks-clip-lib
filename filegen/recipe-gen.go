package filegen

import (
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func CreateRecipeFiles(baseDirpath string, queryToken string) error {
	err := os.RemoveAll(baseDirpath)
	if err != nil {
		return err
	}

	recipeScaffold := globals.ScaffoldInfoTListT{
		{
			Filepath: filepath.Join(baseDirpath, globals.BLUEPRINTS_DIRNAME, globals.TOKENS_DIRNAME, globals.QUERY_DIRNAME, queryToken + globals.JSON_EXT),
			Content: `
{
  "__SCHEMA": "s_context",
  "__CONTENT": [
  ]
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, globals.BLUEPRINTS_DIRNAME, "{{target}}", globals.MISC_DIRNAME, globals.DIRECTIVES_JSON),
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
			Filepath: filepath.Join(baseDirpath, globals.BLUEPRINTS_DIRNAME, "{{target}}", globals.PHASES_DIRNAME, "{{phase_name}}" + globals.JSON_EXT),
			Content: `
{
  "__CODE_BLOCK": "{{code_block}}",
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
			Filepath: filepath.Join(baseDirpath, globals.BLUEPRINTS_DIRNAME, "{{target}}", globals.RUN_PY),
			Content: `
from code_transformer.src.main._main_generator import fn_start

error, _ = fn_start(__file__)
if error is not None:
    exit(1)
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, globals.RECIPE_CONFIG_),
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
